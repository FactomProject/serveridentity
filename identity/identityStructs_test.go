package identity_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	. "github.com/FactomProject/serveridentity/identity"
	"testing"
)

var debug bool = false

func TestGenerateIdentityFromPrivateKey(t *testing.T) {
	//fmt.Println("TESTING: Beginning GenerateIdentityFromPrivateKey Test")
	id := NewIdentity()
	seedKey, err := hex.DecodeString("f84a80f204c8e5e4369a80336919f55885d0b093505d84b80d12f9c08b81cd5e")
	if err != nil {
		t.Error(err)
	}
	var seedPriv [64]byte
	copy(seedPriv[:32], seedKey[:])
	//copy(seedPriv[32:], seedKey[:])

	err = id.GenerateIdentityFromPrivateKey(&seedPriv, 0)
	if err != nil {
		t.Error(err)
	}

	if debug == true {
		strPb := hex.EncodeToString(id.GetPublicKey()[:])
		strPv := hex.EncodeToString(id.GetPrivateKey()[:32])
		strId := hex.EncodeToString(id.GetIdentityKey()[:])
		//fmt.Println("Test Data in Hex:")
		//fmt.Print("Public Key:   ")
		//fmt.Println(strPb)
		//fmt.Print("Private Key:  ")
		//fmt.Println(strPv)
		//fmt.Print("Identity Key: ")
		//fmt.Println(strId)
	}

	a, _ := hex.DecodeString("3f2b77bca02392c95149dc769a78bc758b1037b6a546011b163af0d492b1bcc0")

	if bytes.Compare(id.GetIdentityKey()[:], a) != 0 {
		t.Error("Key generation not correct")
	}
}

func TestGenerateIdentitySet(t *testing.T) {
	//fmt.Println("TESTING: Beginning GenerateIdentitySet Test")
	idSet := NewIdentitySet()
	err := idSet.GenerateIdentitySet()
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateIdentitySetFromPrivateKeys(t *testing.T) {
	//fmt.Println("TESTING: Beginning GenerateIdentitySetFromPrivateKeys Test")
	idSet := NewIdentitySet()

	seedKey1, _ := hex.DecodeString("f84a80f204c8e5e4369a80336919f55885d0b093505d84b80d12f9c08b81cd5e")
	//seedKey1, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	seedKey2, _ := hex.DecodeString("2bb967a78b081fafef17818c2a4c2ba8dbefcd89664ff18f6ba926b55e00b601")
	seedKey3, _ := hex.DecodeString("09d51ae7cc0dbc597356ab1ada078457277875c81989c5db0ae6f4bf86ccea5f")
	seedKey4, _ := hex.DecodeString("72644033bdd70b8fec7aa1fea50b0c5f7dfadb1bce76aa15d9564bf71c62b160")
	var seedKeys [4]*[64]byte
	var seedKeyElement1 [64]byte
	var seedKeyElement2 [64]byte
	var seedKeyElement3 [64]byte
	var seedKeyElement4 [64]byte

	copy(seedKeyElement1[:32], seedKey1)
	seedKeys[0] = &seedKeyElement1
	copy(seedKeyElement2[:32], seedKey2)
	seedKeys[1] = &seedKeyElement2
	copy(seedKeyElement3[:32], seedKey3)
	seedKeys[2] = &seedKeyElement3
	copy(seedKeyElement4[:32], seedKey4)
	seedKeys[3] = &seedKeyElement4

	err := idSet.GenerateIdentitySetFromPrivateKeys(seedKeys)
	if err != nil {
		t.Error(err)
	}

	/*for i := 0; i < 4; i++ {
		str := hex.EncodeToString((idSet.IdentityLevel[i].GetPublicKey())[:32])
		fmt.Println("Private Key: " + str)

		str = hex.EncodeToString(idSet.IdentityLevel[i].PrivPrefix[:])
		str2 := hex.EncodeToString(idSet.IdentityLevel[i].PubPrefix[:])
		_ = str2
		//fmt.Println("Priv Prefix: " + str + "  |  Pub Prefix: " + str2)

		str = idSet.IdentityLevel[i].HumanReadablePublic()
		fmt.Println("Human Readable: " + str)
	}*/
}
