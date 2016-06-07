package identity

import (
	"github.com/FactomProject/factom"
)

func GetEntryCommitString(e *factom.Entry, ec *factom.ECAddress) (string, error) {
	json, err := factom.ComposeEntryCommit(e, ec)
	if err != nil {
		return "error", err
	}

	str, err := factom.EncodeJSONString(json)
	if err != nil {
		return "error", err
	}
	return str, nil
}

func GetEntryRevealString(e *factom.Entry) (string, error) {
	json, err := factom.ComposeEntryReveal(e)
	if err != nil {
		return "error", err
	}

	str, err := factom.EncodeJSONString(json)
	if err != nil {
		return "error", err
	}
	return str, nil
}
