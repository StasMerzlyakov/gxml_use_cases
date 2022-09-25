package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type ComplexDateTypeDayTypeValidator struct {
	state complexDateTypeDayTypeState
}

func (cv *ComplexDateTypeDayTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *ComplexDateTypeDayTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *ComplexDateTypeDayTypeValidator) CheckValue(runes []rune) (any, error) {
	return xsd2.NewGDay(string(runes))
}

func (cv *ComplexDateTypeDayTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *ComplexDateTypeDayTypeValidator) CompleteElement() (bool, error) {
	acceptableStates := complexDateTypeDayTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, complexDateTypeDayTypeStateEnd) {
		return false, nil
	} else {
		return false, cv.unexpectedEndOfElement()
	}
}

func (cv *ComplexDateTypeDayTypeValidator) ResolveValidator(elementData xsd.ElementData) (any, xsd2.IElementValidator) {
	return nil, nil
}

type complexDateTypeDayTypeState int

var complexDateTypeDayTypeStateAcceptableMap = map[complexDateTypeDayTypeState][]complexDateTypeDayTypeState{
	complexDateTypeDayTypeStateInit: {complexDateTypeDayTypeStateEnd},
}

const (
	complexDateTypeDayTypeStateInit complexDateTypeDayTypeState = 0
	complexDateTypeDayTypeStateEnd  complexDateTypeDayTypeState = 1
)
