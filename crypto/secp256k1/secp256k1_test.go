package secp256k1

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := NewPrivateKey(S256())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key)
}
