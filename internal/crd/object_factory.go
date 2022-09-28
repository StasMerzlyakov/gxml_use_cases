package crd

import (
	"github.com/StasMerzlyakov/gxml/xsd"
)

var cardRequestElement = xsd.NameAndNamespace{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CardRequest",
}

var cardResponseElement = xsd.NameAndNamespace{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CardResponse",
}

func ResolveValidator(nameAndNamespace xsd.NameAndNamespace) xsd.IElementValidator {
	switch nameAndNamespace {
	case cardRequestElement:
		str := cardRequestTypeValidator{}
		return &str
	case cardResponseElement:
		str := cardResponseTypeValidator{}
		return &str
	}
	return nil
}
