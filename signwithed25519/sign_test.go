package main_test

import (
	"encoding/hex"
	"testing"

	ed "github.com/FactomProject/ed25519"
	. "github.com/FactomProject/serveridentity/signwithed25519"
)

func TestSignature(t *testing.T) {
	args := []string{"0000000000000000000000000000000000000000000000000000000000000000",
		"4c38c72fc5cdad68f13b74674d3ffb1f3d63a112710868c9b08946553448d26d"}

	signature := "f35d9b47b49c52873e64d198c44e0f7ac70df243b065c818ff7e2dc4ae21991073a1a9c7bc454091a59b68206e4a4506708ba1cc9a5b278894ee47bf5c7a620a"
	sig := SignatureOfMessage(args)

	if sig != signature {
		t.Errorf("Incorrect signature")
	}

	sigB, err := hex.DecodeString(sig)
	if err != nil {
		t.Error(err)
	}

	mes, err := hex.DecodeString(args[0])
	if err != nil {
		t.Error(err)
	}

	pub, err := hex.DecodeString("cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a")
	if err != nil {
		t.Error(err)
	}

	pubE := new([ed.PublicKeySize]byte)
	copy(pubE[:ed.PublicKeySize], pub[:ed.PublicKeySize])

	sigE := new([ed.SignatureSize]byte)
	copy(sigE[:ed.SignatureSize], sigB[:ed.SignatureSize])

	b := ed.VerifyCanonical(pubE, mes, sigE)
	if !b {
		t.Error("Not a valid sig")
	}
}
