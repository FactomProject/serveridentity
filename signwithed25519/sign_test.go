package main_test

import (
	"testing"

	. "github.com/FactomProject/serveridentity/signwithed25519"
)

func TestSignature(t *testing.T) {
	args := []string{"0000000000000000000000000000000000000000000000000000000000000000",
		"df4bc74a3511b4a13ec4848ba50470cc45436078cccab63bda835f13773e7ed3"}

	signature := "d682f8aa34890c9d04ffae83fbd1b7af604b6628f0499e6c4e980a04b1ea0579f724c5cd21157064e8690a776b141d9224d6ac52c1e6b1533e7facd1ebf2e50b"
	sig := SignatureOfMessage(args)
	if sig != signature {
		t.Errorf("Incorrect signature")
	}
}
