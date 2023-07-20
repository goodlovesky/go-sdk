package test3

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func TestString3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(randStr(16))
}

func BenchmarkTestString3(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_ = randStr(16)
	}
}
