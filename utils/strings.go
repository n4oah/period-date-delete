package utils

import (
	"log"
	"strings"
)

func RemoveArrayInSideSpace(paths []string) []string {
	for i := range paths {
		paths[i] = strings.Trim(paths[i], " ")
		log.Println(paths[i])
	}
	return paths
}
