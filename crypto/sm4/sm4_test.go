package sm4

import (
	"encoding/hex"
	"fmt"
	"github.com/goodlovesky/sdk/strs"
	"math/rand"
	"testing"
	"time"
)

func TestSM4(t *testing.T) {

	sm4 := "2yr2exh2372h2837" // 16位
	msg := []byte("你好！我是准备测试的数据,aaaaa")
	key := strs.RandString(16)
	encryptedMsg := Encrypt(sm4, msg, key)
	fmt.Println(len(encryptedMsg))
	decryptMsg := Decrypt(sm4, encryptedMsg, key)
	fmt.Printf("%v\n%v", hex.EncodeToString(encryptedMsg), string(decryptMsg))
}

func BenchmarkSM4(b *testing.B) {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_ = strs.RandString(16)
	}
}
