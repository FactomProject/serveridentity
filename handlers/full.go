package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/serveridentity/functions"
	"github.com/FactomProject/serveridentity/identity"
	"io"
	"os"
)

var SCRIPTNAME string = "fullidentity"
var PRINT_OUT bool = true

/*
 * This file is only used for testing purposes
 */

var Full = func() *sevCmd {
	identity.ShowBruteForce = PRINT_OUT
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity full 'fresh'|ESAddress|elements"
	cmd.description = "Create new identity and subchain as well as entries in the subchain."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
		c.HandleFunc("elements", elementsFull)
		c.HandleFunc("fresh", freshFull)
		c.HandleDefaultFunc(existingECFull)
		c.HandleFunc("help", func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Create a full Identity", cmd)
	return cmd
}()

func existingECFull(args []string) {
	if len(args) == 0 {
		Help.All()
		return
	}
	os.Args = args
	filename := flag.String("n", "fullidentity", "Change the script name")
	flag.Parse()
	SCRIPTNAME = *filename
	l := len(args[0])
	if l != 64 && l != 52 {
		fmt.Println("serveridentity full 'fresh'|ESAddress")
		fmt.Println("Invalid ES Address entered, exiting program...")
		return
	} else if l == 52 {
		// Generate all new Keys from EC
		sid := generateKeysFromEC(args[0], PRINT_OUT)
		if sid == nil {
			return
		}
		fullStart(sid)
	} else if l == 64 {
		fmt.Println("Only base58 human readable key accepted.")
	}
}

func freshFull(args []string) {
	os.Args = args
	filename := flag.String("n", "fullidentity", "Change the script name")
	flag.Parse()
	SCRIPTNAME = *filename
	// Generate all new Keys
	sid := generateKeys(PRINT_OUT)
	if sid == nil {
		return
	}
	fullStart(sid)
}

func elementsFull(args []string) {
	if len(args) > 1 {
		os.Args = args[1:]
	} else {
		os.Args = args
	}
	filename := flag.String("n", "fullidentity", "Change the script name")
	flag.Parse()
	SCRIPTNAME = *filename
	var sid *functions.ServerIdentity

	if len(args) > 1 {
		fmt.Println(args[1])
		l := len(args[1])
		if l != 52 {
			fmt.Println("serveridentity full elements ESAddress")
			fmt.Println("Invalid ES Address entered, exiting program...")
			return
		} else {
			// Generate all new Keys from EC
			sid = generateKeysFromEC(args[1], PRINT_OUT)
			if sid == nil {
				return
			}
		}
	} else {
		sid = generateKeysFromEC("Es2Rf7iM6PdsqfYCo3D1tnAR65SkLENyWJG1deUzpRMQmbh9F3eG", PRINT_OUT)
		if sid == nil {
			return
		}
	}
	fullStartElements(sid)
}

func fullStart(sid *functions.ServerIdentity) {
	file := makeFile(SCRIPTNAME + ".sh")
	defer file.Close()
	var bar string
	for i := 0; i < 76; i++ {
		bar = bar + "\\*"
	}
	if PRINT_OUT {
		PrintHeader("Root Chain Curls")
	}
	createIdentityChain(sid, PRINT_OUT, file)
	registerIdentityChain(sid, PRINT_OUT, file)
	if PRINT_OUT {
		PrintHeader("Sub Chain Curls")
	}
	createSubChain(sid, PRINT_OUT, file)
	registerSubChain(sid, PRINT_OUT, file)

	random := rand.Reader
	var r [20]byte
	_, _ = io.ReadFull(random, r[:20])
	btcKeyHex := r[:20]

	p := sid.IDSet.IdentityLevel[0].GetPrivateKey()
	priv := p[:32]

	file.WriteString("sleep 1\n")

	strCom, strRev, err := functions.CreateNewBitcoinKey(sid.RootChainID, sid.SubChainID, 0, 0, btcKeyHex, priv, sid.ECAddr)
	if err != nil {
		panic(err)
	}
	writeCurlCmd(file, "New Bitcoin Key", strCom, strRev)

	strCom, strRev, newPriv, err := functions.CreateNewBlockSignEntry(sid.RootChainID, sid.SubChainID, priv, sid.ECAddr)
	if err != nil {
		panic(err)
	}
	writeCurlCmd(file, "New Block Signing Key", strCom, strRev)

	strCom, strRev, mHash, err := functions.CreateNewMHash(sid.RootChainID, sid.SubChainID, priv, sid.RootChainID, sid.ECAddr)
	if err != nil {
		panic(err)
	}
	writeCurlCmd(file, "New Matryoshka Hash", strCom, strRev)

	file.WriteString("echo " + bar + "\n")
	file.WriteString("echo  Identity Info\n")
	file.WriteString("echo " + bar + "\n")
	file.WriteString("echo  Identity Chain: " + sid.RootChainID + "\n")
	file.WriteString("echo  Identity SubChain: " + sid.SubChainID + "\n")

	file.WriteString("echo EC Public : " + sid.ECAddr.PubString() + "\n")
	file.WriteString("echo EC Private: " + sid.ECAddr.SecString() + "\n")
	file.WriteString("echo  \n")
	file.WriteString("echo  Private Keys\n")
	for i, r := range sid.IDSet.IdentityLevel {
		file.WriteString(fmt.Sprintf("echo  Level %d: %s\n", i+1, r.HumanReadablePrivate()))
	}
	file.WriteString("echo  \n")
	file.WriteString("echo  BTC Key: " + hex.EncodeToString(btcKeyHex) + "\n")
	keyString := hex.EncodeToString(newPriv)
	keyString = "\n echo - Sec: " + keyString[:64] + "\n echo - Pub: " + keyString[64:]
	file.WriteString("echo  Block Signing Key: " + keyString + "\n")
	file.WriteString("echo  \n")
	file.WriteString("echo  MHashSeed: " + sid.RootChainID + "\n")
	file.WriteString("echo  MHash: " + mHash + "\n")
}

func cliFormat(cliCommand string, ECaddress string) string {
	cliLine := "echo -n \"\" | factom-cli "
	cliLine += cliCommand
	cliLine += " "
	cliLine += ECaddress

	return cliLine
}

func fullStartElements(sid *functions.ServerIdentity) {
	file := makeFile(SCRIPTNAME + ".sh")
	defer file.Close()
	configfile := makeFile(SCRIPTNAME + ".config")
	defer configfile.Close()

	PrintHeader("Creating Identity Chains/Keys")
	ice, err := functions.CreateIdentityChainElements(sid)
	if err != nil {
		panic(err)
	}
	fice := cliFormat(ice, sid.ECAddr.String())

	icr, err := functions.RegisterServerIdentityElements(sid)
	if err != nil {
		panic(err)
	}
	ficr := cliFormat(icr, sid.ECAddr.String())

	sce, err := functions.CreateSubChainElements(sid)
	if err != nil {
		panic(err)
	}
	fsce := cliFormat(sce, sid.ECAddr.String())

	scr, err := functions.RegisterSubChainElements(sid)
	if err != nil {
		panic(err)
	}
	fscr := cliFormat(scr, sid.ECAddr.String())

	fmt.Println("Root Chain: " + sid.RootChainID)
	fmt.Println("Management Chain: " + sid.RootChainID)
	fmt.Println()

	p := sid.IDSet.IdentityLevel[0].GetPrivateKey()
	lowestLevelSigningKey := p[:32]
	lowestLevelSigningKeyHex := fmt.Sprintf("%032x", lowestLevelSigningKey)

	// Block signing key
	bse, bsPriv, bsPublic, err := functions.CreateNewBlockSignEntryElements(sid)
	if err != nil {
		panic(err)
	}

	unsignedUntimedBse, _ := functions.CreateNewBlockSignEntryUnsigned(sid, bsPublic)

	fbse := cliFormat(bse, sid.ECAddr.String())

	bsPrivHex := fmt.Sprintf("%032x", bsPriv)
	fmt.Printf("block signing private key: %s\n", bsPrivHex)

	bsPubHex := fmt.Sprintf("%032x", bsPublic)
	fmt.Printf("block signing public key: %s\n", bsPubHex)
	fmt.Println()

	// Create a Bitcoin Key
	random := rand.Reader
	var r [20]byte
	_, _ = io.ReadFull(random, r[:20])
	btcKeyHex := r[:20]

	fmt.Printf("BTC Key: %x\n", btcKeyHex)

	bke, err := functions.CreateNewBitcoinKeyElements(sid.RootChainID, sid.SubChainID, 0, 0, btcKeyHex, lowestLevelSigningKey, sid.ECAddr)
	if err != nil {
		panic(err)
	}

	fbke := cliFormat(bke, sid.ECAddr.String())

	unsignedUntimesBKe, err := functions.CreateNewBitcoinKeyElementsUnsigned(sid.RootChainID, sid.SubChainID, 0, 0, btcKeyHex, lowestLevelSigningKey, sid.ECAddr)
	if err != nil {
		panic(err)
	}

	// New MHash
	randomSeed := rand.Reader
	var rS [20]byte
	_, _ = io.ReadFull(randomSeed, rS[:20])
	seed := fmt.Sprintf("%x", rS[:20])

	fmt.Printf("MHash Seed (hex): %x\n", seed)

	mhe, err := functions.CreateNewMHashElements(sid.RootChainID, sid.SubChainID, lowestLevelSigningKey, seed, sid.ECAddr)
	if err != nil {
		panic(err)
	}

	fmhe := cliFormat(mhe, sid.ECAddr.String())

	unsignedUntimesMHe, err := functions.CreateNewMHashElementsUnsigned(sid.RootChainID, sid.SubChainID, lowestLevelSigningKey, seed, sid.ECAddr)
	if err != nil {
		panic(err)
	}

	PrintHeader("Factom-cli commands")
	/*****************************
	 * Begin factom-cli commands *
	 *****************************/

	// factom-cli commands to be run will be outputted to the script. Default name is 'fullidentity.sh'
	fileText := ""

	// "Identity Chain"
	fmt.Println(fice)
	fileText += fice + "\n"
	// "Register Factom Identity"
	fmt.Println(ficr)
	fileText += ficr + "\n"
	// "Server Management"
	fmt.Println(fsce)
	fileText += fsce + "\n"
	// "Register Server Management"
	fmt.Println(fscr)
	fileText += fscr + "\n"

	// Declare now
	nowBash := "now=$(printf '%016x' $(date +%s))"
	fileText += nowBash + "\n"
	// Block signing key
	fmt.Println(nowBash)
	sigBashBlock := fmt.Sprintf("sig=$(signwithed25519 %s$now %s)\n", unsignedUntimedBse, lowestLevelSigningKeyHex)
	fmt.Printf(sigBashBlock)
	fmt.Println(fbse)

	fileText += sigBashBlock
	fileText += fbse + "\n"

	// Bitcoin Key
	sigBashBTC := fmt.Sprintf("sigBTC=$(signwithed25519 %s$now %s)\n", unsignedUntimesBKe, lowestLevelSigningKeyHex)
	fmt.Printf(sigBashBTC)
	fmt.Println(fbke)

	fileText += sigBashBTC
	fileText += fbke + "\n"

	// MHash
	sigBashMHash := fmt.Sprintf("sigMHASH=$(signwithed25519 %s$now %s)\n", unsignedUntimesMHe, lowestLevelSigningKeyHex)
	fmt.Printf(sigBashMHash)
	fmt.Println(fmhe)

	fileText += sigBashMHash
	fileText += fmhe + "\n"

	// Write fileText to file
	file.WriteString(fileText)

	configFileText := "IdentityChainID                       = "
	configFileText += sid.RootChainID
	configFileText += "\nLocalServerPrivKey                    = "
	configFileText += bsPrivHex
	configFileText += "\nLocalServerPublicKey                  = "
	configFileText += bsPubHex
	configFileText += "\n"
	configfile.WriteString(configFileText)

}
