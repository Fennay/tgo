package strings

import (
	"math/rand"
	"time"
)

const (
	LettersUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersLower  = "abcdefghijklmnopqrstuvwxyz"
	Num           = "0123456789"
	Letters       = LettersLower + LettersUpper
	LettersAndNum = LettersLower + LettersUpper + Num
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
