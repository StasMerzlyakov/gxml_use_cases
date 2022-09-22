package internal

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type expirationDateTypeValidator struct {
	state expirationDateTypeState
}

func (cv *expirationDateTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *expirationDateTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *expirationDateTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *expirationDateTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *expirationDateTypeValidator) CompleteElement() error {
	acceptableStates := expirationDateTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, expirationDateTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type expirationDateTypeCreator struct {
}

func (expirationDateTypeCreator) Create() xsd.IElementValidator {
	validator := expirationDateTypeValidator{
		state: expirationDateTypeStateInit,
	}
	return &validator
}

type expirationDateTypeState int

var expirationDateTypeStateAcceptableMap = map[expirationDateTypeState][]expirationDateTypeState{
	expirationDateTypeStateInit: {expirationDateTypeStateEnd},
}

const (
	expirationDateTypeStateInit expirationDateTypeState = 0
	expirationDateTypeStateEnd  expirationDateTypeState = 1
)
