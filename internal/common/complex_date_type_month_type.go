package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type ComplexDateTypeMonthTypeValidator struct {
	state complexDateTypeMonthTypeState
}

func (cv *ComplexDateTypeMonthTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *ComplexDateTypeMonthTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *ComplexDateTypeMonthTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *ComplexDateTypeMonthTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *ComplexDateTypeMonthTypeValidator) CompleteElement() error {
	acceptableStates := complexMonthTypeYearTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, complexMonthTypeYearTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *ComplexDateTypeMonthTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

type complexDateTypeMonthTypeState int

var complexMonthTypeYearTypeStateAcceptableMap = map[complexDateTypeMonthTypeState][]complexDateTypeMonthTypeState{
	complexMonthTypeYearTypeStateInit: {complexMonthTypeYearTypeStateEnd},
}

const (
	complexMonthTypeYearTypeStateInit complexDateTypeMonthTypeState = 0
	complexMonthTypeYearTypeStateEnd  complexDateTypeMonthTypeState = 1
)