package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(_aes string, src []byte, iv string) []byte {
	key := []byte(_aes)
	block, err := aes.NewCipher(key)
	if nil != err {
		panic(err)
	}
	//填充
	src = paddingText(src, block.BlockSize())
	//初始化向量
	iv0 := []byte(iv)
	//创建加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv0)
	dst := make([]byte, len(src))
	//加密
	blockMode.CryptBlocks(dst, src)
	return dst
}

func Decrypt(_aes string, src []byte, iv string) []byte {
	key := []byte(_aes)
	block, err := aes.NewCipher(key)
	if nil != err {
		panic(err)
	}
	//初始化向量
	iv0 := []byte(iv)
	blockMode := cipher.NewCBCDecrypter(block, iv0)
	//解密
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	//去除填充
	dst = unPaddingText(dst)
	return dst
}

//func main() {
//	src := []byte("你好！潘丽萍")
//	key := []byte("8765432112345678")
//	encryptedMsg := EncryptAES(src, key)
//	fmt.Println("加密后:", hex.EncodeToString(encryptedMsg))
//	decryptedMsg := DecryptAES(encryptedMsg, key)
//	fmt.Println("解密后的明文:", string(decryptedMsg))
//}

// paddingText 给最后一组数据填充至64字节
func paddingText(src []byte, blockSize int) []byte {
	//求出最后一个分组需要填充的字节数
	padding := blockSize - len(src)%blockSize
	//创建新的切片，切片字节数为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//将新创建的切片和带填充的数据进行拼接
	nextText := append(src, padText...)
	return nextText
}

// unPaddingText 取出数据尾部填充的赘余字符
func unPaddingText(src []byte) []byte {
	//获取待处理数据长度
	len := len(src)
	//取出最后一个字符
	num := int(src[len-1])
	newText := src[:len-num]
	return newText
}
