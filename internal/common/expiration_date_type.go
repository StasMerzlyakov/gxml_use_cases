package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type ExpirationDateTypeValidator struct {
	state expirationDateTypeState
}

func (cv *ExpirationDateTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *ExpirationDateTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *ExpirationDateTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *ExpirationDateTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *ExpirationDateTypeValidator) CompleteElement() error {
	acceptableStates := expirationDateTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, expirationDateTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *ExpirationDateTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

type expirationDateTypeState int

var expirationDateTypeStateAcceptableMap = map[expirationDateTypeState][]expirationDateTypeState{
	expirationDateTypeStateInit: {expirationDateTypeStateEnd},
}

const (
	expirationDateTypeStateInit expirationDateTypeState = 0
	expirationDateTypeStateEnd  expirationDateTypeState = 1
)
