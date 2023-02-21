/*
Created on Tue Feb 21 17:56:44 2023
@author: Dmitry White
*/
package kata

import "strings"

/*
	TODO: Write a function that takes a string input,
	and returns the first character that is not repeated anywhere in the string.
*/
func FirstNonRepeating(str string) string {
	strArr := strings.Split(str, "")
	for _, char := range strArr {
		strLowerCount := strings.Count(strings.ToLower(str), char)
		strUpperCount := strings.Count(strings.ToUpper(str), char)
		if strLowerCount == 1 || strUpperCount == 1 {
			return char
		}
	}
	return ""
}
