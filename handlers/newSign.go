package handlers

import (
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
		c.Handle("block", blockKey)
		c.Handle("btc", btcKey)
		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Create a new signing key", cmd)
	return cmd
}()

var btcKey = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity newkey btc [-s]"
	cmd.description = "Create a new bitcoin key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		sh := flag.Bool("s", false, "generate sh script")
		flag.Parse()
		newBtcKey(*sh)

	}
	Help.Add("Create new bitcoin signing key", cmd)
	return cmd
}()

var blockKey = func() *sevCmd {

	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity newkey block signing key [-s]"
	cmd.description = "Create a new block signing key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		sh := flag.Bool("s", false, "generate sh script")
		flag.Parse()
		newBlockKey(*sh)
	}
	Help.Add("Create new block signing key", cmd)
	return cmd
}()

/********************************
 *        CLI Functions         *
 ********************************/
func newBtcKey(sh bool) {
	PrintBanner()
	if sh == true {
		fmt.Println("A script to run the curl commands will be generated under: 'BtcKey.sh'.")
	}
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

	raw = GetInput("btcAddr", "Input your bitcoin address or type 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	btcKey := raw.([]byte)

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

	// Script Generating
	if sh == true {
		fileB := makeFile("BtcKey")
		defer file.Close()
		writeCurlCmd(fileB, "New Bitcoin Key", strCom, strRev)
	}
}

func newBlockKey(sh bool) {
	if sh == true {
		fmt.Println("A script to run the curl commands will be generated under: 'BlockKey.sh'.")
	}
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
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")

	// Script Generating
	if sh == true {
		fileB := makeFile("BlockKey")
		writeCurlCmd(fileB, "New Bitcoin Key", strCom, strRev)
	}
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
