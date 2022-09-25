package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
)

type DateTimeTypeValidator struct {
	state dateTimeTypeState
}

func (cv *DateTimeTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *DateTimeTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *DateTimeTypeValidator) CheckValue(runes []rune) error {
	return nil
}

func (cv *DateTimeTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *DateTimeTypeValidator) CompleteElement() error {
	acceptableStates := dateTimeTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, dateTimeTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *DateTimeTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

type dateTimeTypeState int

var dateTimeTypeStateAcceptableMap = map[dateTimeTypeState][]dateTimeTypeState{
	dateTimeTypeStateInit: {dateTimeTypeStateEnd},
}

const (
	dateTimeTypeStateInit dateTimeTypeState = 0
	dateTimeTypeStateEnd  dateTimeTypeState = 1
)
