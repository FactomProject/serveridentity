package identity

import ()

var (
	// Level to use in signing when registering
	RegisterIdentityLevel int = 4
	// For ChainIDs
	ProofOfWorkChainID     = []byte{0x88, 0x88, 0x88}
	ProofOfWorkLength  int = 1 // DEFAULT = 3, set lower for faster tests
	// For Registering Root Chain
	RootRegisterChain string = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	// If true will show "." on screen to show user program is computing
	ShowBruteForce bool = false
)
