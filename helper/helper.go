package helper

import (
	"crypto/sha1"
	"encoding/hex"
	"z-blog/modules/setting"
)

func Encrypt(raw string) string {
	salt := setting.Cfg.Section("app").Key("SECRET").String()

	sha1 := sha1.New()
	sha1.Write([]byte(raw + salt))
	encrypt := hex.EncodeToString(sha1.Sum([]byte("")))

	return encrypt
}
