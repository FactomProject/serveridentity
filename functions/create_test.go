package functions_test

import (
//"bytes"
//"encoding/hex"
//"fmt"
//ed "github.com/FactomProject/ed25519"
//"github.com/FactomProject/factom"
//. "github.com/FactomProject/serveridentity/functions"
//"testing"
)

/*
func TestBTCKey(t *testing.T) {
	priv, _ := hex.DecodeString("f84a80f204c8e5e4369a80336919f55885d0b093505d84b80d12f9c08b81cd5e")
	var p [64]byte
	copy(p[:32], priv[:32])
	_ = ed.GetPublicKey(&p)

	btcKey, _ := hex.DecodeString("c5b7fd920dce5f61934e792c7e6fcc829aff533d")
	time, _ := hex.DecodeString("00000000495EAA80")
	ec, err := factom.MakeECAddress(priv[:32])

	b, err := CreateNewBitcoinKey("888888d027c59579fc47a6fc6c4a5c0409c7c39bc38a86cb5fc0069978493762", "888888d027c59579fc47a6fc6c4a5c0409c7c39bc38a86cb5fc0069978493762", 0, 0, btcKey, priv, ec)

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
	version, _ = hex.DecodeString("00")
	a, _ = hex.DecodeString("c5b7fd920dce5f61934e792c7e6fcc829aff533d")
	buf.Write(a)
	a, _ = hex.DecodeString("00000000495EAA80")
	buf.Write(a)

	sig := ed.Sign(&p, buf.Bytes())

	fmt.Println("BTC: " + hex.EncodeToString(sig[:]))
}*/

/*
func TestSig(t *testing.T) {
	priv, _ := hex.DecodeString("f84a80f204c8e5e4369a80336919f55885d0b093505d84b80d12f9c08b81cd5e")
	var p [64]byte
	copy(p[:32], priv[:32])
	_ = ed.GetPublicKey(&p)
	fmt.Println(hex.EncodeToString(p[32:]))

	buf := new(bytes.Buffer)
	a, _ := hex.DecodeString("00")
	buf.Write(a)
	a, _ = hex.DecodeString("4E657720426C6F636B205369676E696E67204B6579")
	buf.Write(a)
	a, _ = hex.DecodeString("888888d027c59579fc47a6fc6c4a5c0409c7c39bc38a86cb5fc0069978493762")
	buf.Write(a)
	a, _ = hex.DecodeString("8473745873ec04073ecf005b0d2b6cfe2f05f88f025e0c0a83a40d1de696a9cb")
	buf.Write(a)
	a, _ = hex.DecodeString("00000000495EAA80")
	buf.Write(a)

	sig := ed.Sign(&p, buf.Bytes())
	fmt.Println(buf.Len())
	fmt.Println(hex.EncodeToString(sig[:]))
}*/
