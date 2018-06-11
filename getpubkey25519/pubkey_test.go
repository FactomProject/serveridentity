package main_test

import (
	"encoding/hex"
	"testing"
	"bytes"

	. "github.com/FactomProject/serveridentity/getpubkey25519"
)

//https://tools.ietf.org/html/draft-josefsson-eddsa-ed25519-02 testvector
func TestPubkey(t *testing.T) {
	args := []string{"9d61b19deffd5a60ba844af492ec2cc44449c5697b326919703bac031cae7f60"}

	pubkey := PubkeyOfPrivate(args)

	pubkyeB, err := hex.DecodeString(pubkey)
	if err != nil {
		t.Error(err)
	}
	testpub, err := hex.DecodeString("d75a980182b10ab7d54bfed3c964073a0ee172f3daa62325af021a68f707511a")
	if err != nil {
		t.Error(err)
	}


	b := bytes.Compare(pubkyeB, testpub)
	if 0 != b {
		t.Error("Not the right pubkey")
	}
}

