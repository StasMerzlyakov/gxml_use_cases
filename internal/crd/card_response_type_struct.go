package crd

import (
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
)

type CardResponseTypeStruct struct {
	CardData common.CardDataType
	CVC      xsd.String
}

type CardResponseType interface {
	GetCardData() common.CardDataType
	SetCardData(value common.CardDataType)
	GetCVC() xsd.String
	SetCVC(value xsd.String)
}

func (sv *CardResponseTypeStruct) GetCardData() common.CardDataType {
	return sv.CardData
}
func (sv *CardResponseTypeStruct) SetCardData(value common.CardDataType) {
	sv.CardData = value
}

func (sv *CardResponseTypeStruct) GetCVC() xsd.String {
	return sv.CVC
}

func (sv *CardResponseTypeStruct) SetCVC(value xsd.String) {
	sv.CVC = value
}

func NewCardResponseType() CardResponseType {
	obj := CardResponseTypeStruct{}
	return &obj
}
