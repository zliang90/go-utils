package user

import (
	"errors"
	"fmt"
	"github.com/zliang90/go-utils/file"
	"os/user"
)

func GetHomeDir(userName string) (string, error) {
	// get user
	var (
		u   *user.User
		err error
	)

	if userName != "" {
		u, err = user.Lookup(userName)
	} else {	// current user
		u, err = user.Current()
	}

	if err != nil {
		return "", err
	}

	if !file.IsExists(u.HomeDir) {
		return "", errors.New(fmt.Sprintf("[%s] homeDir does not exists", u.HomeDir))
	}

	return u.HomeDir, nil
}
