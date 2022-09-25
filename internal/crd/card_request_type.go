package crd

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
)

type cardRequestTypeValidator struct {
	state cardRequestTypeState
}

func (cv *cardRequestTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("unexpected element %s: expected", elementType.ToString())
	expectedStates := cardRequestTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := cardRequestTypeStateToElement[est]
		if idx == 0 {
			result += expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *cardRequestTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element: expected "
	expectedStates := cardRequestTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := cardRequestTypeStateToElement[est]
		if idx == 0 {
			result += " " + expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *cardRequestTypeValidator) CheckValue(runes []rune) error {
	if !util.IsEmpty(runes) {
		return errors.New("value unexpected")
	}
	return nil
}

func (cv *cardRequestTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type == xsd.CharData {
		return nil
	}
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

func (cv *cardRequestTypeValidator) CompleteElement() error {
	// Проверка достижимости конечного состояния из текущего
	acceptableStates := cardRequestTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardRequestTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type cardRequestTypeState int

var cardRequestTypeElementData1 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "FirstName",
	Type:      xsd.ElementNode,
}

var cardRequestTypeElementData2 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "LastName",
	Type:      xsd.ElementNode,
}

var cardRequestTypeElementData3 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "Patronymic",
	Type:      xsd.ElementNode,
}

var cardRequestTypeElementData4 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "BirthDate",
	Type:      xsd.ElementNode,
}

var cardRequestTypeElementData5 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
	Name:      "CardType",
	Type:      xsd.ElementNode,
}

func (cv *cardRequestTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	switch elementData {
	case cardRequestTypeElementData1:
		validator1 := common.MiddleStringTypeValidator{}
		return &validator1
	case cardRequestTypeElementData2:
		validator2 := common.MiddleStringTypeValidator{}
		return &validator2
	case cardRequestTypeElementData3:
		validator3 := common.MiddleStringTypeValidator{}
		return &validator3
	case cardRequestTypeElementData4:
		nameAndNamespace := xsd.NameAndNamespace{
			Namespace: elementData.Namespace,
			Name:      elementData.Name,
		}
		validator4 := common.ResolveValidator(nameAndNamespace)
		return validator4
	case cardRequestTypeElementData5:
		validator5 := common.CardTypeTypeValidator{}
		return &validator5
	default:
		return nil
	}
}

var cardRequestTypeStateToElement = map[cardRequestTypeState]xsd.ElementData{
	cardRequestTypeState1: cardRequestTypeElementData1,
	cardRequestTypeState2: cardRequestTypeElementData2,
	cardRequestTypeState3: cardRequestTypeElementData3,
	cardRequestTypeState4: cardRequestTypeElementData4,
	cardRequestTypeState5: cardRequestTypeElementData5,
}

var cardRequestTypeElementToState = map[xsd.ElementData]cardRequestTypeState{
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
	cardRequestTypeState5:    {cardRequestTypeStateEnd},
}

const (
	cardRequestTypeStateInit cardRequestTypeState = 0
	cardRequestTypeState1    cardRequestTypeState = 1
	cardRequestTypeState2    cardRequestTypeState = 2
	cardRequestTypeState3    cardRequestTypeState = 3
	cardRequestTypeState4    cardRequestTypeState = 4
	cardRequestTypeState5    cardRequestTypeState = 5
	cardRequestTypeStateEnd  cardRequestTypeState = 6
)
