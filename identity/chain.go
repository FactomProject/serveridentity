package identity

import (
	"github.com/FactomProject/factom"
)

func GetChainCommitString(c *factom.Chain, ec *factom.ECAddress) (string, error) {
	json, err := factom.ComposeChainCommit(c, ec)
	if err != nil {
		return "error", err
	}

	str, err := factom.EncodeJSONString(json)
	if err != nil {
		return "error", err
	}
	return str, nil
}

func GetChainRevealString(c *factom.Chain) (string, error) {
	json, err := factom.ComposeChainReveal(c)
	if err != nil {
		return "error", err
	}

	str, err := factom.EncodeJSONString(json)
	if err != nil {
		return "error", err
	}
	return str, nil
}

// EC address
// https://github.com/FactomProject/factom/blob/m2-v2/addresses.go
// ComposeChainCommit
// https://github.com/FactomProject/factom/blob/m2-v2/chain.go
