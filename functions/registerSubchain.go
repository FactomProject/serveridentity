package functions

import (
	"github.com/FactomProject/serveridentity/identity"
	"os"
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

	// TODO: REMOVE
	sub, _ := os.OpenFile("subhash.txt", os.O_RDWR|os.O_APPEND, 0660)
	sub.WriteString(identity.NUMBER + "#" + sid.SubChainID + "#")
	sub.Close()
	//END

	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}
