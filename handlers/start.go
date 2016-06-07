package handlers

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/functions"
	"github.com/FactomProject/serveridentity/identity"
	"os"
	"strings"
)

var Start = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "serveridentity start 'fresh'|ESAddress"
	cmd.description = "Create new identity and subchain."
	cmd.execFunc = func(args []string) {
		if len(args) > 1 && strings.Compare(args[1], "fresh") == 0 {
			fresh(args)
		} else if len(args) > 1 {
			existingEC(args)
		} else {
		}
	}
	Help.Add("Create new identity and subchain", cmd)
	return cmd
}()

func existingEC(args []string) {
	if len(args[1]) != 64 {
		fmt.Println("Invalid EC Address, exiting program...")
		return
	}
	// Generate all new Keys from EC
	sid := generateKeysFromEC(args[1])
	if sid == nil {
		return
	}
	start(sid)
}

func fresh(args []string) {
	// Generate all new Keys
	sid := generateKeys()
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

	PrintHeader("Root Chain Curls")
	createIdentityChain(sid)
	registerIdentityChain(sid)
	PrintHeader("Sub Chain Curls")
	createSubChain(sid)
	registerSubChain(sid)
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
func createIdentityChain(sid *functions.ServerIdentity) {
	strCom, strRev, err := functions.CreateIdentityChain(sid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Root Chain ID: " + sid.RootChainID + "\n")

	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

// Step 2 : Root Chain Register
func registerIdentityChain(sid *functions.ServerIdentity) {
	strCom, strRev, err := functions.RegisterServerIdentity(sid)
	if err != nil {
		panic(err)
	}
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

// Step 1 : Subchain Create
func createSubChain(sid *functions.ServerIdentity) {
	strCom, strRev, err := functions.CreateSubChain(sid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sub Chain ID: " + sid.SubChainID + "\n")
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

// Step 2 : Subchain Register
func registerSubChain(sid *functions.ServerIdentity) {
	strCom, strRev, err := functions.RegisterSubchain(sid)
	if err != nil {
		panic(err)
	}
	fmt.Println(strCom + "\n")
	fmt.Println(strRev + "\n")
}

func generateKeys() *functions.ServerIdentity {
	fmt.Print("All new keys for an identity will be created as well as a new Entry \nCredit address. Are you sure you want to continue? y/n :  ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// If not yes, exit
	if strings.Compare("y\n", input) != 0 {
		return nil
	}

	// Key generation
	sid, err := functions.MakeServerIdentity()

	if err != nil {
		panic(err)
	}

	fmt.Println("Key Generation Complete")
	str := formatIDKeysString(sid.IDSet) + "\n"
	str = str + formatECKeyString(sid.ECAddr)

	fmt.Println(str)
	return sid
}

func generateKeysFromEC(ecStr string) *functions.ServerIdentity {
	fmt.Print("All new keys for an identity will be created. Your EC address will \nbe used. Are you sure you want to continue? y/n :  ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// If not yes, exit
	if strings.Compare("y\n", input) != 0 {
		return nil
	}

	// Key generation
	//sec, _ := hex.DecodeString("9FAA5D459E16C50F192630487B52D78EAB2442B29E23BAD433C83986DBC5DA29")
	// TODO Validate input
	sec, err := hex.DecodeString(ecStr)
	if err != nil {
		panic(err)
	}

	sid, err := functions.MakeServerIdentityFromEC(sec)
	if err != nil {
		panic(err)
	}

	fmt.Println("Key Generation Complete")
	str := formatIDKeysString(sid.IDSet) + "\n"
	str = str + formatECKeyString(sid.ECAddr)

	fmt.Println(str)
	return sid
}

func formatIDKeysString(i *identity.IdentitySet) string {
	var str string
	str = "\nPrivate keys and their corresponding levels. Copy these down, this program will \nnot save them for you.\n"
	str = str + "Level 1: " + i.IdentityLevel[0].HumanReadablePrivate() + "\n"
	str = str + "Level 2: " + i.IdentityLevel[1].HumanReadablePrivate() + "\n"
	str = str + "Level 3: " + i.IdentityLevel[2].HumanReadablePrivate() + "\n"
	str = str + "Level 4: " + i.IdentityLevel[3].HumanReadablePrivate() + "\n"

	str = str + "\nPublic keys (hex) and their corresponding levels. Copy these down, this program will \nnot save them for you.\n"
	str = str + "Level 1: " + i.IdentityLevel[0].HumanReadablePublic() + "\n"
	str = str + "Level 2: " + i.IdentityLevel[1].HumanReadablePublic() + "\n"
	str = str + "Level 3: " + i.IdentityLevel[2].HumanReadablePublic() + "\n"
	str = str + "Level 4: " + i.IdentityLevel[3].HumanReadablePublic() + "\n"

	str = str + "\nIdentity keys and their corresponding levels. Copy these down, this program will \nnot save them for you.\n"
	str = str + "Level 1: " + i.IdentityLevel[0].HumanReadableIdentity() + "\n"
	str = str + "Level 2: " + i.IdentityLevel[1].HumanReadableIdentity() + "\n"
	str = str + "Level 3: " + i.IdentityLevel[2].HumanReadableIdentity() + "\n"
	str = str + "Level 4: " + i.IdentityLevel[3].HumanReadableIdentity() + "\n"
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
