package card_service

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml_use_cases/xml"
)

type CardRequestTypeValidator struct {
	state cardRequestTypeState
}

func (cv *CardRequestTypeValidator) unexpectedElementError(elementType xml.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected ", elementType.ToString())
	expectedStates := cardRequestTypeStateAcceptableMap[cv.state]
	if expectedStates == nil {
		result += " end of element."
	} else {
		for idx, est := range expectedStates {
			expectedElement := cardRequestTypeStateToElement[est]
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

func (cv *CardRequestTypeValidator) AcceptElement(elementType xml.ElementData) error {
	if state, ok := cardRequestTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := cardRequestTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
}

type cardRequestTypeCreator struct {
}

func (cardRequestTypeCreator) create() xml.IElementValidator {
	validator := CardRequestTypeValidator{
		state: cardRequestTypeStateInit,
	}
	return &validator
}

type cardRequestTypeState int

const cardRequestTypeCard = "https://github.com/StasMerzlyakov/gxml/card-service"
const cardRequestTypeCd = "https://github.com/StasMerzlyakov/gxml/common-data"

var cardRequestTypeElementData1 = xml.ElementData{
	Namespace: cardRequestTypeCard,
	Name:      "FirstName",
	Type:      xml.ElementNode,
}

var cardRequestTypeElementData2 = xml.ElementData{
	Namespace: cardRequestTypeCard,
	Name:      "LastName",
	Type:      xml.ElementNode,
}
var cardRequestTypeElementData3 = xml.ElementData{
	Namespace: cardRequestTypeCard,
	Name:      "Patronymic",
	Type:      xml.ElementNode,
}
var cardRequestTypeElementData4 = xml.ElementData{
	Namespace: cardRequestTypeCd,
	Name:      "BirthDate",
	Type:      xml.ElementNode,
}

var cardRequestTypeElementData5 = xml.ElementData{
	Namespace: cardRequestTypeCard,
	Name:      "CardType",
	Type:      xml.ElementNode,
}

var cardRequestTypeStateToElement = map[cardRequestTypeState]xml.ElementData{
	cardRequestTypeState1: cardRequestTypeElementData1,
	cardRequestTypeState2: cardRequestTypeElementData2,
	cardRequestTypeState3: cardRequestTypeElementData3,
	cardRequestTypeState4: cardRequestTypeElementData4,
	cardRequestTypeState5: cardRequestTypeElementData5,
}

var cardRequestTypeElementToState = map[xml.ElementData]cardRequestTypeState{
	cardRequestTypeElementData1: cardRequestTypeState1,
	cardRequestTypeElementData2: cardRequestTypeState2,
	cardRequestTypeElementData3: cardRequestTypeState3,
	cardRequestTypeElementData4: cardRequestTypeState4,
	cardRequestTypeElementData5: cardRequestTypeState5,
}

var cardRequestTypeStateAcceptableMap = map[cardRequestTypeState][]cardRequestTypeState{
	cardRequestTypeStateInit: {cardRequestTypeState1, cardRequestTypeState2, cardRequestTypeState3, cardRequestTypeState4},
	cardRequestTypeState1:    {cardRequestTypeState2, cardRequestTypeState3, cardRequestTypeState4},
	cardRequestTypeState2:    {cardRequestTypeState3, cardRequestTypeState4},
	cardRequestTypeState3:    {cardRequestTypeState4},
	cardRequestTypeState4:    {cardRequestTypeState5},
	cardRequestTypeState5:    nil,
}

const (
	cardRequestTypeStateInit cardRequestTypeState = 0
	cardRequestTypeState1    cardRequestTypeState = 1
	cardRequestTypeState2    cardRequestTypeState = 2
	cardRequestTypeState3    cardRequestTypeState = 3
	cardRequestTypeState4    cardRequestTypeState = 4
	cardRequestTypeState5    cardRequestTypeState = 5
)
