package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"strings"
)

type ComplexDateTypeDayTypeValidator struct {
	state complexDateTypeDayTypeState
	sb    strings.Builder
}

func (cv *ComplexDateTypeDayTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *ComplexDateTypeDayTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *ComplexDateTypeDayTypeValidator) CheckValue(runes []rune) error {
	cv.sb.WriteString(string(runes))
	return nil
}

func (cv *ComplexDateTypeDayTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *ComplexDateTypeDayTypeValidator) CompleteElement() error {
	acceptableStates := complexDateTypeDayTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, complexDateTypeDayTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *ComplexDateTypeDayTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

func (cv *ComplexDateTypeDayTypeValidator) GetInstance() (any, error) {
	return xsd.NewGDay(cv.sb.String())
}

func (cv *ComplexDateTypeDayTypeValidator) IsComplexType() bool {
	return false
}

func (cv *ComplexDateTypeDayTypeValidator) GetStates() []xsd.ElementData {
	return nil
}

type complexDateTypeDayTypeState int

var complexDateTypeDayTypeStateAcceptableMap = map[complexDateTypeDayTypeState][]complexDateTypeDayTypeState{
	complexDateTypeDayTypeStateInit: {complexDateTypeDayTypeStateEnd},
}

const (
	complexDateTypeDayTypeStateInit complexDateTypeDayTypeState = 0
	complexDateTypeDayTypeStateEnd  complexDateTypeDayTypeState = 1
)
