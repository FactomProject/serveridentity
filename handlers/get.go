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
	cmd.helpMsg = "serveridentity get pubkey|idkey|btc|privkey KEY"
	cmd.description = "Get a public key or identity key from a private key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()
		c.Handle("pubkey", gpubKey)
		c.Handle("idkey", gidKey)
		c.Handle("btckey", gbtcKey)
		c.Handle("privkey", gprivKey)
		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Get a public key or identity key from a private key", cmd)
	return cmd
}()

var gidKey = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity get idkey KEY"
	cmd.description = "Get a identity key from a private key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No key given, 'serveridentity get idkey KEY'")
			return
		}
		getIDKey(args[1])

	}
	Help.Add("Get a identity key from a private key", cmd)
	return cmd
}()

var gpubKey = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity get pubkey KEY"
	cmd.description = "Get a public key from a private key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No key given, 'serveridentity get pubkey KEY'")
			return
		}
		getPubKey(args[1])

	}
	Help.Add("Get a public key from a private key", cmd)
	return cmd
}()

var gprivKey = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity get priv KEY"
	cmd.description = "Get the private key hex from a SK# key"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No key given, 'serveridentity get privkey KEY'")
			return
		}
		getPrivateKeyHex(args[1])

	}
	Help.Add("Get the private key hex from a SK# key", cmd)
	return cmd
}()

var gbtcKey = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity get btckey KEY"
	cmd.description = "Get the 20byte bitcoin address from a ~34character btc pubkey"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No key given, 'serveridentity get btckey KEY'")
			return
		}
		getBtcKey(args[1])

	}
	Help.Add("Get hash160", cmd)
	return cmd
}()

/********************************
 *        CLI Functions         *
 ********************************/

func getBtcKey(key string) {
	PrintBanner()
	if len(key) > 34 || len(key) < 25 {
		fmt.Println("Error: Invalid btc key length")
	} else if strings.Compare(key[:1], "1") == 0 {
		x := base58.Decode(key)
		fmt.Println("Btc Type: P2PKH")
		fmt.Printf("Btc Hash160: %x\n", x[1:21])
	} else if strings.Compare(key[:1], "3") == 0 {
		x := base58.Decode(key)
		fmt.Println("Btc Type: P2SH")
		fmt.Printf("Btc Hash160: %x\n", x[1:21])
	} else {
		fmt.Println("Error: Invalid btc key prefix")
	}
}

func getPubKey(key string) {
	PrintBanner()
	if len(key) == 53 && strings.Compare(key[:2], "sk") == 0 {
		if lev, err := strconv.Atoi(key[2:3]); err != nil {
			fmt.Println("Error in input: " + err.Error())
			return
		} else if lev < 1 || lev > 4 {
			fmt.Println("Error: Key level is outside range (1-4)")
			return
		}
		fmt.Println(key)
		pub, _ := getPrivateKey(key)
		fmt.Printf("Public Key: %x\n", pub[:])
	} else {
		if len(key) != 53 {
			fmt.Println("Error: Invalid private key length")
		} else if strings.Compare(key[:2], "sk") != 0 {
			fmt.Println("Error: Invalid private key prefix")
		} else {
			fmt.Println("Error: Invalid private key")
		}
	}
}

func getIDKey(key string) {
	PrintBanner()
	if len(key) == 53 && strings.Compare(key[:2], "sk") == 0 {
		var lev int
		if lev, err := strconv.Atoi(key[2:3]); err != nil {
			fmt.Println("Error in input: " + err.Error())
			return
		} else if lev < 1 || lev > 4 {
			fmt.Println("Error: Key level is outside range (1-4)")
			return
		}

		_, priv := getPrivateKey(key)
		i := identity.NewIdentity()
		i.GenerateIdentityFromPrivateKey(priv, lev)
		fmt.Println(i.HumanReadableIdentity())
	} else {
		if len(key) != 53 {
			fmt.Println("Error: Invalid private key length")
		} else if strings.Compare(key[:2], "sk") != 0 {
			fmt.Println("Error: Invalid private key prefix")
		} else {
			fmt.Println("Error: Invalid private key")
		}
	}
}

func getPrivateKeyHex(key string) {
	PrintBanner()
	if len(key) == 53 && strings.Compare(key[:2], "sk") == 0 {
		if lev, err := strconv.Atoi(key[2:3]); err != nil {
			fmt.Println("Error in input: " + err.Error())
			return
		} else if lev < 1 || lev > 4 {
			fmt.Println("Error: Key level is outside range (1-4)")
			return
		}

		_, priv := getPrivateKey(key)
		fmt.Printf("Private Key Hex: %x\n", priv[:32])
	} else {
		if len(key) != 53 {
			fmt.Println("Error: Invalid private key length")
		} else if strings.Compare(key[:2], "sk") != 0 {
			fmt.Println("Error: Invalid private key prefix")
		} else {
			fmt.Println("Error: Invalid private key")
		}
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
