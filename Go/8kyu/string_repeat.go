/*
Created on Wed Feb 22 17:57:03 2023
@author: Dmitry White
*/
package kata

/*
	TODO: Write a function that accepts an integer n and a string s as parameters,
	and returns a string of s repeated exactly n times.
*/

import "strings"

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}
