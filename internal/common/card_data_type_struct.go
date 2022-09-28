package common

import (
	"github.com/StasMerzlyakov/gxml/xsd"
)

type CardDataTypeStruct struct {
	CardType       xsd.String
	CardNumber     xsd.String
	ExpirationDate xsd.String
}

type CardDataType interface {
	GetCardType() xsd.String
	SetCardType(value xsd.String)
	GetCardNumber() xsd.String
	SetCardNumber(value xsd.String)
	GetExpirationDate() xsd.String
	SetExpirationDate(value xsd.String)
}

func (st *CardDataTypeStruct) GetCardType() xsd.String {
	return st.CardType
}

func (st *CardDataTypeStruct) SetCardType(value xsd.String) {
	st.CardType = value
}

func (st *CardDataTypeStruct) GetCardNumber() xsd.String {
	return st.CardNumber
}

func (st *CardDataTypeStruct) SetCardNumber(value xsd.String) {
	st.CardNumber = value
}

func (st *CardDataTypeStruct) GetExpirationDate() xsd.String {
	return st.ExpirationDate
}

func (st *CardDataTypeStruct) SetExpirationDate(value xsd.String) {
	st.ExpirationDate = value
}

func NewCardDataType() CardDataType {
	res := CardDataTypeStruct{}
	return &res
}
