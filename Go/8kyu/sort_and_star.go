package kata

import (
	"sort"
	"strings"
)

func sortStrings(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func modifyString(str string) string {
	strArr := strings.Split(str, "")
	modifiedStr := strings.Join(strArr, "***")
	return modifiedStr
}

func TwoSort(arr []string) string {
	sortedArr := sortStrings(arr)
	modifiedStr := modifyString(sortedArr[0])
	return modifiedStr
}
