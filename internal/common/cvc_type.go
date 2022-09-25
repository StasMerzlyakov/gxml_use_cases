package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type CvcTypeValidator struct {
	state cvcTypeState
}

func (cv *CvcTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *CvcTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *CvcTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *CvcTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *CvcTypeValidator) CompleteElement() error {
	acceptableStates := cvcTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cvcTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *CvcTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

type cvcTypeState int

var cvcTypeStateAcceptableMap = map[cvcTypeState][]cvcTypeState{
	cvcTypeStateInit: {cvcTypeStateEnd},
}

const (
	cvcTypeStateInit cvcTypeState = 0
	cvcTypeStateEnd  cvcTypeState = 1
)
