package crd

import (
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type CardRequestTypeStruct struct {
	FirstName  xsd2.String
	LastName   xsd2.String
	Patronymic xsd2.String
	BirthDate  common.ComplexDateType
	CardType   xsd2.String
}

type CardRequestType interface {
	GetFirstName() xsd2.String
	SetFirstName(value xsd2.String)
	GetLastName() xsd2.String
	SetLastName(value xsd2.String)
	GetPatronymic() xsd2.String
	SetPatronymic(value xsd2.String)
	GetBirthDate() common.ComplexDateType
	SetBirthDate(value common.ComplexDateType)
}

func (cv *CardRequestTypeStruct) GetFirstName() xsd2.String {
	return cv.FirstName
}

func (cv *CardRequestTypeStruct) SetFirstName(value xsd2.String) {
	cv.FirstName = value
}

func (cv *CardRequestTypeStruct) GetLastName() xsd2.String {
	return cv.LastName
}

func (cv *CardRequestTypeStruct) SetLastName(value xsd2.String) {
	cv.LastName = value
}

func (cv *CardRequestTypeStruct) GetPatronymic() xsd2.String {
	return cv.Patronymic
}

func (cv *CardRequestTypeStruct) SetPatronymic(value xsd2.String) {
	cv.Patronymic = value
}

func (cv *CardRequestTypeStruct) GetBirthDate() common.ComplexDateType {
	return cv.BirthDate
}

func (cv *CardRequestTypeStruct) SetBirthDate(value common.ComplexDateType) {
	cv.BirthDate = value
}

func (cv *CardRequestTypeStruct) GetCardType() xsd2.String {
	return cv.CardType
}

func (cv *CardRequestTypeStruct) SetCardType(value xsd2.String) {
	cv.CardType = value
}

func NewCardRequestType() CardRequestType {
	res := CardRequestTypeStruct{}
	return &res
}
