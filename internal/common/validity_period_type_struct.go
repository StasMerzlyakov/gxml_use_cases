package common

import (
	"github.com/StasMerzlyakov/gxml/xsd"
)

type ValidityPeriodTypeStruct struct {
	From xsd.DateTime
	To   xsd.DateTime
}

type ValidityPeriodType interface {
	GetFrom() xsd.DateTime
	SetFrom(value xsd.DateTime)
	GetTo() xsd.DateTime
	SetTo(value xsd.DateTime)
}

func (cv *ValidityPeriodTypeStruct) GetFrom() xsd.DateTime {
	return cv.From
}

func (cv *ValidityPeriodTypeStruct) GetTo() xsd.DateTime {
	return cv.To
}

func (cv *ValidityPeriodTypeStruct) SetFrom(value xsd.DateTime) {
	cv.From = value
}

func (cv *ValidityPeriodTypeStruct) SetTo(value xsd.DateTime) {
	cv.To = value
}

func NewValidityPeriodType() ValidityPeriodType {
	res := ValidityPeriodTypeStruct{}
	return &res
}
