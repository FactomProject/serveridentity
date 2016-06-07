package handlers

/************************************
 **			Input Sanitation	   **
 ************************************/

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/btcutil/base58"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
	"os"
	"strconv"
	"strings"
)

func GetInput(inputType string, message string) interface{} {
	switch inputType {
	case "btcKeyLevel":
		i := newIntIN(0, 3, message)
		return i.Input.ReadIn()
	case "btcType":
		i := newIntIN(0, 1, message)
		return i.Input.ReadIn()
	case "chainID":
		i := newChainIDIN(message, 64)
		return i.Input.ReadIn()
	case "ecAddr":
		i := newECAddrIN(message)
		return i.Input.ReadIn()
	}
	// Should never reach
	return nil
}

/********************************
 *     Controls Sanitation      *
 ********************************/
type ReadInput interface {
	sanitize(string) bool
}

type Input struct {
	parent  ReadInput
	message string
	value   interface{}
}

func newInput(parent ReadInput, message string) *Input {
	i := new(Input)
	i.parent = parent
	i.message = message
	return i
}

func (i *Input) ReadIn() interface{} {
	for true {
		fmt.Print(i.message)
		reader := bufio.NewReader(os.Stdin)
		rawInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in input: " + err.Error())
		} else {
			input := rawInput[:len(rawInput)-1]
			if strings.Compare(input, "exit") == 0 {
				return nil
			}
			if i.parent.sanitize(input) == true {
				return i.value
			}
		}
	}
	// Should never get here
	return nil
}

/********************************
 *        Input Structs         *
 ********************************/

// Integer input
type intIN struct {
	Input *Input
	ReadInput
	min int
	max int
}

func newIntIN(min int, max int, message string) *intIN {
	i := new(intIN)
	i.min = min
	i.max = max
	i.Input = newInput(i, message)
	return i
}

func (i *intIN) sanitize(input string) bool {
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error in input: " + err.Error())
		return false
	} else {
		if val < i.min || val > i.max {
			fmt.Println("Value is outside range: must be between " + strconv.Itoa(i.min) + " and " + strconv.Itoa(i.max))
			return false
		} else {
			i.Input.value = val
			return true
		}
	}
}

// ChainID input
type chainIDIN struct {
	Input *Input
	ReadInput
	max int
}

func newChainIDIN(message string, length int) *chainIDIN {
	i := new(chainIDIN)
	i.Input = newInput(i, message)
	i.max = length
	return i
}

func (i *chainIDIN) sanitize(input string) bool {
	if len(input) < 64 {
		fmt.Println("Input not long enough")
		return false
	}
	chainIDbeg := "888888"
	if strings.Compare(input[:identity.ProofOfWorkLength*2], chainIDbeg[:identity.ProofOfWorkLength*2]) != 0 {
		fmt.Println("Invalid identity chain id. Input a valid identity chainID")
		return false
	} else if len(input) > i.max {
		fmt.Println("ChainID is too long, input a valid chainID")
	} else {
		i.Input.value = input
		return true
	}
	return false
}

// EC Address
type ecAddrIN struct {
	ReadInput
	Input *Input
}

func newECAddrIN(message string) *ecAddrIN {
	i := new(ecAddrIN)
	i.Input = newInput(i, message)
	return i
}

func (i *ecAddrIN) sanitize(input string) bool {
	if strings.Compare(input, "any") == 0 { // New EC Key
		_, sec, err := ed.GenerateKey(rand.Reader)
		if err != nil {
			panic(err)
		}
		ecAddr, err := factom.MakeECAddress(sec[:32])
		if err != nil {
			panic(err)
		}
		fmt.Println("New entry credit address. Credits must be in this address before curl commands executed")
		fmt.Println(" * Private Key: " + ecAddr.SecString())
		fmt.Println(" * Public Key : " + ecAddr.SecString())

		i.Input.value = ecAddr
		return true
	} else if len(input) == 52 {
		// Base58 Encoded
		if strings.Compare("Es", input[:2]) != 0 || !factom.IsValidAddress(input[:52]) {
			fmt.Println("Invalid entry credit private address. \nInput the entry credit private address or type 'any' for a random new one, or 'exit' to exit.")
		} else {
			ecAddr, err := factom.MakeECAddress(base58.Decode(input[:52])[2:34])
			if err != nil {
				fmt.Println("Error in input: " + err.Error())
				return false
			}
			i.Input.value = ecAddr
			return true
		}
	} else if len(input) == 64 {
		// Hex
		sec, err := hex.DecodeString(input[:64])
		if err != nil {
			fmt.Println("Error in input: " + err.Error())
			return false
		}
		ecAddr, err := factom.MakeECAddress(sec)
		if err != nil {
			fmt.Println("Error in input: " + err.Error())
			return false
		}
		i.Input.value = ecAddr
		return true
	} else {
		fmt.Println("Invalid input. Input the entry credit private address or type 'any' \nfor a random new one, or 'exit' to exit.")
		return false
	}
	// Should never reach
	return false
}

// Private Key
type privIn struct {
	ReadInput
	Input *Input
}

func newPrivIN(message string) *privIn {
	i := new(privIn)
	return i
}

// TODO MORE
