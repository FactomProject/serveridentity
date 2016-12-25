package functions

/*
 * Creates a subchain for messages related to being a Federated/Audit/Candidate server.
 */

import (
	"github.com/FactomProject/serveridentity/identity"
)

func CreateSubChain(sid *ServerIdentity) (string, string, error) {
	sub, err := identity.MakeSubChain(sid.RootChainID)
	if err != nil {
		return "error", "error", err
	}

	chain := sub.GetFactomChain()
	strCom, err := identity.GetChainCommitString(chain, sid.ECAddr)
	if err != nil {
		return "error", "error", err
	}

	strRev, err := identity.GetChainRevealString(chain)
	if err != nil {
		return "error", "error", err
	}

	sid.SubChainID = chain.ChainID
	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}

func CreateSubChainElements(sid *ServerIdentity) (string, error) {
	sub, err := identity.MakeSubChain(sid.RootChainID)
	if err != nil {
		return "error", "error", err
	}

	chain := sub.GetFactomChain()
	strCom, err := identity.GetChainCommitString(chain, sid.ECAddr)
	if err != nil {
		return "error", "error", err
	}

	strRev, err := identity.GetChainRevealString(chain)
	if err != nil {
		return "error", "error", err
	}

	sid.SubChainID = chain.ChainID
	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}
