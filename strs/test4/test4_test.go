package test4

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const (
	// 6 bits to represent a letters index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
)

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		if idx := int(rand.Int63() & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i++
		}
	}
	return string(b)
}

func TestString4(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(randStr(16))
}

func BenchmarkTestString4(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_ = randStr(16)
	}
}
