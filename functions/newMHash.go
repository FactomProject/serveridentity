package functions

import (
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
)

// Gets curl commands to make a new Matryoshka Hash
// Params:
// 		rootChainID 	string
//		seed			string	= Must be hex
//		privateKey		[]byte	=	len:32
//		ec				*factom.ECAddress
func CreateNewMHash(rootChainID string, subChainID string, privateKey []byte, seed string, ec *factom.ECAddress) (string, string, string, error) {
	var priv [64]byte
	copy(priv[:32], privateKey[:32])
	_ = ed.GetPublicKey(&priv)

	mh, err := identity.MakeMHash(rootChainID, subChainID, seed, &priv)

	e := mh.GetEntry()
	strCom, err := identity.GetEntryCommitString(e, ec)
	if err != nil {
		return "error", "error", "error", err
	}

	strRev, err := identity.GetEntryRevealString(e)
	if err != nil {
		return "error", "error", "error", err
	}

	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), mh.GetMHash(), nil
}
