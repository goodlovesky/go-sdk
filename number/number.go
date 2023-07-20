package number

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenValidateCode 生产随机多位数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[r1.Intn(r)])
	}
	return sb.String()
}
