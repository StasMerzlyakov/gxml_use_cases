package internal

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type cardNumberTypeValidator struct {
	state cardNumberTypeState
}

func (cv *cardNumberTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *cardNumberTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *cardNumberTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *cardNumberTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *cardNumberTypeValidator) CompleteElement() error {
	acceptableStates := cardNumberTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardNumberTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

type cardNumberTypeCreator struct {
}

func (cardNumberTypeCreator) Create() xsd.IElementValidator {
	validator := cardNumberTypeValidator{
		state: cardNumberTypeStateInit,
	}
	return &validator
}

type cardNumberTypeState int

var cardNumberTypeStateAcceptableMap = map[cardNumberTypeState][]cardNumberTypeState{
	cardNumberTypeStateInit: {cardNumberTypeStateEnd},
}

const (
	cardNumberTypeStateInit cardNumberTypeState = 0
	cardNumberTypeStateEnd  cardNumberTypeState = 1
)
