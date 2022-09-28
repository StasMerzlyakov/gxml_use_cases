package api

import (
	"bufio"
	"github.com/StasMerzlyakov/gxml/api"
	"github.com/StasMerzlyakov/gxml/buffer"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal"
)

func NewXsdValidator(r *bufio.Reader) (xsd.IXsdValidator, error) {
	return NewXsdValidatorSize(r, buffer.RuneBufferSize)
}

func NewXsdValidatorSize(r *bufio.Reader, size int) (xsd.IXsdValidator, error) {
	if xmlParser, err := api.NewXmlParserSize(r, size); err != nil {
		return nil, err
	} else {
		validator := xsd.Validator{
			XmlParser: xmlParser,
			Resolver:  internal.Resolver{},
		}
		return &validator, err
	}
}
