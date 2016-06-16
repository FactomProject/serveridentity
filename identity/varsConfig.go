package identity

import ()

var (
	// Level to use in signing when registering
	RegisterIdentityLevel int = 1
	// For ChainIDs
	ProofOfWorkChainID     = []byte{0x88, 0x88, 0x88}
	ProofOfWorkLength  int = 3 // DEFAULT = 3, set lower for faster tests
	// For Registering Root Chain
	RootRegisterChain string = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	// If true will show "." on screen to show user program is computing
	ShowBruteForce bool = true
	// Will change frequency of dots for brute forcing
	BruteForcePeriod int = 400000
	// Amount of hashes on seed to generate MHash
	MHashAmount int = 100000
	//Seed Hex Length (length/2 = bytes)
	SeedMin int = 4
	SeedMax int = 64
)
