package functions

import (
	"fmt"

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

// Up to timestamp
func CreateNewBitcoinKeyElementsUnsigned(rootChainID string, subChainID string, btcKeyLevel int, btcType int, btcKey []byte, privateKey []byte, ec *factom.ECAddress) (string, error) {
	// priv key
	var priv [64]byte
	copy(priv[:32], privateKey[:32])
	_ = ed.GetPublicKey(&priv)

	btcKeyStruct, err := identity.MakeBitcoinKey(rootChainID, subChainID, btcKeyLevel, btcType, btcKey, &priv)
	if err != nil {
		return "", err
	}

	e := btcKeyStruct.GetEntry()
	var toTimestamp []byte

	for _, exID := range e.ExtIDs[:6] {
		toTimestamp = append(toTimestamp, exID[:]...)
	}

	elements := fmt.Sprintf("%x", toTimestamp)

	return elements, nil
}

func CreateNewBitcoinKeyElements(rootChainID string, subChainID string, btcKeyLevel int, btcType int, btcKey []byte, privateKey []byte, ec *factom.ECAddress) (string, error) {
	// priv key
	var priv [64]byte
	copy(priv[:32], privateKey[:32])
	_ = ed.GetPublicKey(&priv)

	btcKeyStruct, err := identity.MakeBitcoinKey(rootChainID, subChainID, btcKeyLevel, btcType, btcKey, &priv)
	if err != nil {
		return "", err
	}

	e := btcKeyStruct.GetEntry()
	extIDs := e.ExtIDs
	elements := "addentry " + ELEMENTS_FLAG
	elements += fmt.Sprintf(" -x %x", extIDs[0][:])             // Version
	elements += fmt.Sprintf(" -e \"%s\"", string(extIDs[1][:])) // "New Bitcoin Key"
	elements += fmt.Sprintf(" -x %x", extIDs[2][:])             // Root chain
	elements += fmt.Sprintf(" -x %x", extIDs[3][:])             // Bitcoin Key Level
	elements += fmt.Sprintf(" -x %x", extIDs[4][:])             // Bitcoin Key Type
	elements += fmt.Sprintf(" -x %x", extIDs[5][:])             // Bitcoin Key
	elements += " -x $now"                                      // Timestamp (6)
	elements += fmt.Sprintf(" -x %x", extIDs[7][:])             // Preimage
	elements += " -x $sigBTC"                                   // Signature (7)

	elements += " -c "
	elements += subChainID

	return elements, nil
}
