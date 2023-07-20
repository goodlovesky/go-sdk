package strs

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(Hex(16))
}

func BenchmarkTestString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Hex(16)
	}
}

func TestString2(t *testing.T) {
	fmt.Println(RandString(16))
}

func BenchmarkTestString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(16)
	}
}
