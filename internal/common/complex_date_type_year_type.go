package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type ComplexDateTypeYearTypeValidator struct {
	state complexDateTypeYearTypeState
}

func (cv *ComplexDateTypeYearTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *ComplexDateTypeYearTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *ComplexDateTypeYearTypeValidator) CheckValue(runes []rune) (any, error) {
	return xsd2.NewGYear(string(runes))
}

func (cv *ComplexDateTypeYearTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *ComplexDateTypeYearTypeValidator) CompleteElement() (bool, error) {
	acceptableStates := complexDateTypeYearTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, complexDateTypeYearTypeStateEnd) {
		return false, nil
	} else {
		return false, cv.unexpectedEndOfElement()
	}
}

func (cv *ComplexDateTypeYearTypeValidator) ResolveValidator(elementData xsd.ElementData) (any, xsd2.IElementValidator) {
	return nil, nil
}

type complexDateTypeYearTypeState int

var complexDateTypeYearTypeStateAcceptableMap = map[complexDateTypeYearTypeState][]complexDateTypeYearTypeState{
	complexDateTypeYearTypeStateInit: {complexDateTypeYearTypeStateEnd},
}

const (
	complexDateTypeYearTypeStateInit complexDateTypeYearTypeState = 0
	complexDateTypeYearTypeStateEnd  complexDateTypeYearTypeState = 1
)
