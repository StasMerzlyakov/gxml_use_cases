package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"strings"
)

type MiddleStringTypeValidator struct {
	state middleStringTypeState
	sb    strings.Builder
}

func (cv *MiddleStringTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *MiddleStringTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *MiddleStringTypeValidator) CheckValue(runes []rune) error {
	cv.sb.WriteString(string(runes))
	return nil
}

func (cv *MiddleStringTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *MiddleStringTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd.IElementValidator {
	return nil
}

func (cv *MiddleStringTypeValidator) CompleteElement() error {
	acceptableStates := middleStringTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, middleStringTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *MiddleStringTypeValidator) GetInstance() (any, error) {
	return xsd.NewString(cv.sb.String())
}

func (cv *MiddleStringTypeValidator) IsComplexType() bool {
	return false
}

func (cv *MiddleStringTypeValidator) GetStates() []xsd.ElementData {
	return nil
}

type middleStringTypeState int

var middleStringTypeStateAcceptableMap = map[middleStringTypeState][]middleStringTypeState{
	middleStringTypeStateInit: {middleStringTypeStateEnd},
}

const (
	middleStringTypeStateInit middleStringTypeState = 0
	middleStringTypeStateEnd  middleStringTypeState = 1
)
