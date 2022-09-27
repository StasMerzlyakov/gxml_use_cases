package xsd2

import (
	"bufio"
	"fmt"
	"github.com/StasMerzlyakov/gxml/parser"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"reflect"
	"strings"
)

type IXsdValidator2 interface {
	Validate() (any, error)
	Write(writer bufio.Writer, any, name xsd.NameAndNamespace) error
}

type IValidatorResolver interface {
	ResolveValidator(nameAndNamespace xsd.NameAndNamespace) IElementValidator
	GetNamespacesMap() map[string]string
}

type Validator2 struct {
	XmlParser parser.IXmlParser
	Resolver  IValidatorResolver
}

type anyStack struct {
	data []any
}

func (s *anyStack) Pop() (v any) {
	l := len(s.data)
	rv := s.data[l-1]
	s.data = s.data[:l-1]
	return rv
}

func (s *anyStack) Push(v any) {
	s.data = append(s.data, v)
}

func (s *anyStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *anyStack) Peek() (v any) {
	l := len(s.data)
	return s.data[l-1]
}

type elementValidatorStack struct {
	data []IElementValidator
}

func (s *elementValidatorStack) Pop() (v IElementValidator) {
	l := len(s.data)
	rv := s.data[l-1]
	s.data = s.data[:l-1]
	return rv
}

func (s *elementValidatorStack) Push(v IElementValidator) {
	s.data = append(s.data, v)
}

func (s *elementValidatorStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *elementValidatorStack) Peek() (v IElementValidator) {
	l := len(s.data)
	return s.data[l-1]
}

func (xv *Validator2) writeElement(writer bufio.Writer, parentObj any, validator IElementValidator, tab string) error {
	for _, state := range validator.GetStates() {
		methodName := "Get" + state.Name
		value := reflect.ValueOf(parentObj).MethodByName(methodName).Call([]reflect.Value{})[0]
		if value.IsNil() {
			continue
		}
		obj := value.Interface()
		prefix := xv.Resolver.GetNamespacesMap()[state.Namespace]
		if !validator.IsComplexType() {
			vl := reflect.ValueOf(obj).MethodByName("GetXmlValue").Call([]reflect.Value{})[0].String()
			switch state.Type {
			case xsd.ElementNode:
				{
					writer.WriteString(fmt.Sprintf(tab+"<%s:%s>", prefix, state.Name))
					writer.WriteString(vl)
					writer.WriteString(fmt.Sprintf(tab+"</%s:%s>", prefix, state.Name))
				}
			case xsd.AttributeNode:
				{
					writer.WriteString(fmt.Sprintf(" %s:%s=\"", prefix, state.Name))
					writer.WriteString(vl)
					writer.WriteString("\"")
				}
			}
		} else {
			if tab == "" {
				writer.WriteString(tab + fmt.Sprintf("<%s:%s", prefix, state.Name))
				for k, v := range xv.Resolver.GetNamespacesMap() {
					writer.WriteString(fmt.Sprintf("\n    xmlns:%s=\"%s\"", v, k))
				}
				writer.WriteString(">\n")
			} else {
				writer.WriteString(tab + fmt.Sprintf("<%s:%s>\n", prefix, state.Name))
			}
			nextValidator := validator.ResolveValidator(state)
			xv.writeElement(writer, obj, nextValidator, tab+"  ")
			writer.WriteString(tab + fmt.Sprintf("</%s:%s>\n", prefix, state.Name))
		}
	}
	return nil
}

func (xv *Validator2) Write(writer bufio.Writer, obj any, name xsd.NameAndNamespace) error {
	currentValidator := xv.Resolver.ResolveValidator(name)
	writer.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>")
	xv.writeElement(writer, obj, currentValidator, "")
	return nil
}

