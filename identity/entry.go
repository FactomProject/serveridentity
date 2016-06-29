package identity

import (
	"os"

	"github.com/FactomProject/factom"
)

func GetEntryCommitString(e *factom.Entry, ec *factom.ECAddress) (string, error) {
	json, err := factom.ComposeEntryCommit(e, ec)
	if err != nil {
		return "error", err
	}
	// TODO REMOVE
	// START
	t := new(Reveal)
	err = MapToObject(json, t)
	com, _ := os.OpenFile("ecom.txt", os.O_RDWR|os.O_APPEND, 0660)
	com.WriteString(NUMBER + "#" + t.Params.Message + "#")
	com.Close()
	//END
	//TODO

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
	// TODO REMOVE
	// START
	t := new(Reveal)
	err = MapToObject(json, t)
	rev, _ := os.OpenFile("erev.txt", os.O_RDWR|os.O_APPEND, 0660)
	rev.WriteString(NUMBER + "#" + t.Params.Message + "#")
	rev.Close()
	//END
	//TODO
	str, err := factom.EncodeJSONString(json)
	if err != nil {
		return "error", err
	}
	return str, nil
}
