package identity

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/interfaces"
)

type BlockSigningKey struct {
	version          []byte
	message          []byte
	rootChainID      []byte
	newPubKey        []byte
	timestamp        []byte
	identityPreimage []byte
	signiture        []byte
}

// Creates a new BlockSigningKey type. Used to change keys in identity chain
func MakeBlockSigningKey(rootChainIDStr string, levelAbovePrivateKey *[64]byte) (*BlockSigningKey, []byte, error) {
	rootChainID, err := hex.DecodeString(rootChainIDStr)
	if err != nil {
		return nil, nil, err
	}
	if bytes.Compare(rootChainID[:ProofOfWorkLength], ProofOfWorkChainID[:ProofOfWorkLength]) != 0 {
		return nil, nil, errors.New("Error making a new block signing key: Root ChainID invalid")
	}
	b := new(BlockSigningKey)
	b.version = []byte{0x00}
	b.message = []byte("New Block Signing Key")
	b.rootChainID = rootChainID
	pub, priv, err := ed.GenerateKey(rand.Reader)
	b.newPubKey = pub[:32]
	if err != nil {
		return nil, nil, err
	}
	t := interfaces.NewTimeStampNow()
	b.timestamp, err = t.MarshalBinary()
	if err != nil {
		return nil, nil, err
	}
	preI := make([]byte, 0)
	preI = append(preI, []byte{0x01}...)
	preI = append(preI, levelAbovePrivateKey[32:]...)
	b.identityPreimage = preI
	sig := ed.Sign(levelAbovePrivateKey, b.versionToTimestamp())
	b.signiture = sig[:]
	return b, priv[:], nil
}

func (b *BlockSigningKey) GetEntry() *factom.Entry {
	e := new(factom.Entry)
	e.ChainID = hex.EncodeToString(b.rootChainID)
	e.Content = []byte{}
	e.ExtIDs = b.extIdList()

	return e
}

func (b *BlockSigningKey) versionToTimestamp() []byte {
	buf := new(bytes.Buffer)
	buf.Write(b.version)
	buf.Write(b.message)
	buf.Write(b.rootChainID)
	buf.Write(b.newPubKey)
	buf.Write(b.timestamp)

	return buf.Bytes()
}

func (b *BlockSigningKey) extIdList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, b.version)
	list = append(list, b.message)
	list = append(list, b.rootChainID)
	list = append(list, b.newPubKey)
	list = append(list, b.timestamp)
	list = append(list, b.identityPreimage)
	list = append(list, b.signiture)

	return list
}
