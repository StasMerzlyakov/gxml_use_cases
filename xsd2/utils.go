package xsd2

import (
	"fmt"
	"github.com/StasMerzlyakov/gxml/xsd"
)

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

func ReverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

type NameAndNamespace struct {
	Namespace string
	Name      string
}

func (n NameAndNamespace) ToString() string {
	return fmt.Sprintf("{%s}%s", n.Namespace, n.Name)
}
