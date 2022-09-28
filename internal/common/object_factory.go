package common

import (
	"github.com/StasMerzlyakov/gxml/xsd"
)

var birthDateElement = xsd.NameAndNamespace{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "BirthDate",
}

func ResolveValidator(nameAndNamespace xsd.NameAndNamespace) xsd.IElementValidator {
	switch nameAndNamespace {
	case birthDateElement:
		birthDateElementValidator := ComplexDateTypeValidator{}
		return &birthDateElementValidator
	}
	return nil
}
