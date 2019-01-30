package utils

import (
	"math/rand"
	"time"
)

// SubString work well with unicode
func SubString(str string, start int, length int) string {
	r := []rune(str)
	n := len(r)
	if start >= n || length <= 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	end := start + length
	if end > n {
		end = n
	}

	return string(r[start:end])
}

func RandomString(n int, format int) string {
	const (
		CHAR_NUMBER        = `0123456789`
		CHAR_LETTERS_LOWER = `abcdefghijklmnopqrstuvwxyz`
		CHAR_LETTERS_UPPER = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
		CHAR_SYMBOLS       = `~!@#$%^&*()_+-={}|[]\:";'<>?,./`
	)

	var chars []rune
	switch format {
	case 0:
		chars = []rune(CHAR_NUMBER)

	case 1:
		chars = []rune(CHAR_NUMBER + CHAR_LETTERS_LOWER)

	case 2:
		chars = []rune(CHAR_NUMBER + CHAR_LETTERS_LOWER + CHAR_LETTERS_UPPER)

	default:
		chars = []rune(CHAR_NUMBER + CHAR_LETTERS_LOWER + CHAR_LETTERS_UPPER + CHAR_SYMBOLS)
	}
	rand.Seed(time.Now().UnixNano())
	randString := make([]rune, n)
	for i := range randString {
		randString[i] = chars[rand.Intn(len(chars))]
	}
	return string(randString)
}
