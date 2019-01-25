package gotuil

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// EncodeMD5 Md5加密(32位大写)
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	md5Str := hex.EncodeToString(m.Sum(nil))
	return strings.ToUpper(md5Str)
}
