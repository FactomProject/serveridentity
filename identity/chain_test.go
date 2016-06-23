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
	e.Content = []byte("Main Identity List")

	chain := factom.NewChain(e)
	ec, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	ecAddr, _ := factom.MakeECAddress(ec[:32])
	str, _ := GetChainCommitString(chain, ecAddr)
	//fmt.Println(f.CurlWrapPOST(str))

	str, _ = GetChainRevealString(chain)
	//fmt.Println(f.CurlWrapPOST(str))
	//fmt.Println(chain.ChainID)

	e.ChainID = "8888889b595fbc374d5e752d7fa3c0d601f10570fa8f3e0ea9ad352c8465d075"
	str, _ = GetEntryCommitString(e, ecAddr)
	fmt.Println(f.CurlWrapPOST(str))

	str, _ = GetEntryRevealString(e)
	fmt.Println(f.CurlWrapPOST(str))
	//fmt.Println(chain.ChainID)

	_ = str
}
