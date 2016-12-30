package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/FactomProject/cli"
	ed "github.com/FactomProject/ed25519"
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
	os.Args = args
	filename := flag.String("n", "fullidentity", "Change the script name")
	flag.Parse()
	SCRIPTNAME = *filename
	var sid *functions.ServerIdentity

	if len(args) > 1 {
		fmt.Println(args[1])
		l := len(args[1])
		if l != 52 {
			fmt.Println("serveridentity elements ESAddress")
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
	file = makeFile(SCRIPTNAME)
	defer file.Close()
	var bar string
	for i := 0; i < 76; i++ {
		bar = bar + "\\*"
	}
	if PRINT_OUT {
		PrintHeader("Root Chain Curls")
	}
	createIdentityChain(sid, PRINT_OUT)
	registerIdentityChain(sid, PRINT_OUT)
	if PRINT_OUT {
		PrintHeader("Sub Chain Curls")
	}
	createSubChain(sid, PRINT_OUT)
	registerSubChain(sid, PRINT_OUT)

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
	file = makeFile(SCRIPTNAME)
	defer file.Close()
	//	var bar string
	if PRINT_OUT {
		//PrintHeader("Root Chain Curls")
	}
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

	p := sid.IDSet.IdentityLevel[0].GetPrivateKey()
	lowestLevelSigningKey := p[:32]
	lowestLevelSigningKeyHex := fmt.Sprintf("%032x", lowestLevelSigningKey)

	bse, bsPriv, err := functions.CreateNewBlockSignEntryElements(sid)
	if err != nil {
		panic(err)
	}

	unsignedUntimedBse, _ := functions.CreateNewBlockSignEntryUnsigned(sid, bsPriv)

	fbse := cliFormat(bse, sid.ECAddr.String())

	bsPrivHex := fmt.Sprintf("%032x", bsPriv)
	fmt.Printf("block signing private key: %s\n", bsPrivHex)

	var priv [64]byte
	copy(priv[:32], bsPriv[:32])
	bsPub := ed.GetPublicKey(&priv)
	fmt.Printf("block signing public key: %032x\n", *bsPub)
	fmt.Println()

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

	// Block signing key
	nowBash := "now=$(printf '%016x' $(date +%s))"
	fmt.Println(nowBash)
	sigBash := fmt.Sprintf("sig=$(signwithed25519 %s$now %s)\n", unsignedUntimedBse, lowestLevelSigningKeyHex)
	fmt.Printf(sigBash)
	fmt.Println(fbse)

	fileText += nowBash + "\n"
	fileText += sigBash
	fileText += fbse + "\n"

	// Write fileText to file
	file.WriteString(fileText)

	//strCom, strRev, newPriv, err := functions.CreateNewBlockSignEntry(sid.RootChainID, sid.SubChainID, priv, sid.ECAddr)

	//modified to here so far

	/*

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

	*/

}
