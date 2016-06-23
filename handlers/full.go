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
	cmd.helpMsg = "serveridentity full 'fresh'|ESAddress"
	cmd.description = "Create new identity and subchain as well as entries in the subchain."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
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
	garble := flag.Bool("b", false, "Make incorrect entries")
	flag.Parse()
	SCRIPTNAME = *filename
	l := len(args[0])
	if l != 64 && l != 52 {
		fmt.Println("server identity full 'fresh'|ESAddress")
		fmt.Println("Invalid ES Address entered, exiting program...")
		return
	} else if l == 52 {
		// Generate all new Keys from EC
		sid := generateKeysFromEC(args[0], PRINT_OUT)
		if sid == nil {
			return
		}
		fullStart(sid, *garble)
	} else if l == 64 {
		fmt.Println("Only base58 human readable key accepted.")
	}
}

func freshFull(args []string) {
	os.Args = args
	filename := flag.String("n", "fullidentity", "Change the script name")
	garble := flag.Bool("b", false, "Make incorrect entries")
	flag.Parse()
	SCRIPTNAME = *filename
	// Generate all new Keys
	sid := generateKeys(PRINT_OUT)
	if sid == nil {
		return
	}
	fullStart(sid, *garble)
}

func fullStart(sid *functions.ServerIdentity, garble bool) {
	if garble {
		fmt.Println("Incorrect curls also provided")
	}
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
	file.WriteString("echo  Block Signing Key: " + hex.EncodeToString(newPriv) + "\n")
	file.WriteString("echo  \n")
	file.WriteString("echo  MHashSeed: " + sid.RootChainID + "\n")
	file.WriteString("echo  MHash: " + mHash + "\n")

	if garble {
		PrintHeader("GARBLE: Wrong Key")
		p = sid.IDSet.IdentityLevel[2].GetPrivateKey()
		priv = p[:32]

		file.WriteString("sleep 1\n")

		strCom, strRev, err = functions.CreateNewBitcoinKey(sid.RootChainID, sid.SubChainID, 0, 0, btcKeyHex, priv, sid.ECAddr)
		if err != nil {
			//panic(err)
		}
		writeCurlCmd(file, "New Bitcoin Key", strCom, strRev)

		strCom, strRev, newPriv, err = functions.CreateNewBlockSignEntry(sid.RootChainID, sid.SubChainID, priv, sid.ECAddr)
		if err != nil {
			//panic(err)
		}
		writeCurlCmd(file, "New Block Signing Key", strCom, strRev)

		strCom, strRev, mhash, err := functions.CreateNewMHash(sid.RootChainID, sid.SubChainID, priv, sid.RootChainID, sid.ECAddr)
		if err != nil {
			//panic(err)
		}
		writeCurlCmd(file, "New Matryoshka Hash", strCom, strRev)

		PrintHeader("GARBLE: Bad Key & BTC KEY")
		btcKeyHex = []byte{0x00, 0x00, 0x00}
		p = sid.IDSet.IdentityLevel[0].GetPrivateKey()
		priv = p[1:33]

		file.WriteString("sleep 1\n")

		strCom, strRev, err = functions.CreateNewBitcoinKey(sid.RootChainID, sid.SubChainID, 0, 0, btcKeyHex, priv, sid.ECAddr)
		if err != nil {
			//panic(err)
		}
		writeCurlCmd(file, "New Bitcoin Key", strCom, strRev)

		strCom, strRev, newPriv, err = functions.CreateNewBlockSignEntry(sid.RootChainID, sid.SubChainID, priv, sid.ECAddr)
		if err != nil {
			//panic(err)
		}
		writeCurlCmd(file, "New Block Signing Key", strCom, strRev)

		strCom, strRev, mhash, err = functions.CreateNewMHash(sid.RootChainID, sid.SubChainID, priv, sid.RootChainID, sid.ECAddr)
		if err != nil {
			//panic(err)
		}
		writeCurlCmd(file, "New Matryoshka Hash", strCom, strRev)
		file.WriteString("echo  \n")
		file.WriteString("echo  BTC Key: " + hex.EncodeToString(btcKeyHex) + "\n")
		file.WriteString("echo  Block Signing Key: " + hex.EncodeToString(newPriv) + "\n")
		file.WriteString("echo  MHash: " + mhash + "\n")
	}

}
