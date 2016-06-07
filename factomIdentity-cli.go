package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/serveridentity/handlers"
	"github.com/FactomProject/serveridentity/identity"
	"os"
	"strconv"
	"strings"
)

// Testing CLI
func main() {
	fmt.Println("***********************************************************************")
	fmt.Println("**              Factom Server Identity Management Tool              **")
	fmt.Println("***********************************************************************")
	flag.Parse()
	args := flag.Args()

	c := cli.New()
	c.Handle("help", handlers.Help)
	c.Handle("start", handlers.Start)
	c.Handle("newkey", handlers.NewKey)
	c.HandleFunc("generatenewkeys", generateNewKeys)
	c.HandleFunc("generatefromkeys", generateFromKeys)

	c.HandleDefault(handlers.Help)
	c.Execute(args)
}

func generateFromKeys(args []string) {
	idSet := identity.NewIdentitySet()
	// TODO: Check Args for human readable
	var keys [4]*[64]byte
	for i := range keys {
		keys[i] = new([64]byte)
	}

	// TODO: Error Handling
	key1, err := hex.DecodeString(args[1])
	if err != nil {
		panic(err)
	}
	key2, err := hex.DecodeString(args[2])
	if err != nil {
		panic(err)
	}
	key3, err := hex.DecodeString(args[3])
	if err != nil {
		panic(err)
	}
	key4, err := hex.DecodeString(args[4])
	if err != nil {
		panic(err)
	}

	copy(keys[0][:32], key1[:])
	copy(keys[1][:32], key2[:])
	copy(keys[2][:32], key3[:])
	copy(keys[3][:32], key4[:])

	idSet.GenerateIdentitySetFromPrivateKeys(keys)
	idChain := identity.MakeRootChainFromIdentitySet(idSet)

	_ = idChain
}

func generateNewKeys(args []string) {
	fmt.Println("The following keys will take up to a minute to generate, do you wish to proceed?\n Type 'y'/'n' followed by the enter key")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if strings.Compare(input, "y\n") == 0 {
		fmt.Println("***********************************************************************")
		fmt.Println("**                          Keys and their Levels                    **")
		fmt.Println("***********************************************************************")
		idChain := identity.MakeRootChain()
		for i := 0; i < 4; i++ {
			//fmt.Println("************************************************************")
			if i != 0 {
				fmt.Println()
			}
			fmt.Println("                           Level " + strconv.Itoa(i))
			//fmt.Println("************************************************************")
			fmt.Println("  Private Key : " + idChain.IdSet.IdentityLevel[i].HumanReadablePrivate())
			fmt.Println("  Identity Key: " + idChain.IdSet.IdentityLevel[i].HumanReadableIdentity())
		}
	}

}
