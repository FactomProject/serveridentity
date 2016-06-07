package functions

import (
	"crypto/rand"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
)

// Individual component in an IdentitySet
type ServerIdentity struct {
	ECAddr      *factom.ECAddress
	IDSet       *identity.IdentitySet
	RootChainID string
	SubChainID  string
}

func newServerIdentity() *ServerIdentity {
	return new(ServerIdentity)
}

func MakeServerIdentity() (*ServerIdentity, error) {
	sid := newServerIdentity()
	_, sec, err := ed.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	sid.ECAddr, err = factom.MakeECAddress(sec[:32])
	if err != nil {
		return nil, err
	}
	sid.IDSet = identity.NewIdentitySet()
	sid.IDSet.GenerateIdentitySet()
	return sid, nil
}

func MakeServerIdentityFromEC(sec []byte) (*ServerIdentity, error) {
	sid := newServerIdentity()
	ecAddr, err := factom.MakeECAddress(sec[:32])
	if err != nil {
		return nil, err
	}

	sid.ECAddr = ecAddr
	sid.IDSet = identity.NewIdentitySet()
	sid.IDSet.GenerateIdentitySet()
	return sid, nil
}

// TODO: Make Server Identity from existing keys
func MakeServerIdentityFromKeys() (*ServerIdentity, error) {
	return nil, nil
}
