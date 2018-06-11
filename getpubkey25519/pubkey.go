package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	ed "github.com/FactomProject/ed25519"
)

var _ = fmt.Sprintf("")

func main() {
	flag.Parse()

	// Arg[1] = {32_BYTE_PRIVATE_KEY}
	args := flag.Args()

	if err := CheckArgs(args); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(PubkeyOfPrivate(args))
}

// Checks to make sure args are correct, and returns an error if they are not.
// Just some extra validation
func CheckArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("ERROR: Only one paramater (32 bytes hex privatekey) is allowed.")
	}

	if len(args[0]) != 64 {
		return fmt.Errorf("ERROR: Private key must be 32 bytes or 64 hex characters.")
	}

	return nil
}

// Sign message and return signature as string
func PubkeyOfPrivate(args []string) string {
	sec, err := hex.DecodeString(args[0])
	if err != nil {
		return fmt.Sprintf("ERROR: Message is invalid hex encoding: %v", err.Error())
	}

	if len(sec) != 32 {
		return fmt.Sprintf("ERROR: Private key must be 32 bytes")
	}

	privateKey := new([ed.PrivateKeySize]byte)
	copy(privateKey[:32], sec[:32])

	// Get public key bytes
	pubkey := ed.GetPublicKey(privateKey)

	pubkeyString := hex.EncodeToString((*pubkey)[:])

	return pubkeyString
}
