package string

import (
	"bytes"
	"errors"
)

// Reverse 字符串反转
func Reverse(s string) (string, error) {
	rs := []rune(s)
	lenRs := len(rs)

	if lenRs == 0 {
		return "", errors.New("empty string")
	}

	buf := bytes.NewBufferString("")

	for i := lenRs - 1; i >= 0; i-- {
		if _, err := buf.WriteRune(rs[i]); err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}

func ToSlice(s string) []string {

	var st []string

	for _, sv := range s {
		st = append(st, string(sv))
	}

	return st
}
