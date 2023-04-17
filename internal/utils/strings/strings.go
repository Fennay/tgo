package strings

import (
	"math/rand"
	"time"
)

const (
	Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersAndNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func init() {
	// 添加随机种子
	rand.Seed(time.Now().UnixNano())
}

// RandStringRunes 随机字符串
func RandStringRunes(n int) string {
	lettersRune := []rune(Letters)
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRune[rand.Intn(len(lettersRune))]
	}
	return string(b)
}
