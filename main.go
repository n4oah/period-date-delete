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

	if checkDateCycle(period.DateCycle) == false {
		log.Println("properties DateCycle format error.")
		return
	}

	log.Println("Delete date cycle is " + period.DateCycle)

	if err == nil {
		year, month, day := getDateCycle(period.DateCycle)
		compareDate := nowDate.AddDate(year, month, day)

		log.Println(compareDate.Format("2006-01-02 15:04:05"))

		var deletePaths []string = utils.RemoveArrayInSideSpace(strings.Split(props["deletePaths"], ","))

		for _, val := range deletePaths {
			if isdirectory, _ := io.IsDirectory(val); isdirectory == true {
				compareDirInFiles(val, compareDate)
			}
		}
	}
}

func compareDirInFiles(path string, compareDate time.Time) {
	for i, value := range io.FilePathWalkDir(path, true) {
		fileInfo, _ := os.Stat(value)

		if fileInfo.IsDir() == true {
			compareDirInFiles(value, compareDate)
			continue
		}
		fileDate := fileInfo.ModTime()

		log.Println("index: " + strconv.Itoa(i) + ", value: " + value + ", CreateTime: " + fileDate.Format("2006-01-02 15:04:05"))

		if compareDate.Year() <= fileDate.Year() && compareDate.Month() <= fileDate.Month() && compareDate.Day() <= fileDate.Day() == false {
			log.Println("삭제 대상")
		}
	}
}

func checkDateCycle(str string) bool {
	for _, val := range dateCycleSuffix {
		if strings.HasSuffix(str, val) == true && utils.IsStringToInt(str[:len(str)-1]) == true {
			return true
		}
	}
	return false
}

func getDateCycle(dateCycle string) (int, int, int) {
	var year, month, day int

	number, _ := strconv.Atoi(dateCycle[:len(dateCycle)-1])
	dcs := dateCycle[len(dateCycle)-1 : len(dateCycle)]

	switch dcs {
	case "y":
		year = number
		break
	case "m":
		month = number
		break
	case "d":
		day = number
		break
	}

	return -year, -month, -day
}
