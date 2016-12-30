package functions

import (
	"crypto/rand"
	"fmt"
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

func CreateNewBlockSignEntryElements(sid *ServerIdentity) (string, []byte, error) {
	pub, priv, err := ed.GenerateKey(rand.Reader)
	if err != nil {
		return "", nil, err
	}
	blockSigningPubkey := pub[:32]
	blockSigningPrivatekey := priv[:32]

	elements := "addentry -x 00 -e \"New Block Signing Key\" -x "
	elements += sid.RootChainID
	elements += " -x "
	elements += fmt.Sprintf("%032x", blockSigningPubkey)
	elements += " -x $now"
	idKey := *(sid.IDSet.IdentityLevel[0].GetPublicKey())
	preImage := append([]byte{0x01}, idKey[:]...)
	elements += fmt.Sprintf(" -x %x", preImage)
	elements += " -x $sig"

	elements += " -c "
	elements += sid.SubChainID
	return elements, blockSigningPrivatekey, nil
}

func CreateNewBlockSignEntryUnsigned(sid *ServerIdentity, blockSigningPubkey []byte) (string, error) {
	elements := "004E657720426C6F636B205369676E696E67204B6579" //00 and ascii New Block Signing Key
	elements += sid.RootChainID
	elements += fmt.Sprintf("%032x", blockSigningPubkey)
	return elements, nil
}
