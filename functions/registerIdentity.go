package functions

import (
	"encoding/hex"
	"github.com/FactomProject/serveridentity/identity"
	"reflect"
)

func RegisterServerIdentity(sid *ServerIdentity) (string, string, error) {
	reg, err := identity.MakeRegisterIdentity(sid.IDSet, sid.RootChainID)
	if err != nil {
		return "error", "error", err
	}

	e := reg.GetEntry()
	strCom, err := identity.GetEntryCommitString(e, sid.ECAddr)
	if err != nil {
		return "error", "error", err
	}

	strRev, err := identity.GetEntryRevealString(e)
	if err != nil {
		return "error", "error", err
	}

	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}

func RegisterServerIdentityElements(sid *ServerIdentity) (string, error) {
	reg, err := identity.MakeRegisterIdentity(sid.IDSet, sid.RootChainID)
	if err != nil {
		return "error", err
	}

	elements := "addentry "

	structValue := reflect.ValueOf(reg)
	structElem := structValue.Elem()

	for i := 0; i < structElem.NumField(); i++ {
		el := structElem.Field(i)
		if i == 1 {
			elements += "-e \""
			elements += string(el.Interface().([]byte))
			elements += "\" "

		} else {
			elements += "-x "
			elements += hex.EncodeToString(el.Interface().([]byte))
			elements += " "
		}
	}

	elements += "-c "
	elements += identity.RootRegisterChain
	elements += " "

	return elements, nil
}
