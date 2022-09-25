package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type CardDataTypeValidator struct {
	state cardDataTypeState
}

func (cv *CardDataTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
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

func (cv *CardDataTypeValidator) unexpectedEndOfElement() error {
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

func (cv *CardDataTypeValidator) CheckValue(runes []rune) error {
	if !util.IsEmpty(runes) {
		return errors.New("value unexpected")
	}
	return nil
}

func (cv *CardDataTypeValidator) AcceptElement(elementType xsd.ElementData) error {
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

func (cv *CardDataTypeValidator) CompleteElement() error {
	// Проверка достижимости конечного состояния из текущего
	acceptableStates := cardDataTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardDataTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type cardDataTypeState int

var cardDataTypeElementData1 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "CardType",
	Type:      xsd.ElementNode,
}

var cardDataTypeElementData2 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "CardNumber",
	Type:      xsd.ElementNode,
}

var cardDataTypeElementData3 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "ExpirationDate",
	Type:      xsd.ElementNode,
}

func (cv *CardDataTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	switch elementData {
	case cardDataTypeElementData1:
		validator1 := ComplexDateTypeYearTypeValidator{}
		return &validator1
	case cardDataTypeElementData2:
		validator2 := ComplexDateTypeMonthTypeValidator{}
		return &validator2
	case cardDataTypeElementData3:
		validator3 := ComplexDateTypeDayTypeValidator{}
		return &validator3
	default:
		return nil
	}
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
