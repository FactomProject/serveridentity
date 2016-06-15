package handlers

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/serveridentity/functions"
	"os"
)

var Full = func() *sevCmd {
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

	PrintBanner()
	l := len(args[0])
	if l != 64 && l != 52 {
		fmt.Println("server identity full 'fresh'|ESAddress")
		fmt.Println("Invalid ES Address entered, exiting program...")
		return
	} else if l == 52 {
		// Generate all new Keys from EC
		sid := generateKeysFromEC(args[0])
		if sid == nil {
			return
		}
		fullStart(sid)
	} else if l == 64 {
		fmt.Println("Only base58 human readable key accepted.")
	}
}

func freshFull(args []string) {
	PrintBanner()
	// Generate all new Keys
	sid := generateKeys()
	if sid == nil {
		return
	}
	fullStart(sid)
}

func fullStart(sid *functions.ServerIdentity) {
	fmt.Print("If you have copied down these keys, then ")
	err := waitForEnter()
	if err != nil {
		panic(err)
	}

	file = makeFile("fullidentity")
	defer file.Close()
	var bar string
	for i := 0; i < 76; i++ {
		bar = bar + "\\*"
	}
	//file.WriteString("echo " + bar + "\n")
	//file.WriteString("echo \\* Setup script will create and register an identity and its subchain \\ \\ \\ \\ \\ \\ \\*\n")
	//file.WriteString("echo \\* Credits must be in " + sid.ECAddr.PubString() + " \\ \\*\n")
	//file.WriteString("echo " + bar + "\n")

	PrintHeader("Root Chain Curls")
	createIdentityChain(sid)
	registerIdentityChain(sid)
	PrintHeader("Sub Chain Curls")
	createSubChain(sid)
	registerSubChain(sid)
	//file.WriteString("echo   \n")
	btcKeyHex := base58.Decode("1D1biEdmKwq6CVkFPsDkYKry8Ng1opJwM3")

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
	file.WriteString("echo  Identity Chain:" + sid.RootChainID + "\n")
	file.WriteString("echo  Identity Chain:" + sid.SubChainID + "\n")

	file.WriteString("echo EC Public : " + sid.ECAddr.PubString() + "\n")
	file.WriteString("echo EC Private: " + sid.ECAddr.SecString() + "\n")
	file.WriteString("echo  \n")
	file.WriteString("echo  Private Keys\n")
	for i, r := range sid.IDSet.IdentityLevel {
		file.WriteString(fmt.Sprintf("echo  Level %d: %s\n", i, r.HumanReadablePrivate()))
	}
	file.WriteString("echo  \n")
	file.WriteString("echo  BTC Key: 1D1biEdmKwq6CVkFPsDkYKry8Ng1opJwM3\n")
	file.WriteString("echo  Block Signing Key: " + hex.EncodeToString(newPriv) + "\n")
	file.WriteString("echo  \n")
	file.WriteString("echo  MHashSeed: " + sid.RootChainID + "\n")
	file.WriteString("echo  MHash: " + mHash + "\n")

}