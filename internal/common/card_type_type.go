package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type CardTypeTypeValidator struct {
	state cardTypeTypeState
}

func (cv *CardTypeTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *CardTypeTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *CardTypeTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *CardTypeTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *CardTypeTypeValidator) CompleteElement() error {
	acceptableStates := cardTypeTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardTypeTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *CardTypeTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

type cardTypeTypeState int

var cardTypeTypeStateAcceptableMap = map[cardTypeTypeState][]cardTypeTypeState{
	cardTypeTypeStateInit: {cardTypeTypeStateEnd},
}

const (
	cardTypeTypeStateInit cardTypeTypeState = 0
	cardTypeTypeStateEnd  cardTypeTypeState = 1
)
