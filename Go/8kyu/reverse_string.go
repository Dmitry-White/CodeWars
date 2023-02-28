/*
Created on Tue Feb 28 20:27:53 2023
@author: Dmitry White
*/
package main

import "strings"

/*
	TODO: Complete the solution so that it reverses the string passed into it.
*/
func Solution(word string) string {
	strArr := make([]string, len(word))

	for i := len(word) - 1; i >= 0; i-- {
		strArr = append(strArr, string(word[i]))
	}

	return strings.Join(strArr, "")
}
