package array

import "strings"

func InArray(v string, arr []string) bool {
	isExist := false

	for _, s := range arr {
		if ok := strings.Contains(s, v); ok {
			isExist = true
			break
		}
	}
	return isExist
}
