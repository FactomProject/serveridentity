package identity_test

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/factom"
	f "github.com/FactomProject/serveridentity/functions"
	. "github.com/FactomProject/serveridentity/identity"
	"testing"
)

/*
func TestNewCommitChain(t *testing.T) {
	idChain := MakeIdentityChainHead()

	// Entry Credit Private key
	ecAddr := factom.NewECAddress()

	chain, _ := NewCommitChainCreate(idChain.IdSet, ecAddr)

	str, _ := GetCommitString(chain, ecAddr)
	fmt.Println(str)
}*/

func TestBlankChain(t *testing.T) {
	e := new(factom.Entry)
	e.Content = []byte{}

	chain := factom.NewChain(e)
	ec, _ := hex.DecodeString("9FAA5D459E16C50F192630487B52D78EAB2442B29E23BAD433C83986DBC5DA29")
	ecAddr, _ := factom.MakeECAddress(ec[:32])
	str, _ := GetChainCommitString(chain, ecAddr)
	fmt.Println(f.CurlWrapPOST(str))

	str, _ = GetChainRevealString(chain)
	fmt.Println(f.CurlWrapPOST(str))

	fmt.Println(chain.ChainID)
}
