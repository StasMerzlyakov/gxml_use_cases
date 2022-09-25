package common

import "github.com/StasMerzlyakov/gxml_use_cases/xsd2"

type CardDataTypeStruct struct {
	CardType       xsd2.String
	CardNumber     xsd2.String
	ExpirationDate xsd2.String
}

type CardDataType interface {
	GetCardType() xsd2.String
	SetCardType(value xsd2.String)
	GetCardNumber() xsd2.String
	SetCardNumber(value xsd2.String)
	GetExpirationDate() xsd2.String
	SetExpirationDate(value xsd2.String)
}

func (st *CardDataTypeStruct) GetCardType() xsd2.String {
	return st.CardType
}

func (st *CardDataTypeStruct) SetCardType(value xsd2.String) {
	st.CardType = value
}

func (st *CardDataTypeStruct) GetCardNumber() xsd2.String {
	return st.CardNumber
}

func (st *CardDataTypeStruct) SetCardNumber(value xsd2.String) {
	st.CardNumber = value
}

func (st *CardDataTypeStruct) GetExpirationDate() xsd2.String {
	return st.ExpirationDate
}

func (st *CardDataTypeStruct) SetExpirationDate(value xsd2.String) {
	st.ExpirationDate = value
}

func NewCardDataType() CardDataType {
	res := CardDataTypeStruct{}
	return &res
}
