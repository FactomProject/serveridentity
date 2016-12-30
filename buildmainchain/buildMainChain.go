package main

import (
	"flag"
	"fmt"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/serveridentity/identity"
)

// Builds the main identity chain. This is helpful for running tests on databases
// without a main chain.

func main() {
	var (
		host = flag.String("host", "localhost:8088", "Change host of factomd instance")
		// EC2DKSYyRcNWf7RS963VFYgMExoHRYLHVeCfQ9PGPmNzwrcmgm2r
		ecSec = flag.String("ec", "Es2Rf7iM6PdsqfYCo3D1tnAR65SkLENyWJG1deUzpRMQmbh9F3eG", "Change entry credit private key to use")
		force = flag.Bool("f", false, "Force creation, even if it exists")
	)

	flag.Parse()
	factom.SetFactomdServer(*host)

	x, _ := factom.GetChainHead(identity.RootRegisterChain)
	if len(x) > 0 {
		if *force {
			fmt.Println("Chain already exists, but creating anyway as it has been forced")
		} else {
			fmt.Println("Chain already exists, exiting...")
			return
		}
	}

	ec, _ := factom.GetECAddress(*ecSec)
	e := new(factom.Entry)
	e.ExtIDs = make([][]byte, 2)
	e.ExtIDs[0] = []byte("Factom Identity Registration Chain")
	e.ExtIDs[1] = []byte("44079090249")
	c := factom.NewChain(e)

	ctx, err := factom.CommitChain(c, ec)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Commit chain:", ctx)

	rtx, err := factom.RevealChain(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Reaveal chain:", rtx)

}
