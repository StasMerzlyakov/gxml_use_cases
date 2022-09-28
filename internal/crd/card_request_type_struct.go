package crd

import (
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
)

type CardRequestTypeStruct struct {
	FirstName  xsd.String
	LastName   xsd.String
	Patronymic xsd.String
	BirthDate  common.ComplexDateType
	CardType   xsd.String
}

type CardRequestType interface {
	GetFirstName() xsd.String
	SetFirstName(value xsd.String)
	GetLastName() xsd.String
	SetLastName(value xsd.String)
	GetPatronymic() xsd.String
	SetPatronymic(value xsd.String)
	GetBirthDate() common.ComplexDateType
	SetBirthDate(value common.ComplexDateType)
}

func (cv *CardRequestTypeStruct) GetFirstName() xsd.String {
	return cv.FirstName
}

func (cv *CardRequestTypeStruct) SetFirstName(value xsd.String) {
	cv.FirstName = value
}

func (cv *CardRequestTypeStruct) GetLastName() xsd.String {
	return cv.LastName
}

func (cv *CardRequestTypeStruct) SetLastName(value xsd.String) {
	cv.LastName = value
}

func (cv *CardRequestTypeStruct) GetPatronymic() xsd.String {
	return cv.Patronymic
}

func (cv *CardRequestTypeStruct) SetPatronymic(value xsd.String) {
	cv.Patronymic = value
}

func (cv *CardRequestTypeStruct) GetBirthDate() common.ComplexDateType {
	return cv.BirthDate
}

func (cv *CardRequestTypeStruct) SetBirthDate(value common.ComplexDateType) {
	cv.BirthDate = value
}

func (cv *CardRequestTypeStruct) GetCardType() xsd.String {
	return cv.CardType
}

func (cv *CardRequestTypeStruct) SetCardType(value xsd.String) {
	cv.CardType = value
}

func NewCardRequestType() CardRequestType {
	res := CardRequestTypeStruct{}
	return &res
}
