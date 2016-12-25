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

	elements := ""

	listingValue := reflect.ValueOf(reg)
	listingElem := listingValue.Elem()

	for i := 0; i < listingElem.NumField(); i++ {
		el := listingElem.Field(i)
		if i == 1 {
			elements += "-n \""
			elements += string(el.Interface().([]byte))
			elements += "\" "

		} else {
			elements += "-h "
			elements += hex.EncodeToString(el.Interface().([]byte))
			elements += " "
		}
	}

	return elements, nil
}
