package identity

import ()

type MHash struct {
	version     []byte
	message     []byte
	rootChainID []byte
	newMHash    []byte
	timestamp   []byte
	preimage    []byte
	signiture   []byte
}
