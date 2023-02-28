/*
Created on Sat Feb 25 19:58:16 2023
@author: Dmitry White
*/
package kata

import "strings"

/*
	TODO: Return the number (count) of vowels in the given string.
	We will consider 'a, e, i, o, u' as vowels for this Kata (but not 'y').
	The input string will only consist of lower case letters and/or spaces.
*/
const vowels = "aeiou"

func GetCount(str string) (count int) {
	for _, char := range str {
		if strings.Contains(vowels, string(char)) {
			count++
		}
	}
	return count
}
