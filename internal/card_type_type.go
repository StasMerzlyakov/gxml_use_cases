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
	result := "Unexpected end of element expected. Expected value."
	return errors.New(result)
}

func (cv *cardTypeTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	if state, ok := cardTypeTypeElementToState[elementType]; !ok {
		return cv.unexpectedElementError(elementType)
	} else {
		acceptableStates := cardTypeTypeStateAcceptableMap[cv.state]
		if acceptableStates == nil || !util.Contains(acceptableStates, state) {
			return cv.unexpectedElementError(elementType)
		}
		cv.state = state
		return nil
	}
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

const cardTypeTypeCommon = "https://github.com/StasMerzlyakov/gxml/common-data"

var cardTypeTypeElementData1 = xsd.ElementData{
	Type: xsd.CharData,
}

var cardTypeTypeElementToState = map[xsd.ElementData]cardTypeTypeState{
	cardTypeTypeElementData1: cardTypeTypeValueState1,
}

var cardTypeTypeStateAcceptableMap = map[cardTypeTypeState][]cardTypeTypeState{
	cardTypeTypeStateInit:   {cardTypeTypeValueState1},
	cardTypeTypeValueState1: {cardTypeTypeStateEnd},
}

const (
	cardTypeTypeStateInit   cardTypeTypeState = 0
	cardTypeTypeValueState1 cardTypeTypeState = 1
	cardTypeTypeStateEnd    cardTypeTypeState = 2
)