func (xv *Validator2) Validate() (any, error) {
	var elementStack util.Stack[parser.Element]
	var objectStack anyStack
	var elementValidatorStack elementValidatorStack
	var rootObject any
	inXmlDecl := false
	for {
		if token, err := xv.XmlParser.Next(); err != nil {
			return nil, err
		} else {
			if token == nil {
				return rootObject, nil
			}
			switch token.XmlEventType {
			case parser.XMLDeclStart:
				inXmlDecl = true
				continue
			case parser.XMLDeclEnd:
				inXmlDecl = false
				continue
			case parser.STag:
				currentElement := *xv.XmlParser.CurrentElement()
				var currentValidator IElementValidator
				if !elementValidatorStack.IsEmpty() {
					currentValidator = elementValidatorStack.Peek()
					elementPrefix := currentElement.Name.Prefix
					if elementPrefix == "" {
						// TODO проверка currentValidator.GetElementFormDefault
						// elementPrefix = ...
					}
					elementNamespace, err := xv.XmlParser.GetNamespaceByPrefix(elementPrefix)
					if err != nil {
						return nil, err
					}

					elementData := xsd.ElementData{
						Namespace: elementNamespace,
						Name:      currentElement.Name.Name,
						Type:      xsd.ElementNode,
					}
					if err := currentValidator.AcceptElement(elementData); err != nil {
						return nil, err
					}

					nextValidator := currentValidator.ResolveValidator(elementData)
					if nextValidator == nil {
						return nil, fmt.Errorf("validator for %s not found", elementData.ToString())
					}
					elementValidatorStack.Push(nextValidator)
					if nextValidator.IsComplexType() {
						obj, _ := nextValidator.GetInstance()
						methodName := "Set" + elementData.Name
						parentObj := objectStack.Peek()
						reflect.ValueOf(parentObj).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(obj)})
						objectStack.Push(obj)
					}
				} else {
					elementName := xv.XmlParser.CurrentElement().Name
					namespace, err := xv.XmlParser.GetNamespaceByPrefix(elementName.Prefix)
					if err != nil {
						return nil, err
					}
					nameAndNamespace := xsd.NameAndNamespace{
						Namespace: namespace,
						Name:      elementName.Name,
					}
					nextValidator := xv.Resolver.ResolveValidator(nameAndNamespace)
					if nextValidator == nil {
						return nil, fmt.Errorf("validator for %s not found", nameAndNamespace.ToString())
					}
					if nextValidator.IsComplexType() {
						obj, _ := nextValidator.GetInstance()
						rootObject = obj
						objectStack.Push(obj)
					}
					elementValidatorStack.Push(nextValidator)
				}
				elementStack.Push(currentElement)

			case parser.ETagEnd, parser.EmptyElemEnd:
				currentValidator := elementValidatorStack.Peek()
				if err := currentValidator.CompleteElement(); err != nil {
					return nil, err
				}
				if currentValidator.IsComplexType() {
					objectStack.Pop()
					elementStack.Pop()
				} else {
					obj, err := currentValidator.GetInstance()
					if err != nil {
						return nil, err
					}
					currentElement := elementStack.Pop()
					methodName := "Set" + currentElement.Name.Name
					reflect.ValueOf(objectStack.Peek()).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(obj)})
				}
				elementValidatorStack.Pop()
			case parser.CharData:
				currentValidator := elementValidatorStack.Peek()
				if err := currentValidator.CheckValue(token.Runes); err != nil {
					return nil, err
				}
			case parser.Attr:
				if inXmlDecl {
					continue
				}
				currentAttribute := parseAttribute(string(token.Runes))
				if currentAttribute.Prefix != xmlns {
					if elementValidatorStack.IsEmpty() {
						return nil, fmt.Errorf("attribute must be in element")
					}
					currentValidator := elementValidatorStack.Peek()
					attributePrefix := currentAttribute.Prefix
					if attributePrefix == "" {
						// TODO проверка currentValidator.GetElementFormDefault
						// attributePrefix = ...
					}
					if attributeNamespace, err := xv.XmlParser.GetNamespaceByPrefix(attributePrefix); err != nil {
						return nil, err
					} else {
						elementData := xsd.ElementData{
							Namespace: attributeNamespace,
							Name:      currentAttribute.Name,
							Type:      xsd.AttributeNode,
						}
						if err := currentValidator.AcceptElement(elementData); err != nil {
							return nil, err
						}
						nextValidator := currentValidator.ResolveValidator(elementData)
						if nextValidator == nil {
							return nil, fmt.Errorf("validator for %s not found", elementData.ToString())
						}
						nextValidator.CheckValue([]rune(currentAttribute.Value))

						obj, err := nextValidator.GetInstance()
						if err != nil {
							return nil, err
						}
						methodName := "Set" + currentAttribute.Name
						reflect.ValueOf(objectStack.Peek()).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(obj)})
					}
				}
			}
		}
	}
}

const attrValueDelimiter = "="

func parseAttribute(attributeDecl string) parser.Attribute {
	list := strings.Split(attributeDecl, attrValueDelimiter)
	value := list[1]
	value = value[1 : len(value)-1]

	prefixAndName := parseElementName(list[0])
	return parser.Attribute{Value: value, Name: prefixAndName.Name, Prefix: prefixAndName.Prefix}
}

func parseElementName(elementName string) parser.ElementName {
	prefixAndName := elementName

	if strings.HasPrefix(prefixAndName, etagStart) {
		prefixAndName = prefixAndName[etagStartLen:]
	}
	if strings.HasPrefix(prefixAndName, stagStart) {
		prefixAndName = prefixAndName[stagStartLen:]
	}

	names := strings.Split(prefixAndName, separator)
	if len(names) == 1 {
		return parser.ElementName{Prefix: "", Name: names[0]}
	} else {
		return parser.ElementName{Prefix: names[0], Name: names[1]}
	}
}

const separator = ":"
const etagStart = "</"
const etagStartLen = len(etagStart)
const stagStart = "<"
const stagStartLen = len(stagStart)
const xmlns = "xmlns"

type IElementValidator interface {
	AcceptElement(elementData xsd.ElementData) error
	CompleteElement() error
	CheckValue(runes []rune) error
	ResolveValidator(elementData xsd.ElementData) IElementValidator
	GetInstance() (any, error)
	IsComplexType() bool
	GetStates() []xsd.ElementData
	//GetValue() string
	// TODO GetAttributeFormDefault and GetElementFormDefault support;
	// current implementation work as AttributeFormDefault=qualified && ElementFormDefault=qualified
}
