package functions

/*
 * Creates a subchain for messages related to being a Federated/Audit/Candidate server.
 */

import (
	"encoding/hex"
	"github.com/FactomProject/serveridentity/identity"
	"reflect"
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
		return "error", err
	}

	elements := "addchain "

	structValue := reflect.ValueOf(sub)
	structElem := structValue.Elem()

	for i := 0; i < structElem.NumField(); i++ {
		el := structElem.Field(i)
		if i == 1 || i == 3 {
			elements += "-n \""
			elements += string(el.Interface().([]byte))
			elements += "\" "

		} else if i < 4 {
			elements += "-h "
			elements += hex.EncodeToString(el.Interface().([]byte))
			elements += " "
		}
	}

	sid.SubChainID = sub.ChainID

	return elements, nil
}
