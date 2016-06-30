package identity

// Used by Factomd

import (
	"encoding/json"
)

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
