package crd

import (
	"fmt"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type sample struct {
}

func (s sample) AcceptElement(elementData xsd.ElementData) error {
	return nil
}

func (s sample) CompleteElement() error {
	fmt.Println("crd !!")
	return nil
}

func (s sample) CheckValue(runes []rune) []string {
	return nil
}

func ResolveValidator(name xsd2.NameAndNamespace) (xsd.IElementValidator, error) {
	validator := sample{}
	return &validator, nil
}
