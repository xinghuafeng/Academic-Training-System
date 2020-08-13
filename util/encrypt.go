package util

import (
	"crypto/md5"
	"encoding/hex"
)

// md5
func Md5Encrypt(in string) string {
	h := md5.New()
	h.Write([]byte(in))
	out := hex.EncodeToString(h.Sum(nil))
	return out
}
