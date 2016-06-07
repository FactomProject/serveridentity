package functions

import (
	"github.com/FactomProject/serveridentity/identity"
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
