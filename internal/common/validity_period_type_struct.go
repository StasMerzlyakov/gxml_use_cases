package common

import "github.com/StasMerzlyakov/gxml_use_cases/xsd2"

type ValidityPeriodTypeStruct struct {
	From xsd2.DateTime
	To   xsd2.DateTime
}

type ValidityPeriodType interface {
	GetFrom() xsd2.DateTime
	SetFrom(value xsd2.DateTime)
	GetTo() xsd2.DateTime
	SetTo(value xsd2.DateTime)
}

func (cv *ValidityPeriodTypeStruct) GetFrom() xsd2.DateTime {
	return cv.From
}

func (cv *ValidityPeriodTypeStruct) GetTo() xsd2.DateTime {
	return cv.To
}

func (cv *ValidityPeriodTypeStruct) SetFrom(value xsd2.DateTime) {
	cv.From = value
}

func (cv *ValidityPeriodTypeStruct) SetTo(value xsd2.DateTime) {
	cv.To = value
}

func NewValidityPeriodType() ValidityPeriodType {
	res := ValidityPeriodTypeStruct{}
	return &res
}
