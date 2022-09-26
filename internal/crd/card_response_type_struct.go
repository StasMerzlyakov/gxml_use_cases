package crd

import (
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type CardResponseTypeStruct struct {
	CardData common.CardDataType
	CVC      xsd2.String
}

type CardResponseType interface {
	GetCardData() common.CardDataType
	SetCardData(value common.CardDataType)
	GetCVC() xsd2.String
	SetCVC(value xsd2.String)
}

func (sv *CardResponseTypeStruct) GetCardData() common.CardDataType {
	return sv.CardData
}
func (sv *CardResponseTypeStruct) SetCardData(value common.CardDataType) {
	sv.CardData = value
}

func (sv *CardResponseTypeStruct) GetCVC() xsd2.String {
	return sv.CVC
}

func (sv *CardResponseTypeStruct) SetCVC(value xsd2.String) {
	sv.CVC = value
}

func NewCardResponseType() CardResponseType {
	obj := CardResponseTypeStruct{}
	return &obj
}
