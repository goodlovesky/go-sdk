package aes

import (
	"encoding/hex"
	"fmt"
	"github.com/goodlovesky/sdk/strs"
	"math/rand"
	"testing"
	"time"
)

func TestAES(t *testing.T) {

	aes := "c&1*~2^#0^=)^7%b"
	msg := []byte("你好！我是准备测试的数据,aaaaa")
	key := strs.RandString(16)
	encryptedMsg := Encrypt(aes, msg, key)
	fmt.Println(len(encryptedMsg))
	decryptMsg := Decrypt(aes, encryptedMsg, key)
	fmt.Printf("%v\n%v", hex.EncodeToString(encryptedMsg), string(decryptMsg))
}

func BenchmarkAES(b *testing.B) {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_ = strs.RandString(16)
	}
}
