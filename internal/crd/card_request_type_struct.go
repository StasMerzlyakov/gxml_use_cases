package crd

import (
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

type CardRequestTypeStruct struct {
	FirstName  xsd2.String
	LastName   xsd2.String
	Patronymic xsd2.String

	BirthDate common.ComplexDateType
	CardType  xsd2.String
}

type CardRequestType interface {
	GetFirstName() xsd2.String
	SetFirstName(value xsd2.String)
	GetLastName() xsd2.String
	SetLastName(value xsd2.String)
	GetPatronymic() xsd2.String
	SetPatronymic(value xsd2.String)
	GetBirthDate() common.ComplexDateTypeStruct
}
