package crd

import (
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

var cardRequestElement = xsd.NameAndNamespace{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CardRequest",
}

var cardResponseElement = xsd.NameAndNamespace{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CardResponse",
}

func ResolveValidator(nameAndNamespace xsd.NameAndNamespace) xsd2.IElementValidator {
	switch nameAndNamespace {
	case cardRequestElement:
		str := cardRequestTypeValidator{}
		return &str
	case cardResponseElement:
		str := CardResponseTypeStruct{}
		return &str
	}
	return nil
}
