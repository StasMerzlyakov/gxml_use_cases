package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
	"strings"
)

type DateTimeTypeValidator struct {
	state dateTimeTypeState
	sb    strings.Builder
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
	cv.sb.WriteString(string(runes))
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

func (cv *DateTimeTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	return nil
}

func (cv *DateTimeTypeValidator) GetInstance() (any, error) {
	return xsd2.NewDateTime(cv.sb.String())
}

func (cv *DateTimeTypeValidator) IsComplexType() bool {
	return false
}

func (cv *DateTimeTypeValidator) GetStates() []xsd.ElementData {
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
