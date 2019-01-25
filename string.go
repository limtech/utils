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

// func RandomString(n int) string {
// 	letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 	randStr := make([]rune, n)
// 	for i := range randStr {
// 		randStr[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(randStr)
// }

func RandomString(n int) string {
	const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// // markdown to html
// func Markdown2Html(markdown string) string {
// 	return string(blackfriday.MarkdownBasic([]byte(markdown)))
// }

// // html to markdown
// func Html2Markdown(html string) string {
// 	return string(blackfriday.MarkdownBasic([]byte(html)))
// }

// // make v4 uuid
// func Uuid() string {
// 	return uuid.Must(uuid.NewV4()).String()
// }

// // parse uuid
// func UuidParse(str string) uuid.UUID {
// 	return uuid.FromStringOrNil(str)
// }
