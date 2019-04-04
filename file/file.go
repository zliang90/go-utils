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

func IsExists(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func ToString(f string) (string, error) {
	b, err := ioutil.ReadFile(f)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ToLineSlice(filePath string) ([]string, error) {
	f, err := os.Open(filePath)

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

func ToTrimString(f string) (string, error) {
	s, err := ToString(f)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(s), nil
}

// Contains 判断文件中是否包含sub字符串
func Contains(filePath, sub string) (bool, error) {
	lineSlice, err := ToLineSlice(filePath)
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
