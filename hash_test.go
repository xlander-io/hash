package hash

import (
	"fmt"
	"testing"
)

func Test_Hash(t *testing.T) {

	h := NewHashFromString("0xa7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a")
	fmt.Println(IsNilHash(nil))
	fmt.Println(IsNilHash(h))

	fmt.Println(h.Equal(NewHashFromString("0xa7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a")))

}
