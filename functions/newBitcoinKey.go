package functions

import (
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
)

// Gets curl commands to make a new bitcoin key
// Params:
// 		rootChainID 	string
//		subChainID		string
//		btcKeyLevel		int		=	1-4
//		btcType			int		=	0-1
//		btcKey			[]byte	=	len:20
//		pubKey			[]byte	=	len:32
//		privateKey		[]byte	=	len:32
//		ec				*factom.ECAddress
func CreateNewBitcoinKey(rootChainID string, subChainID string, btcKeyLevel int, btcType int, btcKey []byte, privateKey []byte, ec *factom.ECAddress) (string, string, error) {
	// priv key
	var priv [64]byte
	copy(priv[:32], privateKey[:32])
	_ = ed.GetPublicKey(&priv)

	bk, err := identity.MakeBitcoinKey(rootChainID, subChainID, btcKeyLevel, btcType, btcKey, &priv)
	if err != nil {
		return "error", "error", err
	}

	e := bk.GetEntry()
	strCom, err := identity.GetEntryCommitString(e, ec)
	if err != nil {
		return "error", "error", err
	}

	strRev, err := identity.GetEntryRevealString(e)
	if err != nil {
		return "error", "error", err
	}

	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), nil
}
