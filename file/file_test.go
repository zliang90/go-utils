package file

import (
	"testing"
)

func TestIsExists(t *testing.T) {
	b := IsExists("test_file.txt")

	if !b {
		t.Error()
	}
}

func TestToString(t *testing.T) {
	s, err := ToString("test_file.txt")

	if err != nil {
		t.Error(err.Error())
	}

	if len(s) == 0 {
		t.Error("ToString() return 0 bytes")
	}
}

func TestToLineSlice(t *testing.T) {
	sl, err := ToLineSlice("test_file.txt")

	if err != nil {
		t.Error(err.Error())
	}

	if len(sl) != 3 {
		t.Error()
	}
}

func TestContains(t *testing.T) {
	b, err := Contains("test_file.txt", "test file")

	if err != nil {
		t.Error(err.Error())
	}

	if !b {
		t.Error()
	}
}
