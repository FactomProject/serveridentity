package functions

import (
	"fmt"

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

// Up to timestamp
func CreateNewMHashElementsUnsigned(rootChainID string, subChainID string, privateKey []byte, seed string, ec *factom.ECAddress) (string, error) {
	var priv [64]byte
	copy(priv[:32], privateKey[:32])
	_ = ed.GetPublicKey(&priv)

	mh, err := identity.MakeMHash(rootChainID, subChainID, seed, &priv)
	if err != nil {
		return "", err
	}

	e := mh.GetEntry()
	var toTimestamp []byte

	for _, exID := range e.ExtIDs[:4] {
		toTimestamp = append(toTimestamp, exID[:]...)
	}

	elements := fmt.Sprintf("%x", toTimestamp)

	return elements, nil
}

func CreateNewMHashElements(rootChainID string, subChainID string, privateKey []byte, seed string, ec *factom.ECAddress) (string, error) { // priv key
	var priv [64]byte
	copy(priv[:32], privateKey[:32])
	_ = ed.GetPublicKey(&priv)

	mh, err := identity.MakeMHash(rootChainID, subChainID, seed, &priv)
	if err != nil {
		return "", err
	}

	e := mh.GetEntry()
	extIDs := e.ExtIDs
	elements := "addentry"
	elements += fmt.Sprintf(" -x %x", extIDs[0][:])             // Version
	elements += fmt.Sprintf(" -e \"%s\"", string(extIDs[1][:])) // "New Matryoshka Hash"
	elements += fmt.Sprintf(" -x %x", extIDs[2][:])             // Root chain
	elements += fmt.Sprintf(" -x %x", extIDs[3][:])             // MHash
	elements += " -x $now"                                      // Timestamp (4)
	elements += fmt.Sprintf(" -x %x", extIDs[5][:])             // Preimage
	elements += " -x $sigMHASH"                                 // Signature (6)

	elements += " -c "
	elements += subChainID

	return elements, nil
}
