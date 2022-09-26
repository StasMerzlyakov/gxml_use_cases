package crd

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type cardResponseTypeValidator struct {
	state cardResponseTypeState
}

func (cv *cardResponseTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("unexpected element %s: expected", elementType.ToString())
	expectedStates := cardResponseTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := cardResponseTypeStateToElement[est]
		if idx == 0 {
			result += expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *cardResponseTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element: expected "
	expectedStates := cardResponseTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := cardResponseTypeStateToElement[est]
		if idx == 0 {
			result += " " + expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *cardResponseTypeValidator) CheckValue(runes []rune) error {
	if !util.IsEmpty(runes) {
		return errors.New("value unexpected")
	}
	return nil
}

func (cv *cardResponseTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type == xsd.CharData {
		return nil
	}
	if state, ok := cardResponseTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := cardResponseTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
}

func (cv *cardResponseTypeValidator) CompleteElement() error {
	// Проверка достижимости конечного состояния из текущего
	acceptableStates := cardResponseTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardResponseTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *cardResponseTypeValidator) GetInstance() (any, error) {
	return NewCardResponseType(), nil
}

func (cv *cardResponseTypeValidator) IsComplexType() bool {
	return true
}

type cardResponseTypeState int

var cardResponseTypeElementData1 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CardData",
	Type:      xsd.ElementNode,
}

var cardResponseTypeElementData2 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CVC",
	Type:      xsd.ElementNode,
}

func (cv *cardResponseTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	switch elementData {
	case cardResponseTypeElementData1:
		validator1 := common.CardDataTypeValidator{}
		return &validator1
	case cardResponseTypeElementData2:
		validator2 := common.CvcTypeValidator{}
		return &validator2
	default:
		return nil
	}
}

var cardResponseTypeStateToElement = map[cardResponseTypeState]xsd.ElementData{
	cardResponseTypeState1: cardResponseTypeElementData1,
	cardResponseTypeState2: cardResponseTypeElementData2,
}

var cardResponseTypeElementToState = map[xsd.ElementData]cardResponseTypeState{
	cardResponseTypeElementData1: cardResponseTypeState1,
	cardResponseTypeElementData2: cardResponseTypeState2,
}

var cardResponseTypeStateAcceptableMap = map[cardResponseTypeState][]cardResponseTypeState{
	cardResponseTypeStateInit: {cardResponseTypeState1},
	cardResponseTypeState1:    {cardResponseTypeState2},
	cardResponseTypeState2:    {cardResponseTypeStateEnd},
}

const (
	cardResponseTypeStateInit cardResponseTypeState = 0
	cardResponseTypeState1    cardResponseTypeState = 1
	cardResponseTypeState2    cardResponseTypeState = 2
	cardResponseTypeStateEnd  cardResponseTypeState = 3
)
