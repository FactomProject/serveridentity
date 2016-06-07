package identity

import (
	"bytes"
	"encoding/hex"
	"errors"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/interfaces"
)

type BitcoinKey struct {
	version          []byte
	message          []byte
	rootChainID      []byte
	btcKeyLevel      []byte // 1 - 4
	btcType          []byte // 0 - 1
	btcKey           []byte
	timestamp        []byte
	identityPreimage []byte
	signiture        []byte
	subchain         string // Not included in ExtIDs
}

// Creates a new BitcoinKey type. Used to add bitcoinkeys to subchain
// Params:
// 		rootChainID 	string
//		subChainID		string
//		btcKeyLevel		int		=	1-4
//		btcType			int		=	0-1
//		btcKey			[]byte	=	len:20
//		identiyPrimage	[]byte	= 	len:33
//		privateKey		*[64]byte
func MakeBitcoinKey(rootChainID string, subChainID string, btcKeyLevel int, btcType int, btcKey []byte, privateKey *[64]byte) (*BitcoinKey, error) {
	bk := new(BitcoinKey)
	bk.version = []byte{0x00}
	bk.message = []byte("New Bitcoin Key")
	root, err := hex.DecodeString(rootChainID)
	bk.rootChainID = root
	if err != nil {
		return nil, err
	}
	idChainCheck := ProofOfWorkChainID[:ProofOfWorkLength]
	if bytes.Compare(bk.rootChainID[:ProofOfWorkLength], idChainCheck) != 0 {
		return nil, errors.New("Error creating new BTC key: Invalid root chain id")
	}
	bk.btcKeyLevel, err = intToOneByte(btcKeyLevel)
	if err != nil {
		return nil, err
	}
	if btcType > 1 || btcType < 0 {
		return nil, errors.New("Error creating new BTC key: Bitcoin key type must be either 0 or 1")
	}
	bk.btcType, err = intToOneByte(btcType)
	if err != nil {
		return nil, err
	}
	if len(btcKey) != 20 {
		return nil, errors.New("Error creating new BTC key: Incorrect bitcoin key length")
	}
	bk.btcKey = btcKey
	t := interfaces.NewTimeStampNow()
	bk.timestamp, err = t.MarshalBinary()

	preI := make([]byte, 0)
	preI = append(preI, []byte{0x01}...)
	preI = append(preI, privateKey[32:]...)
	bk.identityPreimage = preI
	mes := ed.Sign(privateKey, bk.versionToTimestamp())
	bk.signiture = mes[:]

	return bk, nil
}

func (b *BitcoinKey) GetEntry() *factom.Entry {
	e := new(factom.Entry)
	e.ChainID = b.subchain
	e.Content = []byte{}
	e.ExtIDs = b.extIdList()

	return e
}

func intToOneByte(i int) ([]byte, error) {
	switch i {
	case 0:
		return []byte{0x00}, nil
	case 1:
		return []byte{0x01}, nil
	case 2:
		return []byte{0x02}, nil
	case 3:
		return []byte{0x03}, nil
	}
	return nil, errors.New("Error creating new BTC key: Bitcoin key level must be between 0 and 3")
}

func (b *BitcoinKey) versionToTimestamp() []byte {
	buf := new(bytes.Buffer)
	buf.Write(b.version)
	buf.Write(b.message)
	buf.Write(b.rootChainID)
	buf.Write(b.btcKeyLevel)
	buf.Write(b.btcType)
	buf.Write(b.btcKey)
	buf.Write(b.timestamp)

	return buf.Bytes()
}

func (b *BitcoinKey) extIdList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, b.version)
	list = append(list, b.message)
	list = append(list, b.rootChainID)
	list = append(list, b.btcKeyLevel)
	list = append(list, b.btcType)
	list = append(list, b.btcKey)
	list = append(list, b.timestamp)
	list = append(list, b.identityPreimage)
	list = append(list, b.signiture)

	return list
}
