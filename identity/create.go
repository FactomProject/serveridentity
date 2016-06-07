package identity

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/factom"
	"strconv"
)

var (
	bruteForceDisplay int = 100000
)

/**************************
 *    Shared Functions    *
 **************************/

type generalChainCreate interface {
	upToNonce() []byte
	getNonce() []byte
}

func findValidNonce(i generalChainCreate) []byte {
	upToNonce := i.upToNonce()
	var count int
	count = 000
	exit := false
	for exit == false {
		count++
		exit = CheckNonce(upToNonce, count)
		if ShowBruteForce == true && count%bruteForceDisplay == 0 {
			fmt.Print(".") // So user knows computation is happening
		}
	}
	fmt.Println("")
	return []byte(strconv.Itoa(count))
}

func calcChainID(i generalChainCreate) []byte {
	buf := new(bytes.Buffer)
	buf.Write(i.upToNonce()[:])

	result := sha256.Sum256(i.getNonce())
	buf.Write(result[:])

	result = sha256.Sum256(buf.Bytes())

	return result[:]
}

/**************************
 *       Main Chain       *
 **************************/

func NewRootChainCreate(idSet *IdentitySet, ecAddr *factom.ECAddress) (*factom.Chain, error) {
	root := MakeRootChainFromIdentitySet(idSet)

	// Entry
	entry := new(factom.Entry)
	entry.ExtIDs = root.extIdList()
	entry.Content = make([]byte, 0)
	entry.ChainID = hex.EncodeToString(calcChainID(root))

	// Chain
	chain := factom.NewChain(entry)
	return chain, nil
}

type RootChain struct {
	version    []byte
	message    []byte
	levelOne   []byte
	levelTwo   []byte
	levelThree []byte
	levelFour  []byte
	nonce      []byte
	IdSet      *IdentitySet
}

/*
 * Generates all the keys and nonce
 * Preffered to use MakeRootChainFromIdentitySet
 */
func MakeRootChain() *RootChain {
	i := new(RootChain)
	i.version = []byte{0x00}
	i.message = []byte("Identity Chain")

	idSet := NewIdentitySet()
	idSet.GenerateIdentitySet()

	i.levelOne = idSet.IdentityLevel[0].GetIdentityKey()[:]
	i.levelTwo = idSet.IdentityLevel[1].GetIdentityKey()[:]
	i.levelThree = idSet.IdentityLevel[2].GetIdentityKey()[:]
	i.levelFour = idSet.IdentityLevel[3].GetIdentityKey()[:]
	i.nonce = []byte{0x00}
	i.IdSet = idSet
	i.nonce = findValidNonce(i)

	return i
}

/*
 * If you have a set of 4 private keys, the IdentityChainHead can be
 * constructed from them
 *
 * param:
 * 		IdentitySet id = A type IdentitySet that constructs IdentityKeys
 *						 from a set of private keys
 */
func MakeRootChainFromIdentitySet(idSet *IdentitySet) *RootChain {
	i := new(RootChain)
	i.version = []byte{0x00}
	i.message = []byte("Identity Chain")

	i.levelOne = idSet.IdentityLevel[0].GetIdentityKey()[:]
	i.levelTwo = idSet.IdentityLevel[1].GetIdentityKey()[:]
	i.levelThree = idSet.IdentityLevel[2].GetIdentityKey()[:]
	i.levelFour = idSet.IdentityLevel[3].GetIdentityKey()[:]

	i.nonce = findValidNonce(i)

	return i
}

/*
 * 7 Elements
 * 1: version (0)
 * 2: ASCII bytes "Identity Chain"
 * 3: Level 1 Identity Key
 * 4-6: levels 2-4
 * 7: Nonce
 */
func (i *RootChain) upToNonce() []byte {
	buf := new(bytes.Buffer)

	result := sha256.Sum256(i.version)
	buf.Write(result[:])

	result = sha256.Sum256(i.message)
	buf.Write(result[:])

	result = sha256.Sum256(i.levelOne)
	buf.Write(result[:])

	result = sha256.Sum256(i.levelTwo)
	buf.Write(result[:])

	result = sha256.Sum256(i.levelThree)
	buf.Write(result[:])

	result = sha256.Sum256(i.levelFour)
	buf.Write(result[:])

	return buf.Bytes()
}

func (i *RootChain) getNonce() []byte {
	return i.nonce
}

func CheckNonce(upToNonce []byte, nonceInt int) bool {
	buf := new(bytes.Buffer)
	buf.Write(upToNonce)

	nonce := []byte(strconv.Itoa(nonceInt))
	result := sha256.Sum256(nonce)
	buf.Write(result[:])

	result = sha256.Sum256(buf.Bytes())

	chainFront := result[:ProofOfWorkLength]

	if bytes.Compare(chainFront[:ProofOfWorkLength], ProofOfWorkChainID[:ProofOfWorkLength]) == 0 {
		return true
	}
	return false
}

/*
 * Returns external ids in [][]byte to pass to other functions
 */
func (i *RootChain) extIdList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, i.version)
	list = append(list, i.message)
	list = append(list, i.levelOne)
	list = append(list, i.levelTwo)
	list = append(list, i.levelThree)
	list = append(list, i.levelFour)
	list = append(list, i.nonce)

	return list
}

/**************************
 *       Sub Chain        *
 **************************/

type Subchain struct {
	version     []byte
	message     []byte
	rootChainID []byte
	nonce       []byte
	ChainID     string
}

// Chain ID of root identity chain
func MakeSubChain(chainID string) (*Subchain, error) {
	sub := new(Subchain)
	sub.version = []byte{0x00}
	sub.message = []byte("Server Management")
	// TODO: Why can't sub.rootChainID, err := hex.DecodeString(chainID)
	root, err := hex.DecodeString(chainID)
	sub.rootChainID = root
	if err != nil {
		return nil, err
	}
	sub.nonce = findValidNonce(sub)
	sub.ChainID = hex.EncodeToString(calcChainID(sub))

	return sub, nil
}

func (sub *Subchain) GetFactomChain() *factom.Chain {
	e := new(factom.Entry)
	e.ExtIDs = sub.extIdList()
	e.Content = []byte{}
	e.ChainID = sub.ChainID

	chain := factom.NewChain(e)
	return chain
}

func (sub *Subchain) extIdList() [][]byte {
	list := make([][]byte, 0)
	list = append(list, sub.version)
	list = append(list, sub.message)
	list = append(list, sub.rootChainID)
	list = append(list, sub.nonce)

	return list
}

func (sub *Subchain) upToNonce() []byte {
	buf := new(bytes.Buffer)

	result := sha256.Sum256(sub.version)
	buf.Write(result[:])

	result = sha256.Sum256(sub.message)
	buf.Write(result[:])

	result = sha256.Sum256(sub.rootChainID)
	buf.Write(result[:])

	return buf.Bytes()
}

func (sub *Subchain) getNonce() []byte {
	return sub.nonce
}
