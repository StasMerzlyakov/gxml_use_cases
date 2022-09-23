package internal

import (
	"fmt"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/crd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

var NamespacesToPackageMap = map[string]string{
	"https://github.com/StasMerzlyakov/gxml/card-service": "crd",
	"https://github.com/StasMerzlyakov/gxml/common-data":  "common",
}

var PackageToNamespacesMap = xsd2.ReverseMap(NamespacesToPackageMap)

func ResolveValidator(name xsd2.NameAndNamespace) (xsd.IElementValidator, error) {
	if packageName, ok := NamespacesToPackageMap[name.Namespace]; !ok {
		return nil, fmt.Errorf("namespace %s not known", name.Namespace)
	} else {
		switch packageName {
		case "crd":
			return crd.ResolveValidator(name)
		case "common":
			return common.ResolveValidator(name)
		default:
			return nil, fmt.Errorf("element %s not known", name.ToString())
		}
	}
}

const cardDataElementNS = "https://github.com/StasMerzlyakov/gxml/card-service"
const cardDataElementName = "CardData"

const cardTypeElementNS = "https://github.com/StasMerzlyakov/gxml/common-data"
const cardTypeElementName = "CardType"

const cardNumberElementNS = "https://github.com/StasMerzlyakov/gxml/common-data"
const cardNumberElementName = "CardNumber"

const expirationDateElementNS = "https://github.com/StasMerzlyakov/gxml/common-data"
const expirationDateElementName = "ExpirationDate"

var ElementValidatorMap = map[xsd.ElementData]xsd2.ICreator{
	xsd.ElementData{
		Namespace: cardDataElementNS,
		Name:      cardDataElementName,
		Type:      xsd.ElementNode,
	}: cardDataTypeCreator{},
	xsd.ElementData{
		Namespace: cardTypeElementNS,
		Name:      cardTypeElementName,
		Type:      xsd.ElementNode,
	}: cardTypeTypeCreator{},
	xsd.ElementData{
		Namespace: cardNumberElementNS,
		Name:      cardNumberElementName,
		Type:      xsd.ElementNode,
	}: cardNumberTypeCreator{},
	xsd.ElementData{
		Namespace: expirationDateElementNS,
		Name:      expirationDateElementName,
		Type:      xsd.ElementNode,
	}: cardNumberTypeCreator{},
}
