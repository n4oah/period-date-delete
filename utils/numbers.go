package utils

import (
	"strconv"
)

func IsStringToInt(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}
