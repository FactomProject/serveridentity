package identity

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/interfaces"
)

type MHash struct {
	version     []byte
	message     []byte
	rootChainID []byte
	newMHash    []byte
	timestamp   []byte
	preimage    []byte
	signiture   []byte

	subchain string
}

func MakeMHash(rootChainID string, subChainID string, seed string, privateKey *[64]byte) (*MHash, error) {
	m := new(MHash)
	m.version = []byte{0x00}
	m.message = []byte("New Matryoshka Hash")
	if root, err := hex.DecodeString(rootChainID); err != nil {
		return nil, err
	} else {
		m.rootChainID = root
	}

	m.subchain = subChainID

	if hash, err := generateMHash(seed); err != nil {
		return nil, err
	} else {
		m.newMHash = hash
	}

	t := interfaces.NewTimeStampNow()
	if timestamp, err := t.MarshalBinary(); err != nil {
		return nil, err
	} else {
		m.timestamp = timestamp
	}

	preI := make([]byte, 0)
	preI = append(preI, []byte{0x01}...)
	preI = append(preI, privateKey[32:]...)
	m.preimage = preI

	sig := ed.Sign(privateKey, m.versionToTimestamp())
	m.signiture = sig[:]
	return m, nil
}

func generateMHash(seed string) ([]byte, error) {
	s, err := hex.DecodeString(seed)
	var holder [32]byte
	copy(holder[:], s)
	if err != nil {
		return nil, err
	}
	for i := 0; i < MHashAmount; i++ {
		holder = sha256.Sum256(holder[:])
	}
	return holder[:], nil
}

func (m *MHash) GetEntry() *factom.Entry {
	e := new(factom.Entry)
	e.ChainID = m.subchain
	e.Content = []byte{}
	e.ExtIDs = m.extIdList()

	return e
}

func (m *MHash) GetMHash() string {
	return hex.EncodeToString(m.newMHash)
}

func (m *MHash) versionToTimestamp() []byte {
	buf := new(bytes.Buffer)
	buf.Write(m.version)
	buf.Write(m.message)
	buf.Write(m.rootChainID)
	buf.Write(m.newMHash)
	buf.Write(m.timestamp)

	return buf.Bytes()
}

func (m *MHash) extIdList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, m.version)
	list = append(list, m.message)
	list = append(list, m.rootChainID)
	list = append(list, m.newMHash)
	list = append(list, m.timestamp)
	list = append(list, m.preimage)
	list = append(list, m.signiture)

	return list
}
