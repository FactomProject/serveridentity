package handlers

/************************************
 **			Input Sanitation	   **
 ************************************/

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"errors"
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
	case "privStr":
		i := newPrivIN(message, -1)
		return i.Input.ReadIn()
	case "privStrRoot":
		i := newPrivIN(message, 4)
		return i.Input.ReadIn()
	case "privStrLev1":
		i := newPrivIN(message, 1)
		return i.Input.ReadIn()
	case "hexStr":
		i := newHexIN(message, identity.SeedMin, identity.SeedMax)
		return i.Input.ReadIn()
	case "btcAddr":
		i := newBase58IN(message)
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
		rawInput, prefix, err := reader.ReadLine()

		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				fmt.Println("Error in input: EOF. Program exiting...")
				return nil
			}
			fmt.Println("Error in input: " + err.Error())
		} else if prefix != false {
			fmt.Println("Too many characters, exiting")
			return nil
		} else {
			input := string(rawInput[:])
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
type privIN struct {
	ReadInput
	Input *Input
	level int
}

func newPrivIN(message string, lev int) *privIN {
	i := new(privIN)
	i.level = lev
	i.Input = newInput(i, message)
	return i
}

func (i *privIN) sanitize(input string) bool {
	if len(input) == 64 {
		fmt.Println("Must be human readable base format, start with 'sk#', not hex.")
		return false
	} else if len(input) == 53 {
		if strings.Compare(input[:2], "sk") == 0 {
			levInt, err := strconv.Atoi(input[2:3])
			if i.level > 0 && i.level != levInt {
				fmt.Printf("%s%d%s\n", "Not the correct level private key. Please enter the level ", i.level, " key")
				return false
			}
			if err != nil {
				fmt.Println("Error in input: " + err.Error())
			}
			p := base58.Decode(input[:53])
			if !identity.CheckHumanReadable(p[:]) {
				fmt.Println("Not a valid private key, end hash is incorrect.")
				return false
			}

			pShort := p[3:35]
			oByte, err := intToOneByte(levInt)
			if err != nil {
				fmt.Println("Error in input: " + err.Error())
			}
			i.Input.value = append(pShort[:], oByte[:]...)
			return true
		} else {
			fmt.Println("Not a valid private key.")
			return false
		}
	} else {
		fmt.Println("Not a valid private key.")
		return false
	}
}

// Hex String
type hexIN struct {
	ReadInput
	Input *Input
	min   int
	max   int
}

func newHexIN(message string, min int, max int) *hexIN {
	h := new(hexIN)
	h.min = min
	h.max = max
	h.Input = newInput(h, message)
	return h
}

func (i *hexIN) sanitize(input string) bool {
	if len(input)%2 != 0 {
		fmt.Println("Hex string must be of even length.")
		return false
	}

	if len(input) > i.max || len(input) < i.min {
		fmt.Println("String must be between " + strconv.Itoa(i.min) + " and " + strconv.Itoa(i.max) + " length.")
		return false
	}

	if _, err := hex.DecodeString(input); err != nil {
		fmt.Println("Error in input: " + err.Error())
	}

	i.Input.value = input
	return true
}

// Base58 -- BTC Address
type base58IN struct {
	ReadInput
	Input *Input
	min   int
	max   int
}

func newBase58IN(message string) *base58IN {
	b := new(base58IN)
	b.max = 35 // Largest
	b.min = 26 // Smallest
	b.Input = newInput(b, message)
	return b
}

func (b *base58IN) sanitize(input string) bool {
	bType, err := strconv.Atoi(input[:1])
	if err != nil {
		fmt.Println("Error in input: " + err.Error())
	}
	//bType := 1
	if bType == 1 || bType == 3 {

	} else {
		fmt.Println("Invalid bitcoin key, must start with '1' or '3'")
		return false
	}
	addr := base58.Decode(input)
	if len(input) > b.max || len(input) < b.min {
		fmt.Println("Invalid input, incorrect length.")
		return false
	}
	if strings.ContainsAny(input, "OIl0") {
		fmt.Println("Invalid bitcoin address. Cannot contain 'O', '0', 'I', or 'l'.")
		return false
	}
	b.Input.value = addr
	return true
}

/********************************
 *       Helper Functions       *
 ********************************/

func intToOneByte(i int) ([]byte, error) {
	switch i {
	case 1:
		return []byte{0x01}, nil
	case 2:
		return []byte{0x02}, nil
	case 3:
		return []byte{0x03}, nil
	case 4:
		return []byte{0x04}, nil
	}
	return nil, errors.New("Key level must be between 1 and 4")
}
