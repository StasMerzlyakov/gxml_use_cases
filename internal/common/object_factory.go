package common

import (
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

var birthDateElement = xsd.NameAndNamespace{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "BirthDate",
}

func ResolveValidator(nameAndNamespace xsd.NameAndNamespace) (any, xsd2.IElementValidator) {
	switch nameAndNamespace {
	case birthDateElement:
		birthDateElementValidator := ComplexDateTypeValidator{}
		return &birthDateElementValidator
	}
	return nil
}
