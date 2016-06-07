package handlers

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/functions"
	"github.com/FactomProject/serveridentity/utils"
	"os"
	"strconv"
	"strings"
)

var NewKey = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "serveridentity newkey 'block'|'btc'"
	cmd.description = "Create new block signing key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
		c.HandleFunc("block", newBlockKey)
		c.HandleFunc("btc", newBtcKey)
		c.Execute(args)
	}
	Help.Add("Create new block signing key", cmd)
	return cmd
}()

//
func newBtcKey(args []string) {
	var raw interface{}
	fmt.Println("To create a new bitcoin key multiple inputs will be required.")

	raw = GetInput("chainID", "Input the identity chain ID in hex or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	rootID := raw.(string)

	raw = GetInput("chainID", "Input the subchain ID in hex or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	subChainID := raw.(string)

	raw = GetInput("btcKeyLevel", "Input the bitcoin key level or type 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	btcKeyLevel := raw.(int)

	raw = GetInput("btcType", "Input the bitcoin address type (0=P2PKH 1=P2SH) or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	btcType := raw.(int)

	// TODO: CONVERT
	btcKey := getBase58("Input your bitcoin address or type 'exit':  ")
	if btcKey == nil { // Exit case
		return
	}

	_, privKey := getPrivateKey("Input the human readable base58 private key or type 'exit':  ")
	if privKey == nil { // Exit case
		return
	}
	// END CONVERT

	raw = GetInput("ecAddr", "Input the entry credit address or 'any' for a new one:  ")
	if raw == nil { // Exit case
		return
	}
	ecAddr := raw.(*factom.ECAddress)
	fmt.Println("Your public entry credit address is: \n * " + ecAddr.PubString())

	strCom, strRev, err := functions.CreateNewBitcoinKey(rootID, subChainID, btcKeyLevel, btcType, btcKey, privKey, ecAddr)
	if err != nil {
		panic(err)
	}
	PrintHeader("New Bitcoin Key Curl Commands")
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

func newBlockKey(args []string) {
	var raw interface{}
	fmt.Println("To create a new block signing key multiple inputs will be required.")

	raw = GetInput("chainID", "Input the identity chain ID in hex or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	rootID := raw.(string)

	raw = GetInput("ecAddr", "Input the entry credit address or 'any' for a new one:  ")
	if raw == nil { // Exit case
		return
	}
	ecAddr := raw.(*factom.ECAddress)
	fmt.Println("Your public entry credit address is: \n * " + ecAddr.PubString())

	// Just 32 bytes
	lev, privKeyAbove := getPrivateKey("Input the private key of the level below you wish to replace. \nHumanReadable base58 key expected, or type 'exit' to exit")
	if privKeyAbove == nil { // Exit case
		return
	}
	// END CONVERT

	strCom, strRev, newPriv, err := functions.CreateNewBlockSignEntry(rootID, privKeyAbove, ecAddr)
	if err != nil {
		panic(err)
	}

	PrintHeader("New Block Signing Key Curl Commands")
	fmt.Println("New PrivateKey : " + newHumanReadable(lev, newPriv) + "\n")
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

// TODO: Move all code below to readInput.go
func newHumanReadable(lev int, key []byte) string {
	var prefix []byte
	switch lev {
	case 1:
		// Case 1 should never happen
		prefix = []byte{0x4d, 0xb6, 0xc9}
	case 2:
		prefix = []byte{0x4d, 0xb6, 0xe7}
	case 3:
		prefix = []byte{0x4d, 0xb7, 0x05}
	case 4:
		prefix = []byte{0x4d, 0xb7, 0x23}
	}

	buf := new(bytes.Buffer)
	// Add Prefix
	buf.Write(prefix[:])

	// Add Key
	buf.Write(key[:32])

	o := buf.Bytes()
	// Sha356d
	humanReadable := utils.Sha256d(o)

	// Append first 4 bytes
	o = append(o, humanReadable[:4]...)

	str := base58.Encode(o)
	return str
}

func getBase58(message string) []byte {
	for true {
		fmt.Print(message)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in input: " + err.Error())
		} else if strings.Compare(input[:len(input)-1], "exit") == 0 {
			return nil
		} else {
			// TODO: Sanitize
			b := base58.Decode(input[:len(input)-1])
			return b
		}
	}
	// should never reach here
	return nil
}

// TODO: Santitize better and catch errors
func getPrivateKey(message string) (int, []byte) {
	//fmt.Println("Input the private key of the level below you wish to replace. \nHumanReadable base58 key expected, or type 'exit' to exit")
	fmt.Println(message)
	for true {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.Compare(input, "exit\n") == 0 { // Exit
			return 0, nil
		} else if len(input) == 65 {
			fmt.Println("Must be human readable base format, start with 'sk#', not hex.")
		} else if len(input) == 54 {
			if strings.Compare(input[:2], "sk") == 0 {
				levInt, err := strconv.Atoi(input[3:4])
				if err != nil {
					panic(err)
				}
				levInt++
				// TODO: Check valid human readable hash at end
				p := base58.Decode(input[:53])
				return levInt, p[3:35]
			} else {
				fmt.Println("Not a private key, input the private key of the level below you wish to replace. \nHumanReadable base58 ley expected, or type 'exit' to exit")
			}
		} else {
			fmt.Println("Invalid input, a the private key. HumanReadable base58 key expected, or type 'exit' to exit")
		}
	}
	// Should never get here
	return 0, nil
}
