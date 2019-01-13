package main

import (
	"log"
	"os"
	"period/io"
	"period/utils"
	"period/vo"
	"strconv"
	"strings"
	"time"
)

var dateCycleSuffix = [...]string{"d", "m", "y"}

var propertiy string = "period_delete.properties"
var nowDate time.Time = time.Now()

func main() {
	log.Println("Period Delete Process Start.")

	var period *vo.Period = new(vo.Period)
	props, err := io.ReadPropertiesFile(propertiy)

	period.DateCycle = props["dateCycle"]
	period.DeletePaths = props["deletePaths"]
	period.LogFilePath = props["logFilePath"]

	log.Println("Delete date cycle is " + period.DateCycle)

	if err == nil {
		var deletePaths []string = utils.RemoveArrayInSideSpace(strings.Split(props["deletePaths"], ","))

		for _, val := range deletePaths {
			if isdirectory, _ := io.IsDirectory(val); isdirectory == true {
				for i, value := range io.FilePathWalkDir(val) {
					fileInfo, _ := os.Stat(value)

					log.Println("index: " + strconv.Itoa(i) + ", value: " + value + ", CreateTime: " + fileInfo.ModTime().Format("2006-01-02 15:04:05"))

				}
			}
		}
	}
}

func CheckDateCycle(str string) bool {
	for _, val := range dateCycleSuffix {
		if strings.HasSuffix(str, val) == false || utils.IsStringToInt(str[:len(str)-1]) == false {
			return false
		}
	}
	return true
}
