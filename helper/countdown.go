package helper

import (
	"strconv"
	"strings"
	"time"
)

// accepted input: yyyy/mm/dd hh:mm:ss
// TODO: also accept only yyyy/mm/dd; yyyy/mm/dd hh:mm
func ValidateDateTime(dateTime []string) {

}

func ParseDateTime(dateTime []string) time.Time {
	var dateObj time.Time

	switch len(dateTime) {
	case 2:
		dateStrArr := strings.Split(dateTime[0], "/")
		var dateIntArr = []int{}

		for _, n := range dateStrArr {
			dateInt, err := strconv.Atoi(n)
			if err != nil {
				OutputError(err.Error())
			}
			dateIntArr = append(dateIntArr, dateInt)
		}

		timeStrArr := strings.Split(dateTime[1], ":")
		var timeIntArr = []int{}
		for _, n := range timeStrArr {
			timeInt, err := strconv.Atoi(n)
			if err != nil {
				OutputError(err.Error())
			}
			timeIntArr = append(timeIntArr, timeInt)
		}

		dateObj = time.Date(dateIntArr[0], time.Month(dateIntArr[1]), dateIntArr[2], timeIntArr[0], timeIntArr[1], timeIntArr[2], 0, time.Local)
	case 1:
		nowDate := time.Now()
		timeStrArr := strings.Split(dateTime[0], ":")
		var timeIntArr = []int{}
		for _, n := range timeStrArr {
			timeInt, err := strconv.Atoi(n)
			if err != nil {
				OutputError(err.Error())
			}
			timeIntArr = append(timeIntArr, timeInt)
		}
		dateObj = time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), timeIntArr[0], timeIntArr[1], timeIntArr[2], 0, time.Local)
	}
	return dateObj
}
