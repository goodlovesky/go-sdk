package crypto

import (
	"github.com/goodlovesky/sdk/crypto/aes"
	"github.com/goodlovesky/sdk/crypto/sm4"
)

func EncryptSM4(_sm4 string, msg []byte, iv string) []byte {
	return sm4.Encrypt(_sm4, msg, iv)
}

func DecryptSM4(_sm4 string, msg []byte, iv string) []byte {
	return sm4.Decrypt(_sm4, msg, iv)
}

func EncryptAES(_aes string, msg []byte, iv string) []byte {
	return aes.Encrypt(_aes, msg, iv)
}

func DecryptAES(_aes string, msg []byte, iv string) []byte {
	return aes.Decrypt(_aes, msg, iv)
}
