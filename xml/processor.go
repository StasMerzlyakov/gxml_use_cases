package xml

import (
	"bufio"
	"fmt"
	"github.com/StasMerzlyakov/gxml/api"
	"github.com/StasMerzlyakov/gxml/buffer"
	"github.com/StasMerzlyakov/gxml/parser"
)

type ElementType int

// TODO mixed types
const (
	ElementNode   ElementType = 0
	AttributeNode ElementType = 1
	CharData      ElementType = 2
)

type ElementData struct {
	Namespace string
	Name      string
	Type      ElementType
}

func (ed *ElementData) ToString() string {
	if ed.Namespace == "" {
		return ed.Name
	} else {
		return fmt.Sprintf("{%s}%s", ed.Namespace, ed.Name)
	}
}

type IElementValidator interface {
	AcceptElement(elementData ElementData) error
}

type ElementValidator struct {
	xmlParser parser.IXmlParser
}

func NewElementValidator(r *bufio.Reader) (*ElementValidator, error) {
	return NewElementValidatorSize(r, buffer.RuneBufferSize)
}

func NewElementValidatorSize(r *bufio.Reader, size int) (*ElementValidator, error) {
	if xmlParser, err := api.NewXmlParserSize(r, size); err != nil {
		return nil, err
	} else {
		validator := ElementValidator{
			xmlParser: xmlParser,
		}
		//return &validator, nil
		return &validator, nil
	}
}
