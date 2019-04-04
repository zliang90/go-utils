package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Sum(text []byte) string {
	m := md5.New()

	_, err := m.Write(text)
	if err != nil {
		return ""
	}

	ch := m.Sum(nil)
	return hex.EncodeToString(ch)
}
