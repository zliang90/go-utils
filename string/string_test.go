package string

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := "i'm gopher!"

	r, _ := Reverse(s)

	if r != "!rehpog m'i" {
		t.Error("string reverse not eq '!rehpog m'i' ")
	}

	s = "中国人！"

	r, _ = Reverse(s)
	if r != "！人国中" {
		t.Error("string reverse not eq '！人国中'")
	}
}

func TestCount(t *testing.T) {
	s := "hello, i'm gopher!"

	if len(s) != len(s) {
		t.Error("length not equal")
	}
}

func TestToSlice(t *testing.T) {
	s := "我是中国人"
	ss := ToSlice(s)

	if len(ss) == 0 {
		t.Error("ToSlice() returns 0 elements in array")
	}

	fmt.Println(ss)
}
