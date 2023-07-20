package sm4

import (
	"bytes"
	"crypto/cipher"
	"github.com/tjfoc/gmsm/sm4"
)

// Encrypt 传入的key长度必须和配置的sm4key长度相等
func Encrypt(_sm4 string, src []byte, key string) []byte {
	//创建加密块
	block, err := sm4.NewCipher([]byte(_sm4))
	if nil != err {
		panic(err)
	}
	//填充数据
	src = paddingText(src, block.BlockSize())
	//初始化向量
	iv := []byte(key)
	//设置加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// Decrypt 传入的key长度必须和配置的sm4key长度相等
func Decrypt(_sm4 string, src []byte, key string) []byte {
	//创建解密块
	block, err := sm4.NewCipher([]byte(_sm4))
	if nil != err {
		panic(err)
	}
	//初始化向量
	iv := []byte(key)
	//创建解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(src))
	//解密
	blockMode.CryptBlocks(dst, src)
	//去除填充
	dst = unPaddingText(dst)
	return dst
}

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
