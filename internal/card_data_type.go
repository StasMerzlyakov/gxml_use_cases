package internal

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type cardDataTypeValidator struct {
	state cardDataTypeState
}

func (cv *cardDataTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("unexpected element %s: expected", elementType.ToString())
	expectedStates := cardDataTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := cardDataTypeStateToElement[est]
		if idx == 0 {
			result += expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *cardDataTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element: expected "
	expectedStates := cardDataTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := cardDataTypeStateToElement[est]
		if idx == 0 {
			result += " " + expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *cardDataTypeValidator) CheckValue(runes []rune) error {
	if !xsd2.IsEmpty(runes) {
		return errors.New("value unexpected")
	}
	return nil
}

func (cv *cardDataTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type == xsd.CharData {
		return nil
	}
	if state, ok := cardDataTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := cardDataTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
}

func (cv *cardDataTypeValidator) CompleteElement() error {
	// Проверка достижимости конечного состояния из текущего
	acceptableStates := cardDataTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardDataTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type cardDataTypeCreator struct {
}

func (cardDataTypeCreator) Create() xsd.IElementValidator {
	validator := cardDataTypeValidator{
		state: cardDataTypeStateInit,
	}
	return &validator
}

type cardDataTypeState int

const cardDataTypeCommon = "https://github.com/StasMerzlyakov/gxml/common-data"

var cardDataTypeElementData1 = xsd.ElementData{
	Namespace: cardDataTypeCommon,
	Name:      "CardType",
	Type:      xsd.ElementNode,
}

var cardDataTypeElementData2 = xsd.ElementData{
	Namespace: cardDataTypeCommon,
	Name:      "CardNumberType",
	Type:      xsd.ElementNode,
}

var cardDataTypeElementData3 = xsd.ElementData{
	Namespace: cardDataTypeCommon,
	Name:      "ExpirationDateType",
	Type:      xsd.ElementNode,
}

var cardDataTypeStateToElement = map[cardDataTypeState]xsd.ElementData{
	cardDataTypeState1: cardDataTypeElementData1,
	cardDataTypeState2: cardDataTypeElementData2,
	cardDataTypeState3: cardDataTypeElementData3,
}

var cardDataTypeElementToState = map[xsd.ElementData]cardDataTypeState{
	cardDataTypeElementData1: cardDataTypeState1,
	cardDataTypeElementData2: cardDataTypeState2,
	cardDataTypeElementData3: cardDataTypeState3,
}

var cardDataTypeStateAcceptableMap = map[cardDataTypeState][]cardDataTypeState{
	cardDataTypeStateInit: {cardDataTypeState1},
	cardDataTypeState1:    {cardDataTypeState2},
	cardDataTypeState2:    {cardDataTypeState3},
	cardDataTypeState3:    {cardDataTypeStateEnd},
}

const (
	cardDataTypeStateInit cardDataTypeState = 0
	cardDataTypeState1    cardDataTypeState = 1
	cardDataTypeState2    cardDataTypeState = 2
	cardDataTypeState3    cardDataTypeState = 3
	cardDataTypeStateEnd  cardDataTypeState = 4
)
