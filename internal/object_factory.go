package internal

import (
	"github.com/StasMerzlyakov/gxml/xsd"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/common"
	"github.com/StasMerzlyakov/gxml_use_cases/internal/crd"
	"github.com/StasMerzlyakov/gxml_use_cases/xsd2"
)

var namespacesToPackageMap = map[string]string{
	"https://github.com/StasMerzlyakov/gxml/card-service": "crd",
	"https://github.com/StasMerzlyakov/gxml/common-data":  "common",
}

type Resolver struct {
}

func (Resolver) ResolveValidator(name xsd.NameAndNamespace) (any, xsd2.IElementValidator) {
	if packageName, ok := namespacesToPackageMap[name.Namespace]; !ok {
		return nil
	} else {
		switch packageName {
		case "crd":
			return crd.ResolveValidator(name)
		case "common":
			return common.ResolveValidator(name)
		default:
			return nil, nil
		}
	}
}
