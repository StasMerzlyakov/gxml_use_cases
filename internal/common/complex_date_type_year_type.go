package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
	"strings"
)

type ComplexDateTypeYearTypeValidator struct {
	state complexDateTypeYearTypeState
	sb    strings.Builder
}

func (cv *ComplexDateTypeYearTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *ComplexDateTypeYearTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *ComplexDateTypeYearTypeValidator) CheckValue(runes []rune) error {
	cv.sb.WriteString(string(runes))
	return nil
}

func (cv *ComplexDateTypeYearTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *ComplexDateTypeYearTypeValidator) CompleteElement() error {
	acceptableStates := complexDateTypeYearTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, complexDateTypeYearTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *ComplexDateTypeYearTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	return nil
}

func (cv *ComplexDateTypeYearTypeValidator) GetInstance() (any, error) {
	return xsd2.NewGYear(cv.sb.String())
}

func (cv *ComplexDateTypeYearTypeValidator) IsComplexType() bool {
	return false
}

func (cv *ComplexDateTypeYearTypeValidator) GetStates() []xsd.ElementData {
	return nil
}

type complexDateTypeYearTypeState int

var complexDateTypeYearTypeStateAcceptableMap = map[complexDateTypeYearTypeState][]complexDateTypeYearTypeState{
	complexDateTypeYearTypeStateInit: {complexDateTypeYearTypeStateEnd},
}

const (
	complexDateTypeYearTypeStateInit complexDateTypeYearTypeState = 0
	complexDateTypeYearTypeStateEnd  complexDateTypeYearTypeState = 1
)
