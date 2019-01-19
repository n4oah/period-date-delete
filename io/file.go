package io

import (
	"io/ioutil"
	"os"
)

func FilePathWalkDir(root string, isDir bool) []string {
	var array []string
	/*file, _ := os.Stat(root)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	if info.IsDir() == false || (info.IsDir() == true && isDir == true && os.SameFile(info, file) == false) {
		files = append(files, path)
	}
	return nil
	})*/

	files, _ := ioutil.ReadDir(root)

	for _, file := range files {
		//fmt.Println(root + "/" + file.Name())
		array = append(array, root+"/"+file.Name())
	}

	return array
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
