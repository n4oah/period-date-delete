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
var nowDate time.Time = time.Now() //.AddDate(0, 0, -7)

func main() {
	log.Println("Period Delete Process Start.")

	var period *vo.Period = new(vo.Period)
	props, err := io.ReadPropertiesFile(propertiy)

	period.DateCycle = props["dateCycle"]
	period.DeletePaths = props["deletePaths"]
	period.LogFilePath = props["logFilePath"]
	period.DeleteOfDirectory = props["deleteOfDirectory"]

	if CheckDateCycle(period.DateCycle) == false {
		log.Println("properties DateCycle format error.")
		return
	}

	log.Println("Delete date cycle is " + period.DateCycle)

	if err == nil {
		var deletePaths []string = utils.RemoveArrayInSideSpace(strings.Split(props["deletePaths"], ","))

		for _, val := range deletePaths {
			if isdirectory, _ := io.IsDirectory(val); isdirectory == true {
				CompareDirInFiles(val, period.DateCycle)
			}
		}
	}
}

func CompareDirInFiles(path string, dateCycle string) {
	for i, value := range io.FilePathWalkDir(path, true) {
		fileInfo, _ := os.Stat(value)

		if fileInfo.IsDir() == true {
			CompareDirInFiles(value, dateCycle)
		} else {
			fileDate := fileInfo.ModTime()

			log.Println("index: " + strconv.Itoa(i) + ", value: " + value + ", CreateTime: " + fileDate.Format("2006-01-02 15:04:05"))

			log.Println(nowDate.Format("2006-01-02 15:04:05"))
		}
	}
}

func CheckDateCycle(str string) bool {
	for _, val := range dateCycleSuffix {
		if strings.HasSuffix(str, val) == true && utils.IsStringToInt(str[:len(str)-1]) == true {
			return true
		}
	}
	return false
}
