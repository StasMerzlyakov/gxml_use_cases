package api

import (
	"bufio"
	"github.com/StasMerzlyakov/gxml/api"
	"github.com/StasMerzlyakov/gxml/buffer"
	"github.com/StasMerzlyakov/gxml_use_cases/internal"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

func NewXsdValidator(r *bufio.Reader) (xsd2.IXsdValidator2, error) {
	return NewXsdValidatorSize(r, buffer.RuneBufferSize)
}

func NewXsdValidatorSize(r *bufio.Reader, size int) (xsd2.IXsdValidator2, error) {
	if xmlParser, err := api.NewXmlParserSize(r, size); err != nil {
		return nil, err
	} else {
		validator := xsd2.Validator2{
			XmlParser: xmlParser,
			Resolver:  internal.Resolver{},
		}
		return &validator, err
	}
}
