package identity

import (
	"encoding/hex"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
)

/**************************
 *       Main Chain       *
 **************************/

type RegisterIdentity struct {
	version          []byte
	message          []byte
	encodedChainID   []byte
	identityPreimage []byte
	signiture        []byte
}

func MakeRegisterIdentity(idSet *IdentitySet, chainID string) (*RegisterIdentity, error) {
	r := new(RegisterIdentity)
	r.version = []byte{0x00}
	r.message = []byte("Register Factom Identity")
	hex, err := hex.DecodeString(chainID)
	if err != nil {
		return nil, err
	}
	r.encodedChainID = hex
	r.identityPreimage = append([]byte{0x01}, idSet.IdentityLevel[RegisterIdentityLevel-1].GetPublicKey()[:]...)

	sigMsg := make([]byte, 0)
	sigMsg = append(sigMsg, r.version[:]...)
	sigMsg = append(sigMsg, r.message[:]...)
	sigMsg = append(sigMsg, r.encodedChainID[:]...)
	//sigMsg = append(sigMsg, r.identityPreimage[:]...)

	priv := idSet.IdentityLevel[RegisterIdentityLevel-1].GetPrivateKey()
	r.signiture = ed.Sign(priv, sigMsg)[:]

	return r, nil
}

func (r *RegisterIdentity) extIDList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, r.version)
	list = append(list, r.message)
	list = append(list, r.encodedChainID)
	list = append(list, r.identityPreimage)
	list = append(list, r.signiture)

	return list
}

func (r *RegisterIdentity) GetEntry() *factom.Entry {
	e := new(factom.Entry)
	e.ChainID = RootRegisterChain
	e.ExtIDs = r.extIDList()
	e.Content = []byte{}

	return e
}

/**************************
 *        Sub Chain       *
 **************************/

type RegisterSubchain struct {
	version          []byte
	message          []byte
	subChainID       []byte
	identityPreimage []byte
	signiture        []byte
}

func MakeRegisterSubchain(idSet *IdentitySet, chainID string) (*RegisterSubchain, error) {
	r := new(RegisterSubchain)
	r.version = []byte{0x00}
	r.message = []byte("Register Server Management")
	hex, err := hex.DecodeString(chainID)
	if err != nil {
		return nil, err
	}
	r.subChainID = hex
	r.identityPreimage = append([]byte{0x01}, idSet.IdentityLevel[RegisterIdentityLevel-1].GetPublicKey()[:]...)

	sigMsg := make([]byte, 0)
	sigMsg = append(sigMsg, r.version[:]...)
	sigMsg = append(sigMsg, r.message[:]...)
	sigMsg = append(sigMsg, r.subChainID[:]...)
	//sigMsg = append(sigMsg, r.identityPreimage[:]...)

	priv := idSet.IdentityLevel[RegisterIdentityLevel-1].GetPrivateKey()
	r.signiture = ed.Sign(priv, sigMsg)[:]

	return r, nil
}

func (r *RegisterSubchain) extIDList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, r.version)
	list = append(list, r.message)
	list = append(list, r.subChainID)
	list = append(list, r.identityPreimage)
	list = append(list, r.signiture)

	return list
}

func (r *RegisterSubchain) GetEntry(rootChainID string) *factom.Entry {
	e := new(factom.Entry)
	e.ChainID = rootChainID
	e.ExtIDs = r.extIDList()
	e.Content = []byte{}

	return e
}
