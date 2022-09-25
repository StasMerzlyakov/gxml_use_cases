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

func ResolveValidator(nameAndNamespace xsd.NameAndNamespace) (any, xsd2.IElementValidator) {
	switch nameAndNamespace {
	case cardRequestElement:
		cardRequestElementValidator := cardRequestTypeValidator{}
		str := CardRequestTypeStruct{}
		return &str, &cardRequestElementValidator
	case cardResponseElement:
		str := CardResponseTypeStruct{}
		cardResponseElementValidator := cardResponseTypeValidator{}
		return &str, &cardResponseElementValidator
	}
	return nil
}
