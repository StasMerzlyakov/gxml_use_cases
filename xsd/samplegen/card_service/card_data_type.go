package card_service

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml_use_cases/xml"
)

type cardDataTypeValidator struct {
	state cardDataTypeState
}

func (cv *cardDataTypeValidator) unexpectedElementError(elementType xml.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected ", elementType.ToString())
	expectedStates := CardDataTypeStateAcceptableMap[cv.state]
	if expectedStates == nil {
		result += " end of element."
	} else {
		for idx, est := range expectedStates {
			expectedElement := cardDataTypeStateToElement[est]
			if idx == 0 {
				result += " " + expectedElement.ToString()
			} else {
				result += ", " + expectedElement.ToString()
			}
		}
		result += "."
	}
	return errors.New(result)
}

func (cv *cardDataTypeValidator) AcceptElement(elementType xml.ElementData) error {
	if state, ok := CardDataTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := CardDataTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
}

type cardDataTypeCreator struct {
}

func (cardDataTypeCreator) create() xml.IElementValidator {
	validator := cardDataTypeValidator{
		state: cardDataTypeStateInit,
	}
	return &validator
}

type cardDataTypeState int

const cardDataTypeCommon = "https://github.com/StasMerzlyakov/gxml/card-service"

var cardDataTypeElementData1 = xml.ElementData{
	Namespace: cardDataTypeCommon,
	Name:      "CardType",
	Type:      xml.ElementNode,
}

var cardDataTypeElementData2 = xml.ElementData{
	Namespace: cardDataTypeCommon,
	Name:      "CardNumberType",
	Type:      xml.ElementNode,
}

var cardDataTypeElementData3 = xml.ElementData{
	Namespace: cardDataTypeCommon,
	Name:      "ExpirationDateType",
	Type:      xml.ElementNode,
}

var cardDataTypeStateToElement = map[cardDataTypeState]xml.ElementData{
	cardDataTypeState1: cardDataTypeElementData1,
	cardDataTypeState2: cardDataTypeElementData2,
	cardDataTypeState3: cardDataTypeElementData3,
}

var CardDataTypeElementToState = map[xml.ElementData]cardDataTypeState{
	cardDataTypeElementData1: cardDataTypeState1,
	cardDataTypeElementData2: cardDataTypeState2,
	cardDataTypeElementData3: cardDataTypeState3,
}

var CardDataTypeStateAcceptableMap = map[cardDataTypeState][]cardDataTypeState{
	cardDataTypeStateInit: {cardDataTypeState1},
	cardDataTypeState1:    {cardDataTypeState2},
	cardDataTypeState2:    {cardDataTypeState3},
	cardDataTypeState3:    nil,
}

const (
	cardDataTypeStateInit cardDataTypeState = 0
	cardDataTypeState1    cardDataTypeState = 1
	cardDataTypeState2    cardDataTypeState = 2
	cardDataTypeState3    cardDataTypeState = 3
)
