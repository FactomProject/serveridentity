package handlers

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/functions"
	"github.com/FactomProject/serveridentity/identity"
	"os"
	"strings"
)

var file *os.File

var Start = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity start 'fresh'|ESAddress"
	cmd.description = "Create new identity and subchain."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
		c.HandleFunc("fresh", fresh)
		c.HandleDefaultFunc(existingEC)
		c.HandleFunc("help", func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Create new Identity", cmd)
	return cmd
}()

func existingEC(args []string) {
	if len(args) == 0 {
		Help.All()
		return
	}

	PrintBanner()
	l := len(args[0])
	if l != 64 && l != 52 {
		fmt.Println("server identity start 'fresh'|ESAddress")
		fmt.Println("Invalid ES Address entered, exiting program...")
		return
	} else if l == 52 {
		// Generate all new Keys from EC
		sid := generateKeysFromEC(args[0], true)
		if sid == nil {
			return
		}
		start(sid)
	} else if l == 64 {
		fmt.Println("Only base58 human readable key accepted.")
	}
}

func fresh(args []string) {
	PrintBanner()
	// Generate all new Keys
	sid := generateKeys(true)
	if sid == nil {
		return
	}
	start(sid)
}

func start(sid *functions.ServerIdentity) {
	fmt.Print("If you have copied down these keys, then ")
	err := waitForEnter()
	if err != nil {
		panic(err)
	}

	file = makeFile("startidentity")
	defer file.Close()
	var bar string
	for i := 0; i < 76; i++ {
		bar = bar + "\\*"
	}
	file.WriteString("echo " + bar + "\n")
	file.WriteString("echo \\* Setup script will create and register an identity and its subchain \\ \\ \\ \\ \\ \\ \\*\n")
	file.WriteString("echo \\* Credits must be in " + sid.ECAddr.PubString() + " \\ \\*\n")
	file.WriteString("echo " + bar + "\n")

	PrintHeader("Root Chain Curls")
	createIdentityChain(sid, true)
	registerIdentityChain(sid, true)
	PrintHeader("Sub Chain Curls")
	createSubChain(sid, true)
	registerSubChain(sid, true)
	file.WriteString("echo   \n")
}

