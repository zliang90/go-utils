package crypto

import (
	"testing"
)

func TestMd5Sum(t *testing.T) {
	m := Md5Sum([]byte("gopher"))

	if m == "" {
		t.Error("Md5Sum() return empty string")
	}

	if m != "ca654591d4ac97414391907f882b3c05" {
		t.Error("'gopher' md5 sum not eq 'ca654591d4ac97414391907f882b3c05'")
	}

}
