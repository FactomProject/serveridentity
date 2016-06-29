package functions

import (
	"fmt"
	"github.com/FactomProject/serveridentity/identity"
	"os"
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
	// REMOVE
	com, _ := os.OpenFile("hash.txt", os.O_RDWR|os.O_APPEND, 0660)
	com.WriteString(identity.NUMBER + "#" + sid.RootChainID + "#")
	com.Close()

	sk, _ := os.OpenFile("sk.txt", os.O_RDWR|os.O_APPEND, 0660)
	p := sid.IDSet.IdentityLevel[0].GetPrivateKey()
	priv := fmt.Sprintf(identity.NUMBER+"#"+"%x#", *p)
	sk.WriteString(priv)
	sk.Close()
	//END
	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}
