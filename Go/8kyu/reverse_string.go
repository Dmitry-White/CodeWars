/*
Created on Tue Feb 28 20:27:53 2023
@author: Dmitry White
*/
package main

/*
	TODO: Complete the solution so that it reverses the string passed into it.
*/
func Solution(word string) string {
	resultWord := ""

	for _, char := range word {
		resultWord = string(char) + resultWord
	}

	return resultWord
}
