package internal

import (
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

const cardDataElementNS = "https://github.com/StasMerzlyakov/gxml/card-service"
const cardDataElementName = "CardData"

const cardTypeElementNS = "https://github.com/StasMerzlyakov/gxml/common-data"
const cardTypeElementName = "CardType"

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
}
