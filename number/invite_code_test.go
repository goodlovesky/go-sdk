package number

import (
	"testing"
)

func TestEncode(t *testing.T) {
	for i := 1; i < 100; i++ {
		t.Log(GenerateInviteCode(uint64(i)))
	}
}

func TestDecode(t *testing.T) {
	for i := 1000000000; i < 1000000100; i++ {
		result := GenerateInviteCode(uint64(i))
		t.Log(i, result)
		num := DecodeInviteCode(result)
		if num != uint64(i) {
			t.Fatal("num != i")
		}
	}
}
