package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func ReformatDate(date string) string {
	months := map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6, "Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}
	arr := strings.Split(date, " ")
	month := months[arr[1]]
	day, _ := strconv.Atoi(arr[0][:len(arr[0])-2])
	return fmt.Sprintf("%s-%02d-%02d", arr[2], month, day)
}
