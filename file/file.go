package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	DefaultBufferSize = 4096
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ToString(f string) (string, error) {
	b, err := ioutil.ReadFile(f)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ToLineSlice(path string) ([]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	lineSlice := make([]string, 0)
	buf := bufio.NewReaderSize(f, DefaultBufferSize)

	for {
		line, _, err := buf.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lineSlice = append(lineSlice, string(line))
	}
	return lineSlice, nil
}

func ToTrimString(path string) (string, error) {
	s, err := ToString(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(s), nil
}

func Contains(path, sub string) (bool, error) {
	lineSlice, err := ToLineSlice(path)
	if err != nil {
		return false, err
	}

	exits := false

	for _, line := range lineSlice {
		if strings.Contains(line, sub) {
			exits = true
			break
		}
	}

	return exits, nil
}
