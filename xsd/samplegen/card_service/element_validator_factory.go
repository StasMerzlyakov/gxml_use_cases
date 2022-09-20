package card_service

import (
	"github.com/StasMerzlyakov/gxml_use_cases/xml"
)

type creator interface {
	create() xml.IElementValidator
}

var elementValidatorMap = map[xml.ElementData]creator{
	xml.ElementData{
		Namespace: cardRequestTypeCard,
		Name:      "CardRequest",
		Type:      xml.ElementNode,
	}: cardRequestTypeCreator{},

	xml.ElementData{
		Namespace: cardDataTypeCommon,
		Name:      "CardDataType",
		Type:      xml.ElementNode,
	}: cardDataTypeCreator{},
}

func Resolve(elementData xml.ElementData) xml.IElementValidator {
	return elementValidatorMap[elementData].create()
}
