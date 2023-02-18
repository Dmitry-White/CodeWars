/*
Created on Sat Feb 18 22:35:13 2023
@author: Dmitry White
*/
package kata

import "strings"

/*
	TODO: Write a function that checks if a given string (case insensitive) is a palindrome.
*/
func IsPalindrome(str string) bool {
	strSlice := strings.Split(str, "")

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}

	strReversed := strings.Join(strSlice, "")

	s1 := strings.ToUpper(str)
	s2 := strings.ToUpper(strReversed)

	return strings.Compare(s1, s2) == 0
}
