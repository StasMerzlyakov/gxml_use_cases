package xsd2

import (
	"fmt"
	"regexp"
	"strconv"
)

type Integer interface {
	GetValue() int
}

type integerStruct struct {
	value int
}

func (str integerStruct) GetValue() int {
	return str.value
}

func NewInteger(value string) (Integer, error) {
	intVar, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return integerStruct{value: intVar}, nil
}

type Decimal interface {
	GetValue() float64
}

type decimalStruct struct {
	value float64
}

func (str decimalStruct) GetValue() float64 {
	return str.value
}

func NewDecimal(value string) (Decimal, error) {
	floatVar, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	return &decimalStruct{value: floatVar}, nil
}

type String interface {
	GetValue() string
}

type stringStruct struct {
	value string
}

func (str stringStruct) GetValue() string {
	return str.value
}

func NewString(value string) (String, error) {
	return stringStruct{value: value}, nil
}

type GYear interface {
	GetValue() int
}

type gYearStruct struct {
	value int
}

func (str gYearStruct) GetValue() int {
	return str.value
}

func NewGYear(value string) (GYear, error) {
	matched, err := regexp.MatchString("\\d{4}", value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, fmt.Errorf("value %s not match gYear pattern", value)
	}
	intVar, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return gYearStruct{value: intVar}, nil
}

type GMonth interface {
	GetValue() int
}

type gMonthStruct struct {
	value int
}

func (str gMonthStruct) GetValue() int {
	return str.value
}

func NewGMonth(value string) (GMonth, error) {

	matched, err := regexp.MatchString("--\\d{2}", value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, fmt.Errorf("value %s not match gMonth pattern", value)
	}

	value = value[2:4]

	intVar, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return gMonthStruct{value: intVar}, nil
}

type GDay interface {
	GetValue() int
}

type gDayStruct struct {
	value int
}

func (str gDayStruct) GetValue() int {
	return str.value
}

func NewGDay(value string) (GDay, error) {

	matched, err := regexp.MatchString("---\\d{2}", value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, fmt.Errorf("value %s not match gDay pattern", value)
	}

	value = value[3:5]

	intVar, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return gDayStruct{value: intVar}, nil
}

type DateTime interface {
}

type dateTimeStruct struct {
}

func NewDateTime(value string) (DateTime, error) {
	return dateTimeStruct{}, nil
}
