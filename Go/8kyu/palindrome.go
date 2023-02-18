/*
Created on Sat Feb 18 22:50:34 2023
@author: Dmitry White
*/
package kata

import "strings"

/*
	TODO: Write a function that checks if a given string (case insensitive) is a palindrome.
*/
func IsPalindrome(str string) bool {
	strNormalized := strings.ToLower(str)
	strSlice := strings.Split(strNormalized, "")

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}

	strReversed := strings.Join(strSlice, "")

	return strings.Compare(strNormalized, strReversed) == 0
}
