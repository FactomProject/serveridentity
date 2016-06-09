package functions

import (
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
)

func CreateNewBlockSignEntry(rootChainID string, levelAbovePrivate []byte, ec *factom.ECAddress) (string, string, []byte, error) {
	/* if level == 1 {
		return "error", "error", nil, errors.New("Error creating new block signing key: Cannot replace level 1 key")
	} */

	// Pub key to be replaced
	//idPreimage := make([]byte, 0)
	//idPreimage = append(idPreimage, []byte{0x01}...)
	//idPreimage = append(idPreimage, keyReplace[:32]...)

	// priv key level above to approve change
	var priv [64]byte
	copy(priv[:32], levelAbovePrivate[:32])
	_ = ed.GetPublicKey(&priv)
	bs, newPriv, err := identity.MakeBlockSigningKey(rootChainID, &priv)
	if err != nil {
		return "error", "error", nil, err
	}

	e := bs.GetEntry()
	strCom, err := identity.GetEntryCommitString(e, ec)
	if err != nil {
		return "error", "error", nil, err
	}

	strRev, err := identity.GetEntryRevealString(e)
	if err != nil {
		return "error", "error", nil, err
	}

	return CurlWrapPOST(strCom), CurlWrapPOST(strRev), newPriv, nil
}