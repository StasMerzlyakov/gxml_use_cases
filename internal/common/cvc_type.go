package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
	"strings"
)

type CvcTypeValidator struct {
	state cvcTypeState
	sb    strings.Builder
}

func (cv *CvcTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *CvcTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *CvcTypeValidator) CheckValue(runes []rune) error {
	cv.sb.WriteString(string(runes))
	return nil
}

func (cv *CvcTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *CvcTypeValidator) CompleteElement() error {
	acceptableStates := cvcTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cvcTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *CvcTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	return nil
}

func (cv *CvcTypeValidator) GetInstance() (any, error) {
	return xsd2.NewString(cv.sb.String())
}

func (cv *CvcTypeValidator) IsComplexType() bool {
	return false
}

func (cv *CvcTypeValidator) GetStates() []xsd.ElementData {
	return nil
}

type cvcTypeState int

var cvcTypeStateAcceptableMap = map[cvcTypeState][]cvcTypeState{
	cvcTypeStateInit: {cvcTypeStateEnd},
}

const (
	cvcTypeStateInit cvcTypeState = 0
	cvcTypeStateEnd  cvcTypeState = 1
)
