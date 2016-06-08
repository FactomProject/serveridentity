package handlers

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/functions"
	"github.com/FactomProject/serveridentity/utils"
	"os"
	"strings"
)

/********************************
 *          Cli Control         *
 ********************************/
var NewKey = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity newkey 'block'|'btc'"
	cmd.description = "Create a new key to add/replace"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
		c.HandleFunc("block", newBlockKey)
		c.HandleFunc("btc", newBtcKey)
		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Create new block signing key", cmd)
	return cmd
}()

/********************************
 *        CLI Functions         *
 ********************************/
func newBtcKey(args []string) {
	PrintBanner()
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
	// END CONVERT

	raw = GetInput("privStr", "Input the private key of the level below you wish to replace. \nHumanReadable base58 key expected, or type 'exit':  \n")
	if raw == nil { // Exit case
		return
	}
	privKey := raw.([]byte)[:32]

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
	PrintBanner()
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
	fmt.Println(" -  Your public entry credit address is: \n * " + ecAddr.PubString())

	// Just 32 bytes
	raw = GetInput("privStr", "Input the private key to sign. HumanReadable base58 key expected, or type 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	privKeyAbove := raw.([]byte)[:32]
	//lev := raw.([]byte)[32:33] // No longer used, is level of key signing

	strCom, strRev, newPriv, err := functions.CreateNewBlockSignEntry(rootID, privKeyAbove, ecAddr)
	if err != nil {
		panic(err)
	}

	PrintHeader("New Block Signing Key Curl Commands")
	fmt.Println("New PrivateKey : " + hex.EncodeToString(newPriv)[:32] + "\n")
	// makeHumanReadable(lev, newPriv) + "\n")
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

/********************************
 *       Helper Functions       *
 ********************************/

// No longer used
func makeHumanReadable(lev []byte, key []byte) string {
	var prefix []byte
	if bytes.Compare([]byte{0x01}, lev) == 0 {
		// Case 1 should never happen
		prefix = []byte{0x4d, 0xb6, 0xc9}
	} else if bytes.Compare([]byte{0x02}, lev) == 0 {
		prefix = []byte{0x4d, 0xb6, 0xe7}
	} else if bytes.Compare([]byte{0x03}, lev) == 0 {
		prefix = []byte{0x4d, 0xb7, 0x05}
	} else if bytes.Compare([]byte{0x04}, lev) == 0 {
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

// TODO: Move all code below to readInput.go
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
