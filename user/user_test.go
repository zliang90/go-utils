package user

import (
	"fmt"
	"testing"
)

func TestGetHomeDir(t *testing.T) {
	path, err := GetHomeDir("")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("home path: ", path)
}
