package hash

import (
	"bytes"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

var NIL_HASH = NewHashFromString("0xa7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a")

type Hash struct {
	content []byte
}

func (hash *Hash) Bytes() []byte {
	return hash.content
}

func (hash *Hash) PrePend(prefix []byte) []byte {
	return append(prefix, hash.content...)
}

func (hash *Hash) Append(suffix []byte) []byte {
	return append(hash.content, suffix...)
}

func (hash *Hash) Equal(target *Hash) bool {
	if target == nil {
		return false
	}
	return bytes.Equal(hash.content, target.content)
}

func (hash *Hash) Clone() *Hash {
	return NewHashFromBytes(hash.content)
}

func IsNilHash(target *Hash) bool {
	if target == nil {
		return true
	}
	return bytes.Equal(target.content, NIL_HASH.content)
}

// sha3-256 hash
func CalHash(input []byte) *Hash {
	sha3_g := sha3.New256()
	// Create a new hash & write input string
	sha3_g.Write([]byte(input))
	// Get the resulting encoded byte slice
	return &Hash{
		content: sha3_g.Sum(nil),
	}
}

// case 1 , input is nil , return const of empty hash
// case 2 , input is longer then 32 , return the first 32 bytes of hash
func NewHashFromBytes(input []byte) *Hash {
	result := Hash{
		content: make([]byte, 32),
	}
	if len(input) > 32 {
		copy(result.content[0:32], input[len(input)-32:])
	} else {
		copy(result.content[32-len(input):], input)
	}
	return &result
}

func NewHashFromString(input string) *Hash {
	input_ := input
	if len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X') {
		input_ = input[2:]
	}
	if len(input)%2 == 1 {
		input_ = "0" + input_
	}

	h, _ := hex.DecodeString(input_)
	return NewHashFromBytes(h)
}
