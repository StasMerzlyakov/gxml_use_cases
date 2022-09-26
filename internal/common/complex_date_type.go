package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type ComplexDateTypeValidator struct {
	state complexDateTypeState
}

func (cv *ComplexDateTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("unexpected element %s: expected", elementType.ToString())
	expectedStates := complexDateTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := complexDateTypeStateToElement[est]
		if idx == 0 {
			result += expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *ComplexDateTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element: expected "
	expectedStates := complexDateTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := complexDateTypeStateToElement[est]
		if idx == 0 {
			result += " " + expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *ComplexDateTypeValidator) CheckValue(runes []rune) error {
	if !util.IsEmpty(runes) {
		return errors.New("value unexpected")
	}
	return nil
}

func (cv *ComplexDateTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type == xsd.CharData {
		return nil
	}
	if state, ok := complexDateTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := complexDateTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
}

func (cv *ComplexDateTypeValidator) CompleteElement() error {
	// Проверка достижимости конечного состояния из текущего
	acceptableStates := complexDateTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, complexDateTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type complexDateTypeState int

var complexDateTypeElementData1 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "Year",
	Type:      xsd.ElementNode,
}

var complexDateTypeElementData2 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "Month",
	Type:      xsd.ElementNode,
}

var complexDateTypeElementData3 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "Day",
	Type:      xsd.ElementNode,
}

func (cv *ComplexDateTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	switch elementData {
	case complexDateTypeElementData1:
		validator1 := ComplexDateTypeYearTypeValidator{}
		return &validator1
	case complexDateTypeElementData2:
		validator2 := ComplexDateTypeMonthTypeValidator{}
		return &validator2
	case complexDateTypeElementData3:
		validator3 := ComplexDateTypeDayTypeValidator{}
		return &validator3
	default:
		return nil
	}
}

func (cv *ComplexDateTypeValidator) GetInstance() (any, error) {
	return NewComplexDateType(), nil
}

func (cv *ComplexDateTypeValidator) IsComplexType() bool {
	return true
}

var complexDateTypeStateToElement = map[complexDateTypeState]xsd.ElementData{
	complexDateTypeState1: complexDateTypeElementData1,
	complexDateTypeState2: complexDateTypeElementData2,
	complexDateTypeState3: complexDateTypeElementData3,
}

var complexDateTypeElementToState = map[xsd.ElementData]complexDateTypeState{
	complexDateTypeElementData1: complexDateTypeState1,
	complexDateTypeElementData2: complexDateTypeState2,
	complexDateTypeElementData3: complexDateTypeState3,
}

var complexDateTypeStateAcceptableMap = map[complexDateTypeState][]complexDateTypeState{
	complexDateTypeStateInit: {complexDateTypeState1},
	complexDateTypeState1:    {complexDateTypeState2, complexDateTypeState3, complexDateTypeStateEnd},
	complexDateTypeState2:    {complexDateTypeState3, complexDateTypeStateEnd},
	complexDateTypeState3:    {complexDateTypeStateEnd},
}

const (
	complexDateTypeStateInit complexDateTypeState = 0
	complexDateTypeState1    complexDateTypeState = 1
	complexDateTypeState2    complexDateTypeState = 2
	complexDateTypeState3    complexDateTypeState = 3
	complexDateTypeStateEnd  complexDateTypeState = 4
)
