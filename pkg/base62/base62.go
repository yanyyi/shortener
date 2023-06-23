package base62

import (
	"math"
	"strings"
)

var (
	base62Str  string
	baseStrLen uint64
)

// MustInit 要使用base62这个包必须要调用该函数完成初始化
func MustInit(baseString string) {
	if len(baseString) == 0 {
		panic("need base string!")
	}
	base62Str = baseString
	baseStrLen = uint64(len(baseString))
}

// 为了避免被人恶意请求，我们可以将上面的字符串打乱

// Int2String 十进制数转62进制字符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	var bl []byte
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, base62Str[mod])
		seq = div
	}
	return string(reverse(bl))
}

// Stirng2Int 62进制数转十进制字符串
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	// 从右往左遍历
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += uint64(strings.Index(base62Str, string(b))) * uint64(base)
	}
	return seq
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
