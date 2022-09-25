package common

import "github.com/StasMerzlyakov/gxml_use_cases/xsd2"

type ComplexDateTypeStruct struct {
	Year  xsd2.GYear
	Month xsd2.GMonth
	Day   xsd2.GDay
}

type ComplexDateType interface {
	GetYear() xsd2.GYear
	SetYear(value xsd2.GYear)
	GetMonth() xsd2.GMonth
	SetMonth(value xsd2.GMonth)
	GetDay() xsd2.GDay
	SetDay(value xsd2.GDay)
}

func (sv *ComplexDateTypeStruct) GetYear() xsd2.GYear {
	return sv.Year
}

func (sv *ComplexDateTypeStruct) SetYear(value xsd2.GYear) {
	sv.Year = value
}

func (sv *ComplexDateTypeStruct) GetMonth() xsd2.GMonth {
	return sv.Month
}

func (sv *ComplexDateTypeStruct) SetMonth(value xsd2.GMonth) {
	sv.Month = value
}

func (sv *ComplexDateTypeStruct) GetDay() xsd2.GDay {
	return sv.Day
}

func (sv *ComplexDateTypeStruct) SetDay(value xsd2.GDay) {
	sv.Day = value
}

func NewComplexDateType() ComplexDateType {
	res := ComplexDateTypeStruct{}
	return &res
}
