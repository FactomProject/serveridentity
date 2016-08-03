package identity

import (
	"encoding/json"
	"fmt"

	"github.com/FactomProject/factom"
)

type ParamEntry struct {
	Entry string `json:"entry"`
}

func GetEntryCommitString(e *factom.Entry, ec *factom.ECAddress) (string, error) {
	jsonString, err := factom.ComposeEntryCommit(e, ec)
	if err != nil {
		return "error", err
	}

	// Milestone 2
	if Version == 2 {
		str, err := factom.EncodeJSONString(jsonString)
		if err != nil {
			return "error", err
		}
		return str, nil
	}
	// Milestone 1
	param, err := jsonString.Params.MarshalJSON()
	if err != nil {
		return "error", err
	} else {
		ec := new(ParamEntry)
		err := json.Unmarshal(param, ec)
		if err != nil {
			return "error", err
		}
		return `{"CommitEntryMsg":"` + ec.Entry + `"}`, nil
	}
	return "", fmt.Errorf("Error: Unreachable code")
}

func GetEntryRevealString(e *factom.Entry) (string, error) {
	jsonString, err := factom.ComposeEntryReveal(e)
	if err != nil {
		return "error", err
	}

	// Milestone 2
	if Version == 2 {
		str, err := factom.EncodeJSONString(jsonString)
		if err != nil {
			return "error", err
		}
		return str, nil
	}
	// Milestone 1
	param, err := jsonString.Params.MarshalJSON()
	if err != nil {
		return "error", err
	} else {
		er := new(ParamEntry)
		err := json.Unmarshal(param, er)
		if err != nil {
			return "error", err
		}
		return `{"Entry":"` + er.Entry + `"}`, nil
	}
	return "", fmt.Errorf("Error: Unreachable code")
}
