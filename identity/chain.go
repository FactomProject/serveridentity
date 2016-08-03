package identity

import (
	"encoding/json"
	"fmt"

	"github.com/FactomProject/factom"
)

type ParamCommitChain struct {
	Message string `json:"message"`
}

func GetChainCommitString(c *factom.Chain, ec *factom.ECAddress) (string, error) {
	jsonString, err := factom.ComposeChainCommit(c, ec)
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
		cc := new(ParamCommitChain)
		err := json.Unmarshal(param, cc)
		if err != nil {
			return "error", err
		}
		return `{"CommitChainMsg":"` + cc.Message + `"}`, nil
	}
	return "", fmt.Errorf("Error: Unreachable code")
}

func GetChainRevealString(c *factom.Chain) (string, error) {
	jsonString, err := factom.ComposeChainReveal(c)
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
		cr := new(ParamEntry)
		err := json.Unmarshal(param, cr)
		if err != nil {
			return "error", err
		}
		return `{"Entry":"` + cr.Entry + `"}`, nil
	}
	return "", fmt.Errorf("Error: Unreachable code")
}

// EC address
// https://github.com/FactomProject/factom/blob/m2-v2/addresses.go
// ComposeChainCommit
// https://github.com/FactomProject/factom/blob/m2-v2/chain.go
