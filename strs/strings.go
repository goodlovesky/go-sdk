package strs

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

var randSource = rand.NewSource(time.Now().UnixNano())

const (
	Base64Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	HexChars    = "0123456789abcdef"
	DecChars    = "0123456789"

	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// IsImageExt 是否图片格式
func IsImageExt(s string) bool {
	var ret = false
	var extList = []string{"png", "jpg", "jpeg", "bmp", "webp", "gif", "tif"}
	var s1 = strings.Split(s, ".")
	var imageType = s1[len(s1)-1]
	for _, value := range extList {
		value = strings.ToLower(value)
		if value == imageType {
			ret = true
			break
		}
	}
	return ret
}

// IsVideoExt 是否视频格式
func IsVideoExt(s string) bool {
	var ret = false
	var extList = []string{"mp4", "flv", "mov", "avi", "wmv", "mpg", "ts"}
	var s1 = strings.Split(s, ".")
	var imageType = s1[len(s1)-1]
	for _, value := range extList {
		value = strings.ToLower(value)
		if value == imageType {
			ret = true
			break
		}
	}
	return ret
}

// IsAudioExt 是否音频格式
func IsAudioExt(s string) bool {
	var ret = false
	var extList = []string{"mp3", "wav", "aac", "wma", "ogg", "m4a", "amr", "audio", "m3u8"}
	var s1 = strings.Split(s, ".")
	var imageType = s1[len(s1)-1]
	for _, value := range extList {
		value = strings.ToLower(value)
		if value == imageType {
			ret = true
			break
		}
	}
	return ret
}

// RandString 随机位数的字符串
func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, randSource.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Base64 generates a random Base64 string with length of n
//
// Example: X02+jDDF/exDoqPg9/aXlzbUCN93GIQ5
func Base64(n int) string { return RandString2(n, Base64Chars) }

// Base62 generates a random Base62 string with length of n
//
// Example: 1BsNqB61o4ztSqLC6labKGNf4MYy352X
func Base62(n int) string { return RandString2(n, Base62Chars) }

// Dec generates a random decimal number string with length of n
//
// Example: 37110235710860781655802098192113
func Dec(n int) string { return RandString2(n, DecChars) }

// Hex generates a random Hexadecimal string with length of n
//
// Example: 67aab2d956bd7cc621af22cfb169cba8
func Hex(n int) string { return RandString2(n, HexChars) }

// list of default letters that can be used to make a random string when calling String
// function with no letters provided
var defLetters = []rune(Base62Chars)

// RandString2 随机生成
func RandString2(n int, letters ...string) string {
	var letterRunes []rune
	if len(letters) == 0 {
		letterRunes = defLetters
	} else {
		letterRunes = []rune(letters[0])
	}
	var bb bytes.Buffer
	bb.Grow(n)
	l := uint32(len(letterRunes))
	// on each loop, generate one random rune and append to output
	for i := 0; i < n; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(RandomBytes(4))%l])
	}
	//return bb.String()
	return *(*string)(unsafe.Pointer(&bb))
}

// RandomBytes generates n random bytes
func RandomBytes(n int) []byte {
	b := make([]byte, n)
	//randSource := rand.NewSource(time.Now().UnixNano())
	//randObj := rand.New(randSource)
	//_, err := randObj.Read(b)
	//_, err := rand.Read(b)
	_, err := rand.New(randSource).Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
