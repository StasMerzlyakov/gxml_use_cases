package api

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValidator1(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_type_1.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	xsdValidator, err := NewXsdValidator(r)
	if err != nil {
		t.Fatal(err)
	}
	_, err = xsdValidator.Validate()
	assert.Nil(t, err)
}

func TestValidator2(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_type_2.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	xsdValidator, err := NewXsdValidator(r)
	if err != nil {
		t.Fatal(err)
	}
	_, err = xsdValidator.Validate()
	assert.NotNil(t, err)
}
