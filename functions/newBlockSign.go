package functions

import (
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
)

func CreateNewBlockSignEntry(rootChainID string, subChainID string, levelAbovePrivate []byte, ec *factom.ECAddress) (string, string, []byte, error) {
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
	bs, newPriv, err := identity.MakeBlockSigningKeyFixed(rootChainID, subChainID, &priv, true)
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

func CreateNewBlockSignEntryElements(sid *ServerIdentity) (string, error) {
	elements := "addentry -x 00 -e \"New Block Signing Key\" -x "
	elements += sid.RootChainID
	elements += " -x "
	elements += " %%%identity key here%%% "
	elements += " -x $now -x $sig"
	
	elements += " -c "
	elements += sid.SubChainID
	return elements, nil
}

func CreateNewBlockSignEntryUnsigned(sid *ServerIdentity) (string, error) {
	elements := "004E657720426C6F636B205369676E696E67204B6579"
	elements += sid.RootChainID
	elements += " %%%identity key here%%% "
	return elements, nil
}

