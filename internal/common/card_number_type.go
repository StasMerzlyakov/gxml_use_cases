package common

import (
	"errors"
	"fmt"
	"github.com/StasMerzlyakov/gxml/util"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
	"strings"
)

type CardNumberTypeValidator struct {
	state cardNumberTypeState
	sb    strings.Builder
}

func (cv *CardNumberTypeValidator) unexpectedElementError(elementType xsd.ElementData) error {
	result := fmt.Sprintf("Unexpected element %s; expected element value.", elementType.ToString())
	return errors.New(result)
}

func (cv *CardNumberTypeValidator) unexpectedEndOfElement() error {
	result := "unexpected end of element expected: expected value"
	return errors.New(result)
}

func (cv *CardNumberTypeValidator) CheckValue(runes []rune) error {
	cv.sb.WriteString(string(runes))
	return nil
}

func (cv *CardNumberTypeValidator) AcceptElement(elementType xsd.ElementData) error {
	if elementType.Type != xsd.CharData {
		return cv.unexpectedElementError(elementType)
	}
	return nil
}

func (cv *CardNumberTypeValidator) CompleteElement() error {
	acceptableStates := cardNumberTypeStateAcceptableMap[cv.state]
	if util.Contains(acceptableStates, cardNumberTypeStateEnd) {
		return nil
	} else {
		return cv.unexpectedEndOfElement()
	}
}

func (cv *CardNumberTypeValidator) ResolveValidator(elementData xsd.ElementData) xsd2.IElementValidator {
	return nil
}

func (cv *CardNumberTypeValidator) GetInstance() (any, error) {
	return xsd2.NewString(cv.sb.String())
}

func (cv *CardNumberTypeValidator) IsComplexType() bool {
	return false
}

type cardNumberTypeState int

var cardNumberTypeStateAcceptableMap = map[cardNumberTypeState][]cardNumberTypeState{
	cardNumberTypeStateInit: {cardNumberTypeStateEnd},
}

const (
	cardNumberTypeStateInit cardNumberTypeState = 0
	cardNumberTypeStateEnd  cardNumberTypeState = 1
)
