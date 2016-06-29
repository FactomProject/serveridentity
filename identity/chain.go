package identity

import (
	"encoding/json"
	"os"

	"github.com/FactomProject/factom"
)

func GetChainCommitString(c *factom.Chain, ec *factom.ECAddress) (string, error) {
	js, err := factom.ComposeChainCommit(c, ec)
	if err != nil {
		return "error", err
	}

	// TODO
	// REMOVE START
	t := new(Commit)
	err = MapToObject(js, t)
	com, _ := os.OpenFile("com.txt", os.O_RDWR|os.O_APPEND, 0660)
	com.WriteString(NUMBER + "#" + t.Params.Message + "#")
	com.Close()
	// END
	//TODO

	str, err := factom.EncodeJSONString(js)
	if err != nil {
		return "error", err
	}
	return str, nil
}

func GetChainRevealString(c *factom.Chain) (string, error) {
	js, err := factom.ComposeChainReveal(c)
	if err != nil {
		return "error", err
	}

	// TODO
	// REMOVE START
	t := new(Reveal)
	err = MapToObject(js, t)
	rev, _ := os.OpenFile("rev.txt", os.O_RDWR|os.O_APPEND, 0660)
	rev.WriteString(NUMBER + "#" + t.Params.Message + "#")
	rev.Close()
	// END
	//TODO

	str, err := factom.EncodeJSONString(js)
	if err != nil {
		return "error", err
	}
	return str, nil
}

// EC address
// https://github.com/FactomProject/factom/blob/m2-v2/addresses.go
// ComposeChainCommit
// https://github.com/FactomProject/factom/blob/m2-v2/chain.go
type Commit struct {
	Jsonrpc string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Params  *ParamC `json:"params"`
	Method  string  `json:"method"`
}
type Reveal struct {
	Jsonrpc string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Params  *ParamR `json:"params"`
	Method  string  `json:"method"`
}
type ParamC struct {
	Message string `json:"message"`
}
type ParamR struct {
	Message string `json:"entry"`
}

func MapToObject(source interface{}, dst interface{}) error {
	b, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}
