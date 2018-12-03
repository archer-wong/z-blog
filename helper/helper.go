package helper

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"
	"z-blog/modules/setting"
)

func Encrypt(raw string) string {
	salt := setting.Cfg.Section("app").Key("SECRET").String()

	sha1 := sha1.New()
	sha1.Write([]byte(raw + salt))
	encrypt := hex.EncodeToString(sha1.Sum([]byte("")))

	return encrypt
}

func DateTime(t time.Time, layout string) string {
	return t.Format(layout) //`2006-01-02 15:04:05`
}

func SubString(str string, num int) string {
	subStr := []rune(str)
	len := len(subStr)
	if len > num {

		return string(subStr[0:num])
	}

	return str
}

var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen = len(codes)
)
func RandStringBytes(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}

