package util


import (
	"strconv"
)


func IsLeapYear(year string) bool {
	year_num, _ := strconv.ParseInt(year, 10, 64)
	return (year_num%4 == 0 && year_num%100 != 0) || year_num%400 == 0 
}