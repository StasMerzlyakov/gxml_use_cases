package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
	"golang.org/x/exp/maps"
)

type ValidityPeriodTypeValidator struct {
	state validityPeriodTypeState
}

func (cv *ValidityPeriodTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("unexpected element %s: expected", elementType.ToString())
	expectedStates := validityPeriodTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := validityPeriodTypeStateToElement[est]
		if idx == 0 {
			result += expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *ValidityPeriodTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element: expected "
	expectedStates := validityPeriodTypeStateAcceptableMap[cv.state]
	for idx, est := range expectedStates {
		expectedElement := validityPeriodTypeStateToElement[est]
		if idx == 0 {
			result += " " + expectedElement.ToString()
		} else {
			result += ", " + expectedElement.ToString()
		}
	}
	return errors.New(result)
}

func (cv *ValidityPeriodTypeValidator) CheckValue(runes []rune) error {
	if !util.IsEmpty(runes) {
		return errors.New("value unexpected")
	}
	return nil
}

func (cv *ValidityPeriodTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type == xsd.CharData {
		return nil
	}
	if state, ok := validityPeriodTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := validityPeriodTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
}

func (cv *ValidityPeriodTypeValidator) CompleteElement() error {
	// Проверка достижимости конечного состояния из текущего
	acceptableStates := validityPeriodTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, validityPeriodTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type validityPeriodTypeState int

var validityPeriodTypeElementData1 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "From",
	Type:      xsd.ElementNode,
}

var validityPeriodTypeElementData2 = xsd.ElementData{
	Namespace: "https://github.com/StasMerzlyakov/gxml/common-data",
	Name:      "To",
	Type:      xsd.ElementNode,
}

func (cv *ValidityPeriodTypeValidator) GetStates() []xsd.ElementData {
	return maps.Values(validityPeriodTypeStateToElement)
}

func (cv *ValidityPeriodTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	switch elementData {
	case validityPeriodTypeElementData1:
		validator1 := DateTimeTypeValidator{}
		return &validator1
	case validityPeriodTypeElementData2:
		validator2 := DateTimeTypeValidator{}
		return &validator2
	default:
		return nil
	}
}

func (cv *ValidityPeriodTypeValidator) GetInstance() (any, error) {
	return NewValidityPeriodType(), nil
}

func (cv *ValidityPeriodTypeValidator) IsComplexType() bool {
	return true
}

var validityPeriodTypeStateToElement = map[validityPeriodTypeState]xsd.ElementData{
	validityPeriodTypeState1: validityPeriodTypeElementData1,
	validityPeriodTypeState2: validityPeriodTypeElementData2,
}

var validityPeriodTypeElementToState = map[xsd.ElementData]validityPeriodTypeState{
	validityPeriodTypeElementData1: validityPeriodTypeState1,
	validityPeriodTypeElementData2: validityPeriodTypeState2,
}

var validityPeriodTypeStateAcceptableMap = map[validityPeriodTypeState][]validityPeriodTypeState{
	validityPeriodTypeStateInit: {validityPeriodTypeState1},
	validityPeriodTypeState1:    {validityPeriodTypeState2},
	validityPeriodTypeState2:    {validityPeriodTypeStateEnd},
}

const (
	validityPeriodTypeStateInit validityPeriodTypeState = 0
	validityPeriodTypeState1    validityPeriodTypeState = 1
	validityPeriodTypeState2    validityPeriodTypeState = 2
	validityPeriodTypeStateEnd  validityPeriodTypeState = 3
)
