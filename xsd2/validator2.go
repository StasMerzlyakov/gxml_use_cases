package xsd2

import (
	"fmt"
	"github.com/StasMerzlyakov/gxml/parser"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"strings"
)

type IXsdValidator2 interface {
	Validate() error
}

type Validator2 struct {
	XmlParser parser.IXmlParser
	Resolver  xsd.IValidatorResolver
}

type elementValidatorStack struct {
	data []xsd.IElementValidator
}

func (s *elementValidatorStack) Pop() (v xsd.IElementValidator) {
	l := len(s.data)
	rv := s.data[l-1]
	s.data = s.data[:l-1]
	return rv
}

func (s *elementValidatorStack) Push(v xsd.IElementValidator) {
	s.data = append(s.data, v)
}

func (s *elementValidatorStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *elementValidatorStack) Peek() (v xsd.IElementValidator) {
	l := len(s.data)
	return s.data[l-1]
}

func (xv *Validator2) Validate() error {
	var elementStack util.Stack[parser.Element]
	var elementValidatorStack elementValidatorStack
	inXmlDecl := false
	for {
		if token, err := xv.XmlParser.Next(); err != nil {
			return err
		} else {
			if token == nil {
				return nil
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
				var currentValidator xsd.IElementValidator
				if !elementValidatorStack.IsEmpty() {
					currentValidator = elementValidatorStack.Peek()
					elementPrefix := currentElement.Name.Prefix
					if elementPrefix == "" {
						// TODO проверка currentValidator.GetElementFormDefault
						// elementPrefix = ...
					}

					elementNamespace, err := xv.XmlParser.GetNamespaceByPrefix(elementPrefix)
					if err != nil {
						return err
					}

					elementData := xsd.ElementData{
						Namespace: elementNamespace,
						Name:      currentElement.Name.Name,
						Type:      xsd.ElementNode,
					}
					if err := currentValidator.AcceptElement(elementData); err != nil {
						return err
					}

					nextValidator := currentValidator.ResolveValidator(elementData)
					if nextValidator == nil {
						return fmt.Errorf("validator for %s not found", elementData.ToString())
					}
					elementValidatorStack.Push(nextValidator)
				} else {
					elementName := xv.XmlParser.CurrentElement().Name
					namespace, err := xv.XmlParser.GetNamespaceByPrefix(elementName.Prefix)
					if err != nil {
						return err
					}
					nameAndNamespace := xsd.NameAndNamespace{
						Namespace: namespace,
						Name:      elementName.Name,
					}
					nextValidator := xv.Resolver.ResolveValidator(nameAndNamespace)
					if nextValidator == nil {
						return fmt.Errorf("validator for %s not found", nameAndNamespace.ToString())
					}
					elementValidatorStack.Push(nextValidator)
				}
				elementStack.Push(currentElement)

			case parser.ETagEnd, parser.EmptyElemEnd:
				currentValidator := elementValidatorStack.Peek()
				if err := currentValidator.CompleteElement(); err != nil {
					return err
				}
				elementValidatorStack.Pop()
				elementStack.Pop()
			case parser.CharData:
				currentValidator := elementValidatorStack.Peek()
				if err := currentValidator.CheckValue(token.Runes); err != nil {
					return err
				}
			case parser.Attr:
				if inXmlDecl {
					continue
				}
				currentAttribute := parseAttribute(string(token.Runes))
				if currentAttribute.Prefix != xmlns {
					if elementValidatorStack.IsEmpty() {
						return fmt.Errorf("attribute must be in element")
					}
					currentValidator := elementValidatorStack.Peek()
					attributePrefix := currentAttribute.Prefix
					if attributePrefix == "" {
						// TODO проверка currentValidator.GetElementFormDefault
						// attributePrefix = ...
					}
					if attributeNamespace, err := xv.XmlParser.GetNamespaceByPrefix(attributePrefix); err != nil {
						return err
					} else {
						elementData := xsd.ElementData{
							Namespace: attributeNamespace,
							Name:      currentAttribute.Name,
							Type:      xsd.AttributeNode,
						}
						if err := currentValidator.AcceptElement(elementData); err != nil {
							return err
						}
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