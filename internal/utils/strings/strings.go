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

// Rand 自定义随机数
func Rand(source string, n int) string {
	return zyRand(source, n)
}

// RandStringWithA2z 随机字符串 [A-z]
func RandStringWithA2z(n int) string {
	return zyRand(Letters, n)
}

// RandNumRunes 随机获取数字
func RandNumRunes(n int) string {
	return zyRand(Num, n)
}

func zyRand(source string, n int) string {
	lettersRune := []rune(source)
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRune[rand.Intn(len(lettersRune))]
	}
	return string(b)
}
