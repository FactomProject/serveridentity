package functions

import (
	"encoding/hex"
	"github.com/FactomProject/serveridentity/identity"
)

func RegisterSubchain(sid *ServerIdentity) (string, string, error) {
	reg, err := identity.MakeRegisterSubchain(sid.IDSet, sid.SubChainID)
	if err != nil {
		return "error", "error", err
	}

	e := reg.GetEntry(sid.RootChainID)
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

func RegisterSubChainElements(sid *ServerIdentity) (string, error) {
	reg, err := identity.MakeRegisterSubchain(sid.IDSet, sid.SubChainID)
	if err != nil {
		return "error", err
	}

	elements := "addentry "

	for i, el := range reg.ExtIDList() {
		if i == 1 {
			elements += "-e \""
			elements += string(el)
			elements += "\" "

		} else {
			elements += "-x "
			elements += hex.EncodeToString(el)
			elements += " "
		}
	}

	elements += "-c "
	elements += sid.RootChainID
	elements += " "

	return elements, nil
}
