package xsd2

import (
	"bufio"
	"github.com/StasMerzlyakov/gxml/api"
	"github.com/StasMerzlyakov/gxml/buffer"
	"github.com/StasMerzlyakov/gxml/parser"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"strings"
)

type IElementValidatorResolver interface {
	Resolve(elementData xsd.ElementData) xsd.IElementValidator
}

type IXsdValidator interface {
	Validate() ([]string, error)
}

type ICreator interface {
	Create() xsd.IElementValidator
}

type Validator struct {
	XmlParser                  parser.IXmlParser
	ElementValidatorCreatorMap map[xsd.ElementData]ICreator
	SkipEmptyCharData          bool
}

func (xv *Validator) Validate() ([]string, error) {

	var validationErrorList []string
	var elementStack util.Stack[parser.Element]
	var elementValidatorStack ElementValidatorStack

	for {
		if token, err := xv.XmlParser.Next(); err != nil {
			return validationErrorList, err
		} else {
			if token == nil {
				return validationErrorList, nil
			}
			switch token.XmlEventType {
			case parser.STag:
				currentElement := *xv.XmlParser.CurrentElement()
				elementStack.Push(currentElement)
				if elementNamespace, err := xv.XmlParser.GetNamespaceByPrefix(currentElement.Name.Prefix); err != nil {
					return validationErrorList, err
				} else {
					elementData := xsd.ElementData{
						Namespace: elementNamespace,
						Name:      currentElement.Name.Name,
						Type:      xsd.ElementNode,
					}

					// check by current validator
					if !elementValidatorStack.IsEmpty() {
						currentValidator := elementValidatorStack.Peek()
						if err := currentValidator.AcceptElement(elementData); err != nil {
							return validationErrorList, err
						}
					}
					nextValidator := xv.ElementValidatorCreatorMap[elementData].Create()
					elementValidatorStack.Push(nextValidator)
				}
			case parser.ETagEnd, parser.EmptyElemEnd:
				currentValidator := elementValidatorStack.Peek()
				if err := currentValidator.CompleteElement(); err != nil {
					return validationErrorList, err
				}
				elementValidatorStack.Pop()
				elementStack.Pop()
			case parser.CharData:
				currentValidator := elementValidatorStack.Peek()
				if currentValidator != nil {
					elementData := xsd.ElementData{
						Type: xsd.CharData,
					}
					if err := currentValidator.AcceptElement(elementData); err != nil {
						return validationErrorList, err
					}

					if err := currentValidator.CheckValue(token.Runes); err != nil {
						return validationErrorList, err
					}
				}

			case parser.Attr:
				currentValidator := elementValidatorStack.Peek()
				// TODO перейти на CurrentAttribute
				currentAttribute := parseAttribute(string(token.Runes))
				if currentAttribute.Prefix != xmlns {
					if attributeNamespace, err := xv.XmlParser.GetNamespaceByPrefix(currentAttribute.Prefix); err != nil {
						return validationErrorList, err
					} else {

						elementData := xsd.ElementData{
							Namespace: attributeNamespace,
							Name:      currentAttribute.Name,
							Type:      xsd.AttributeNode,
						}
						if err := currentValidator.AcceptElement(elementData); err != nil {
							return validationErrorList, err
						}
					}
				}
			}
		}
	}
	return validationErrorList, nil
}

func NewXsdValidator(r *bufio.Reader) (IXsdValidator, error) {
	return NewXsdValidatorSize(r, buffer.RuneBufferSize)
}

func NewXsdValidatorSize(r *bufio.Reader, size int) (IXsdValidator, error) {
	if xmlParser, err := api.NewXmlParserSize(r, size); err != nil {
		return nil, err
	} else {
		validator := Validator{
			XmlParser: xmlParser,
		}
		return &validator, nil
	}
}

const attrValueDelimiter = "="

func isSpace(r rune) bool {
	return r == 0x20 || r == 0x9 || r == 0xd || r == 0xa
}

func IsEmpty(runes []rune) bool {
	isEmpty := true
	for _, r := range runes {
		if !isSpace(r) {
			isEmpty = false
			break
		}
	}
	return isEmpty
}

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
