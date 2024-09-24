package string

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@!#$%&*"
	bytes := []byte(str)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		var result []byte
		for i := 0; i < length; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}

		if strings.ContainsAny(string(result), "0123456789") && strings.ContainsAny(string(result), "@!#$%&*") &&
			strings.ContainsAny(string(result), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			time.Sleep(time.Nanosecond)
			return string(result)
		}
	}
}
