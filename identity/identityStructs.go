package identity

import (
	"bytes"
	"crypto/rand"
	//"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/FactomProject/btcutil/base58"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/serveridentity/utils"
)

// Prefixes for Human Readable
var (
	privPrefix1 = [3]byte{0x4d, 0xb6, 0xc9}
	privPrefix2 = [3]byte{0x4d, 0xb6, 0xe7}
	privPrefix3 = [3]byte{0x4d, 0xb7, 0x05}
	privPrefix4 = [3]byte{0x4d, 0xb7, 0x23}
	pubPrefix1  = [3]byte{0x3f, 0xbe, 0xba}
	pubPrefix2  = [3]byte{0x3f, 0xbe, 0xd8}
	pubPrefix3  = [3]byte{0x3f, 0xbe, 0xf6}
	pubPrefix4  = [3]byte{0x3f, 0xbf, 0x14}
)

// Individual component in an IdentitySet
type Identity struct {
	privateKey  *[ed.PrivateKeySize]byte
	publicKey   *[ed.PublicKeySize]byte
	identityKey *[ed.PublicKeySize]byte
	privPrefix  []byte
	pubPrefix   []byte
}

func (i *Identity) GetPrefix() []byte {
	return i.pubPrefix
}

func (i *Identity) GetPrivateKey() *[ed.PrivateKeySize]byte {
	return i.privateKey
}
func (i *Identity) GetPublicKey() *[ed.PublicKeySize]byte {
	return i.publicKey
}
func (i *Identity) GetIdentityKey() *[ed.PublicKeySize]byte {
	return i.identityKey
}

func NewIdentity() *Identity {
	id := new(Identity)
	//id.privPrefix = new([3]byte)
	return id
}

func (i *Identity) GenerateIdentity(count int) error {
	pub, priv, err := ed.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	i.privateKey = priv
	i.publicKey = pub
	i.generateIdentityKey(count)
	return nil
}

func (i *Identity) GenerateIdentityFromPrivateKey(key *[ed.PrivateKeySize]byte, count int) error {
	i.publicKey = ed.GetPublicKey(key)
	i.privateKey = key
	i.generateIdentityKey(count)
	return nil
}

func (i *Identity) generateIdentityKey(count int) error {
	var temp [33]byte
	if i.privateKey == nil || i.publicKey == nil {
		return errors.New("Public/Private key pair not yet set")
	}

	copy(temp[:1], []byte{0x01})
	copy(temp[1:], i.publicKey[:])

	//shaReturn := sha256.Sum256(temp[:])
	//shaReturn = sha256.Sum256(shaReturn[:])
	shaReturn := utils.Sha256d(temp[:])
	var t [32]byte
	copy(t[:32], shaReturn[:32])
	i.identityKey = &t

	switch count {
	case 0:
		i.privPrefix = append(i.privPrefix[:], privPrefix1[:]...)
		i.pubPrefix = append(i.pubPrefix[:], pubPrefix1[:]...)
	case 1:
		i.privPrefix = append(i.privPrefix[:], privPrefix2[:]...)
		i.pubPrefix = append(i.pubPrefix[:], pubPrefix2[:]...)
	case 2:
		i.privPrefix = append(i.privPrefix[:], privPrefix3[:]...)
		i.pubPrefix = append(i.pubPrefix[:], pubPrefix3[:]...)
	case 3:
		i.privPrefix = append(i.privPrefix[:], privPrefix4[:]...)
		i.pubPrefix = append(i.pubPrefix[:], pubPrefix4[:]...)
	}
	return nil
}

func (i *Identity) HumanReadablePrivate() string {
	buf := new(bytes.Buffer)
	// Add Prefix
	buf.Write(i.privPrefix[:])

	// Add Key
	buf.Write(i.privateKey[:32])

	o := buf.Bytes()
	// Sha356d
	humanReadable := utils.Sha256d(o)

	// Append first 4 bytes
	o = append(o, humanReadable[:4]...)

	str := base58.Encode(o)
	return str
}

func CheckHumanReadable(key []byte) bool {
	// Sha356d
	humanReadable := utils.Sha256d(key[:35])

	// Append first 4 bytes
	if bytes.Compare(humanReadable[:4], key[len(key)-4:]) == 0 {
		return true
	}
	return false
}

func (i *Identity) HumanReadablePublic() string {
	buf := new(bytes.Buffer)

	// Add Key
	buf.Write(i.publicKey[:32])

	o := buf.Bytes()

	str := hex.EncodeToString(o)
	return str
}

func (i *Identity) HumanReadableIdentity() string {
	buf := new(bytes.Buffer)
	// Add Prefix
	buf.Write(i.pubPrefix[:])

	// Add Key
	buf.Write(i.identityKey[:32])

	o := buf.Bytes()
	// Sha356d
	humanReadable := utils.Sha256d(o)

	// Append first 4 bytes
	o = append(o, humanReadable[:4]...)

	str := base58.Encode(o)
	return str
}

// Set of 4 Identities
type IdentitySet struct {
	IdentityLevel [4]*Identity
}

func NewIdentitySet() *IdentitySet {
	idSet := new(IdentitySet)
	return idSet
}

func (i *IdentitySet) GenerateIdentitySet() error {
	for count := 0; count < 4; count++ {
		id := NewIdentity()
		id.GenerateIdentity(count)
		i.IdentityLevel[count] = id
	}
	return nil
}

func MergeKeys(key1 [ed.PrivateKeySize]byte, key2 [ed.PrivateKeySize]byte, key3 [ed.PrivateKeySize]byte, key4 [ed.PrivateKeySize]byte) [4]*[ed.PrivateKeySize]byte {
	var keys [4]*[64]byte
	copy(keys[0][:], key1[:])
	copy(keys[1][:], key2[:])
	copy(keys[2][:], key3[:])
	copy(keys[3][:], key4[:])
	return keys
}

func (i *IdentitySet) GenerateIdentitySetFromPrivateKeys(keys [4]*[ed.PrivateKeySize]byte) error {
	for count := 0; count < 4; count++ {
		id := NewIdentity()
		id.GenerateIdentityFromPrivateKey(keys[count], count)
		i.IdentityLevel[count] = id
	}
	return nil
}
