package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type MiddleStringTypeValidator struct {
	state middleStringTypeState
}

func (cv *MiddleStringTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *MiddleStringTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *MiddleStringTypeValidator) CheckValue(runes []rune) (any, error) {
	return xsd2.NewString(string(runes))
}

func (cv *MiddleStringTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *MiddleStringTypeValidator) ResolveValidator(elementData xsd.ElementData) (any, xsd2.IElementValidator) {
	return nil, nil
}

func (cv *MiddleStringTypeValidator) CompleteElement() (bool, error) {
	acceptableStates := middleStringTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, middleStringTypeStateEnd) {
		return false, nil
	} else {
		return false, cv.unexpectedEndOfElement()
	}
}

type middleStringTypeState int

var middleStringTypeStateAcceptableMap = map[middleStringTypeState][]middleStringTypeState{
	middleStringTypeStateInit: {middleStringTypeStateEnd},
}

const (
	middleStringTypeStateInit middleStringTypeState = 0
	middleStringTypeStateEnd  middleStringTypeState = 1
)
