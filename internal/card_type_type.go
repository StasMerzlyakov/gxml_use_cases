package internal

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type cardTypeTypeValidator struct {
	state cardTypeTypeState
}

func (cv *cardTypeTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *cardTypeTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *cardTypeTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *cardTypeTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *cardTypeTypeValidator) CompleteElement() error {
	acceptableStates := cardTypeTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardTypeTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type cardTypeTypeCreator struct {
}

func (cardTypeTypeCreator) Create() xsd.IElementValidator {
	validator := cardTypeTypeValidator{
		state: cardTypeTypeStateInit,
	}
	return &validator
}

type cardTypeTypeState int

var cardTypeTypeStateAcceptableMap = map[cardTypeTypeState][]cardTypeTypeState{
	cardTypeTypeStateInit: {cardTypeTypeStateEnd},
}

const (
	cardTypeTypeStateInit cardTypeTypeState = 0
	cardTypeTypeStateEnd  cardTypeTypeState = 1
)
