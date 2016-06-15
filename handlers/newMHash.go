package handlers

import (
	"flag"
	"fmt"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/functions"
	"os"
)

/********************************
 *          Cli Control         *
 ********************************/
var NewMHash = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity mhash"
	cmd.description = "Create a new Matryoshka Hash"
	cmd.execFunc = func(args []string) {
		os.Args = args
		sh := flag.Bool("s", false, "generate sh script")
		flag.Parse()
		newMHash(*sh)
	}
	Help.Add("Create a new Matryoshka Hash", cmd)
	return cmd
}()

/********************************
 *        CLI Functions         *
 ********************************/
func newMHash(sh bool) {
	PrintBanner()
	if sh == true {
		fmt.Println("A script to run the curl commands will be generated under: 'mhash.sh'.")
	}
	var raw interface{}
	fmt.Println("To create a new Matryoshka Hash multiple inputs will be required.")

	raw = GetInput("chainID", "Input the identity chain ID in hex or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	rootID := raw.(string)

	raw = GetInput("chainID", "Input the server management chainID in hex or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	subChainID := raw.(string)

	raw = GetInput("hexStr", "Input a random hex seed or 'exit':  ")
	if raw == nil { // Exit case
		return
	}
	seed := raw.(string)

	raw = GetInput("privStrLev1", "Input the level 1 ('sk1') private key. HumanReadable base58 key expected, or type 'exit':  \n")
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

	strCom, strRev, mHash, err := functions.CreateNewMHash(rootID, subChainID, privKey, seed, ecAddr)
	if err != nil {
		panic(err)
	}

	PrintHeader("New MHash Curl Commands")
	fmt.Println("New MHash : " + mHash + "\n")
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")

	if sh == true {
		fileB := makeFile("mhash")
		defer file.Close()
		writeCurlCmd(fileB, "New Matryoshka Hash", strCom, strRev)
	}
}
