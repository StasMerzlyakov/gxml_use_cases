package common

import (
	"github.com/StasMerzlyakov/gxml/xsd"
)

type ComplexDateTypeStruct struct {
	Year  xsd.GYear
	Month xsd.GMonth
	Day   xsd.GDay
}

type ComplexDateType interface {
	GetYear() xsd.GYear
	SetYear(value xsd.GYear)
	GetMonth() xsd.GMonth
	SetMonth(value xsd.GMonth)
	GetDay() xsd.GDay
	SetDay(value xsd.GDay)
}

func (sv *ComplexDateTypeStruct) GetYear() xsd.GYear {
	return sv.Year
}

func (sv *ComplexDateTypeStruct) SetYear(value xsd.GYear) {
	sv.Year = value
}

func (sv *ComplexDateTypeStruct) GetMonth() xsd.GMonth {
	return sv.Month
}

func (sv *ComplexDateTypeStruct) SetMonth(value xsd.GMonth) {
	sv.Month = value
}

func (sv *ComplexDateTypeStruct) GetDay() xsd.GDay {
	return sv.Day
}

func (sv *ComplexDateTypeStruct) SetDay(value xsd.GDay) {
	sv.Day = value
}

func NewComplexDateType() ComplexDateType {
	res := ComplexDateTypeStruct{}
	return &res
}
