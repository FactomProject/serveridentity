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
	Version          []byte
	Message          []byte
	EncodedChainID   []byte
	IdentityPreimage []byte
	Signiture        []byte
}

func MakeRegisterIdentity(idSet *IdentitySet, chainID string) (*RegisterIdentity, error) {
	r := new(RegisterIdentity)
	r.Version = []byte{0x00}
	r.Message = []byte("Register Factom Identity")
	hex, err := hex.DecodeString(chainID)
	if err != nil {
		return nil, err
	}
	r.EncodedChainID = hex
	r.IdentityPreimage = append([]byte{0x01}, idSet.IdentityLevel[RegisterIdentityLevel-1].GetPublicKey()[:]...)

	sigMsg := make([]byte, 0)
	sigMsg = append(sigMsg, r.Version[:]...)
	sigMsg = append(sigMsg, r.Message[:]...)
	sigMsg = append(sigMsg, r.EncodedChainID[:]...)
	//sigMsg = append(sigMsg, r.IdentityPreimage[:]...)

	priv := idSet.IdentityLevel[RegisterIdentityLevel-1].GetPrivateKey()
	r.Signiture = ed.Sign(priv, sigMsg)[:]

	return r, nil
}

func (r *RegisterIdentity) extIDList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, r.Version)
	list = append(list, r.Message)
	list = append(list, r.EncodedChainID)
	list = append(list, r.IdentityPreimage)
	list = append(list, r.Signiture)

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
	Version          []byte
	Message          []byte
	subChainID       []byte
	IdentityPreimage []byte
	Signiture        []byte
}

func MakeRegisterSubchain(idSet *IdentitySet, chainID string) (*RegisterSubchain, error) {
	r := new(RegisterSubchain)
	r.Version = []byte{0x00}
	r.Message = []byte("Register Server Management")
	hex, err := hex.DecodeString(chainID)
	if err != nil {
		return nil, err
	}
	r.subChainID = hex
	r.IdentityPreimage = append([]byte{0x01}, idSet.IdentityLevel[RegisterIdentityLevel-1].GetPublicKey()[:]...)

	sigMsg := make([]byte, 0)
	sigMsg = append(sigMsg, r.Version[:]...)
	sigMsg = append(sigMsg, r.Message[:]...)
	sigMsg = append(sigMsg, r.subChainID[:]...)
	//sigMsg = append(sigMsg, r.IdentityPreimage[:]...)

	priv := idSet.IdentityLevel[RegisterIdentityLevel-1].GetPrivateKey()
	r.Signiture = ed.Sign(priv, sigMsg)[:]

	return r, nil
}

func (r *RegisterSubchain) extIDList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, r.Version)
	list = append(list, r.Message)
	list = append(list, r.subChainID)
	list = append(list, r.IdentityPreimage)
	list = append(list, r.Signiture)

	return list
}

func (r *RegisterSubchain) GetEntry(rootChainID string) *factom.Entry {
	e := new(factom.Entry)
	e.ChainID = rootChainID
	e.ExtIDs = r.extIDList()
	e.Content = []byte{}

	return e
}
