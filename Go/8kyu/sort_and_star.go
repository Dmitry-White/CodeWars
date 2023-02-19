/*
Created on Sun Feb 19 20:20:03 2023
@author: Dmitry White
*/
package kata

import (
	"sort"
	"strings"
)

/*
TODO: Given a list of strings, sort it alphabetically
(case-sensitive, and based on the ASCII values of the chars)
and then return the first value.
The returned value must be a string,
and have "***" between each of its letters.
Do not remove or add elements from/to the array.
*/

func modifyString(str string) string {
	strArr := strings.Split(str, "")
	modifiedStr := strings.Join(strArr, "***")
	return modifiedStr
}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	modifiedStr := modifyString(arr[0])
	return modifiedStr
}
