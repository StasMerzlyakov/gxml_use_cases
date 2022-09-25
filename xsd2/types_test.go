package xsd2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDecimal(t *testing.T) {
	value := "1.23454"
	decimal, err := NewDecimal(value)
	assert.Nil(t, err)
	assert.Equal(t, 1.23454, decimal.GetValue())
}

func TestNewInteger(t *testing.T) {
	value := "123"
	intVar, err := NewInteger(value)
	assert.Nil(t, err)
	assert.Equal(t, 123, intVar.GetValue())
}

func TestNewGDay(t *testing.T) {
	value := "----10"
	gDay, err := NewGDay(value)
	assert.Nil(t, err)
	assert.Equal(t, 10, gDay.GetValue())
}

func TestNewGMonth(t *testing.T) {
	value := "--10"
	gMonth, err := NewGMonth(value)
	assert.Nil(t, err)
	assert.Equal(t, 10, gMonth.GetValue())
}

func TestNewGYear(t *testing.T) {
	value := "2010"
	gYear, err := NewGYear(value)
	assert.Nil(t, err)
	assert.Equal(t, 2010, gYear.GetValue())
}
