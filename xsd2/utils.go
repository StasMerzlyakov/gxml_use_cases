package xsd2

import "github.com/StasMerzlyakov/gxml/xsd"

type ElementValidatorStack struct {
	data []xsd.IElementValidator
}

func (s *ElementValidatorStack) Pop() (v xsd.IElementValidator) {
	l := len(s.data)
	rv := s.data[l-1]
	s.data = s.data[:l-1]
	return rv
}

func (s *ElementValidatorStack) Push(v xsd.IElementValidator) {
	s.data = append(s.data, v)
}

func (s *ElementValidatorStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *ElementValidatorStack) Peek() (v xsd.IElementValidator) {
	l := len(s.data)
	return s.data[l-1]
}
