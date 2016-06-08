package handlers

import (
	"flag"
	"fmt"
	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/cli"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/serveridentity/identity"
	"os"
	"strconv"
	"strings"
)

/********************************
 *          Cli Control         *
 ********************************/
var Get = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity get 'pubkey'|'idkey' KEY"
	cmd.description = "Get a public key or identity key from a private key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
		c.HandleFunc("pubkey", getPubKey)
		c.HandleFunc("idkey", getIDKey)
		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Get a public key or identity key from a private key", cmd)
	return cmd
}()

func getPubKey(args []string) {
	PrintBanner()
	if len(args) > 1 && len(args[1]) == 53 && strings.Compare(args[1][:2], "sk") == 0 {
		pub, _ := getPrivateKey(args[1])
		fmt.Printf("Public Key: %x\n", pub[:])
	} else {
		fmt.Println("No Private Key given")
	}
}

func getIDKey(args []string) {
	PrintBanner()
	if len(args) > 1 && len(args[1]) == 53 && strings.Compare(args[1][:2], "sk") == 0 {
		_, priv := getPrivateKey(args[1])
		i := identity.NewIdentity()
		lev, err := strconv.Atoi(args[1][2:3])
		if err != nil {
			fmt.Println("Error in input: " + err.Error())
			return
		}
		i.GenerateIdentityFromPrivateKey(priv, lev-1)
		fmt.Println("Identity Key: " + i.HumanReadableIdentity())
	} else {
		fmt.Println("No Private Key given")
	}
}

func getPrivateKey(input string) (*[32]byte, *[64]byte) {
	levInt := input[2:3]
	p := base58.Decode(input[:53])
	if !identity.CheckHumanReadable(p[:]) {
		fmt.Println("Not a valid private key, end hash is incorrect.")
		return nil, nil
	}
	pShort := p[3:35]
	fmt.Println("Key Level: " + levInt)

	var priv [64]byte
	copy(priv[:32], pShort[:32])
	pub := ed.GetPublicKey(&priv)

	return pub, &priv
	//oByte, err := intToOneByte(levInt)
	//if err != nil {
	//	fmt.Println("Error in input: " + err.Error())
	//}
	//i.Input.value = append(pShort[:], oByte[:]...)
}
