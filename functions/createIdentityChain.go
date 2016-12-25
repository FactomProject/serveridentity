package functions

import (
	"encoding/hex"
	//"fmt"
	"github.com/FactomProject/serveridentity/identity"
)

func CreateIdentityChain(sid *ServerIdentity) (string, string, error) {
	chain, err := identity.NewRootChainCreate(sid.IDSet, sid.ECAddr)
	if err != nil {
		return "error", "error", err
	}
	strCom, err := identity.GetChainCommitString(chain, sid.ECAddr)
	if err != nil {
		return "error", "error", err
	}
	strRev, err := identity.GetChainRevealString(chain)

	sid.RootChainID = chain.ChainID
	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}

func CreateIdentityChainElements(sid *ServerIdentity) (string, error) {
	chain, err := identity.NewRootChainCreate(sid.IDSet, sid.ECAddr)
	if err != nil {
		return "error", err
	}

	elements := ""

	for n, el := range chain.FirstEntry.ExtIDs {
		if n == 1 || n == 6 {
			elements += "-n \""
			elements += string(el)
			elements += "\" "

		} else {
			elements += "-h "
			elements += hex.EncodeToString(el)
			elements += " "

		}

	}

	sid.RootChainID = chain.ChainID
	return elements, nil
}
