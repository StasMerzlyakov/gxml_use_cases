package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type CardNumberTypeValidator struct {
	state cardNumberTypeState
}

func (cv *CardNumberTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *CardNumberTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *CardNumberTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *CardNumberTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *CardNumberTypeValidator) CompleteElement() error {
	acceptableStates := cardNumberTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardNumberTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *CardNumberTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

type cardNumberTypeState int

var cardNumberTypeStateAcceptableMap = map[cardNumberTypeState][]cardNumberTypeState{
	cardNumberTypeStateInit: {cardNumberTypeStateEnd},
}

const (
	cardNumberTypeStateInit cardNumberTypeState = 0
	cardNumberTypeStateEnd  cardNumberTypeState = 1
)
