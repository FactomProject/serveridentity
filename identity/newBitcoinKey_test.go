package identity_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	ed "github.com/FactomProject/ed25519"
	//"github.com/FactomProject/factom"
	"encoding/binary"
	"github.com/FactomProject/factomd/common/interfaces"
	. "github.com/FactomProject/serveridentity/identity"
	"testing"
)

func TestBTCKey(t *testing.T) {
	for trial := 0; trial < 4; trial++ {
		priv, _ := hex.DecodeString("f84a80f204c8e5e4369a80336919f55885d0b093505d84b80d12f9c08b81cd5e")
		var p [64]byte
		copy(p[:32], priv[:32])
		_ = ed.GetPublicKey(&p)
		btcKey, _ := hex.DecodeString("c5b7fd920dce5f61934e792c7e6fcc829aff533d")
		//ec, err := factom.MakeECAddress(priv[:32])
		chain := "888888d027c59579fc47a6fc6c4a5c0409c7c39bc38a86cb5fc0069978493762"

		b, err := MakeBitcoinKey(chain, chain, 0, 0, btcKey, &p)
		if err != nil {
			fmt.Println(err.Error())
		}

		ti := interfaces.GetTime()
		by := make([]byte, 8)
		binary.BigEndian.PutUint64(by, ti)
		timestamp := by
		fmt.Printf("DEBUG: %x\n", timestamp)

		buf := new(bytes.Buffer)
		a, _ := hex.DecodeString("00")
		buf.Write(a)
		a, _ = hex.DecodeString("4e657720426974636f696e204b6579")
		buf.Write(a)
		a, _ = hex.DecodeString("888888d027c59579fc47a6fc6c4a5c0409c7c39bc38a86cb5fc0069978493762")
		buf.Write(a)
		a, _ = hex.DecodeString("00")
		buf.Write(a)
		a, _ = hex.DecodeString("00")
		buf.Write(a)
		a, _ = hex.DecodeString("c5b7fd920dce5f61934e792c7e6fcc829aff533d")
		buf.Write(a)
		//a, _ = hex.DecodeString("00000000495EAA80")
		buf.Write(timestamp)

		sig := ed.Sign(&p, buf.Bytes())
		//fmt.Println(len(b.GetEntry().ExtIDs[8]))
		//fmt.Println(len(sig[:]))

		//fmt.Println("BTC: " + hex.EncodeToString(sig[:]))
		//fmt.Printf("DEBUG: %x\n", b.GetEntry().ExtIDs[8])
		fmt.Println("Test Failed: Trying test again. Could be timestamp mismatch")
		if bytes.Compare(b.GetEntry().ExtIDs[8][:], sig[:]) != 0 {
			if trial == 3 {
				t.Error("BTCKey make fail. May be due to timestamp not matching")
			}
		} else {
			trial = 10
		}
	}
}
