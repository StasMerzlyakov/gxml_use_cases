package api

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/crd"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValidator1(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_response_1.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	xsdValidator, err := NewXsdValidator(r)
	if err != nil {
		t.Fatal(err)
	}
	obj, err := xsdValidator.Validate()
	crt, ok := obj.(crd.CardResponseType)
	assert.True(t, ok)
	assert.NotNil(t, crt)
	assert.Nil(t, err)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	nameAndNamespace := xsd.NameAndNamespace{
		Namespace: "https://github.com/StasMerzlyakov/gxml/card-service",
		Name:      "CardResponse",
	}
	xsdValidator.Write(writer, crt, nameAndNamespace)
	str := b.String()
	fmt.Printf("%s\n", str)

	//reflect.TypeOf(crd.CardRequestType)

}

func TestValidator2(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_response_2.xml")
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

func TestValidator3(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_response_3.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	xsdValidator, err := NewXsdValidator(r)
	if err != nil {
		t.Fatal(err)
	}
	obj, err := xsdValidator.Validate()
	crt := obj.(crd.CardResponseType)
	assert.NotNil(t, crt)
	assert.Nil(t, err)
}

func TestValidator4(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_request_1.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	xsdValidator, err := NewXsdValidator(r)
	if err != nil {
		t.Fatal(err)
	}
	obj, err := xsdValidator.Validate()
	crt, ok := obj.(crd.CardRequestType)
	assert.True(t, ok)
	assert.NotNil(t, crt)
	assert.Nil(t, err)
}

func TestValidator5(t *testing.T) {
	file, err := os.Open("../test_files/xml/card_request_2.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	xsdValidator, err := NewXsdValidator(r)
	if err != nil {
		t.Fatal(err)
	}
	obj, err := xsdValidator.Validate()
	crt, ok := obj.(crd.CardRequestType)
	assert.True(t, ok)
	assert.NotNil(t, crt)
	assert.Nil(t, err)
}
