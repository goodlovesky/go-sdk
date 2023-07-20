package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V 散列算法
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