func waitForEnter() error {
	fmt.Print("press [ENTER] to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("")
	return nil
}

// Step 1 : Root Chain Create
func createIdentityChain(sid *functions.ServerIdentity, out bool) {
	strCom, strRev, err := functions.CreateIdentityChain(sid)
	if err != nil {
		panic(err)
	}
	if out == true {
		fmt.Println("Root Chain ID: " + sid.RootChainID + "\n")

		fmt.Println(strCom + "\n")
		fmt.Println(strRev + "\n")
	}

	writeCurlCmd(file, "Creating Identity Chain - ChainID: "+sid.RootChainID, strCom, strRev)
}

// Step 2 : Root Chain Register
func registerIdentityChain(sid *functions.ServerIdentity, out bool) {
	strCom, strRev, err := functions.RegisterServerIdentity(sid)
	if err != nil {
		panic(err)
	}
	if out == true {
		fmt.Println(strCom + "\n")
		fmt.Println(strRev + "\n")
	}

	writeCurlCmd(file, "Registering Identity Chain", strCom, strRev)
}

// Step 1 : Subchain Create
func createSubChain(sid *functions.ServerIdentity, out bool) {
	strCom, strRev, err := functions.CreateSubChain(sid)
	if err != nil {
		panic(err)
	}
	if out == true {
		fmt.Println("Sub Chain ID: " + sid.SubChainID + "\n")

		fmt.Println(strCom + "\n")
		fmt.Println(strRev + "\n")
	}

	writeCurlCmd(file, "Creating Server Management SubChain - ChainID: "+sid.SubChainID, strCom, strRev)
}

// Step 2 : Subchain Register
func registerSubChain(sid *functions.ServerIdentity, out bool) {
	strCom, strRev, err := functions.RegisterSubchain(sid)
	if err != nil {
		panic(err)
	}
	if out == true {
		fmt.Println(strCom + "\n")
		fmt.Println(strRev + "\n")
	}

	writeCurlCmd(file, "Registering Server Management SubChain", strCom, strRev)
}

func generateKeys(out bool) *functions.ServerIdentity {
	/*fmt.Print("All new keys for an identity will be created as well as a new Entry \nCredit address. Are you sure you want to continue? y/n :  ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// If not yes, exit
	if strings.Compare("y\n", input) != 0 {
		return nil
	}*/

	// Key generation
	sid, err := functions.MakeServerIdentity()

	if err != nil {
		panic(err)
	}
	if out == true {
		fmt.Println("Key Generation Complete")
		str := formatIDKeysString(sid.IDSet) + "\n"
		str = str + formatECKeyString(sid.ECAddr)

		fmt.Println(str)
	}
	return sid
}

func generateKeysFromEC(ecStr string, out bool) *functions.ServerIdentity {
	/*fmt.Print("All new keys for an identity will be created. Your EC address will \nbe used. Are you sure you want to continue? y/n :  ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// If not yes, exit
	if strings.Compare("y\n", input) != 0 {
		return nil
	}*/

	// Key generation
	if strings.Compare(ecStr[:2], "Es") != 0 {
		fmt.Println("Invalid entry credit private key prefix, exiting program...")
		return nil
	}
	if !factom.IsValidAddress(ecStr) {
		fmt.Println("Invalid entry credit private key, exiting program...")
		return nil
	}
	ec := base58.Decode(ecStr[:52])
	sec := ec[2:34]
	sid, err := functions.MakeServerIdentityFromEC(sec)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
	}
	if out == true {
		fmt.Println("Key Generation Complete")
		str := formatIDKeysString(sid.IDSet) + "\n"
		str = str + formatECKeyString(sid.ECAddr)
		fmt.Println(str)
	}
	return sid
}

func formatIDKeysString(i *identity.IdentitySet) string {
	var str string
	str = "\nPrivate keys and their corresponding levels. Copy these down, this program will \nnot save them for you.\n"
	str = str + "Level 1: " + i.IdentityLevel[0].HumanReadablePrivate() + "\n"
	str = str + "Level 2: " + i.IdentityLevel[1].HumanReadablePrivate() + "\n"
	str = str + "Level 3: " + i.IdentityLevel[2].HumanReadablePrivate() + "\n"
	str = str + "Level 4: " + i.IdentityLevel[3].HumanReadablePrivate() + "\n"

	/*str = str + "\nPublic keys (hex) and their corresponding levels. Copy these down, this program will \nnot save them for you.\n"
	str = str + "Level 1: " + i.IdentityLevel[0].HumanReadablePublic() + "\n"
	str = str + "Level 2: " + i.IdentityLevel[1].HumanReadablePublic() + "\n"
	str = str + "Level 3: " + i.IdentityLevel[2].HumanReadablePublic() + "\n"
	str = str + "Level 4: " + i.IdentityLevel[3].HumanReadablePublic() + "\n"

	str = str + "\nIdentity keys and their corresponding levels. Copy these down, this program will \nnot save them for you.\n"
	str = str + "Level 1: " + i.IdentityLevel[0].HumanReadableIdentity() + "\n"
	str = str + "Level 2: " + i.IdentityLevel[1].HumanReadableIdentity() + "\n"
	str = str + "Level 3: " + i.IdentityLevel[2].HumanReadableIdentity() + "\n"
	str = str + "Level 4: " + i.IdentityLevel[3].HumanReadableIdentity() + "\n"*/
	return str
}

func formatECKeyString(ec *factom.ECAddress) string {
	var str string
	str = "This is the entry key that will be used to add entries/chains to the Factom network. \nCopy these keys down, this program will not save them for you."
	str = str + "\nNote: Entry credits must be in the wallet before commands are executed.\n"
	str = str + "Private Key: " + ec.SecString() + "\n"
	str = str + "Public Key : " + ec.PubString() + "\n"
	return str
}
