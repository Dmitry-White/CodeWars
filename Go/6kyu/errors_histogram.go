/*
Created on Wed May 31 21:24:01 2023
@author: Dmitry White
*/

package kata

import (
	"fmt"
	"strings"
)

/*
	TODO: Write a function `hist`` which given a string
	will output a string representing a histogram of
	the encountered runes 'u', 'w', 'x' or 'z.
	The string has a length greater or equal to one and contains only letters from a to z.
*/

const (
	ERRORS = "uwxz"
	SYMBOL = "*"
)

func hist(s string) string {
	histogram := map[rune]int{}
	output := ""

	for _, char := range s {
		isError := strings.ContainsRune(ERRORS, char)
		if !isError {
			continue
		}

		_, ok := histogram[char]
		if !ok {
			histogram[char] = 1
		} else {
			histogram[char]++
		}
	}

	for _, char := range ERRORS {
		count, ok := histogram[char]
		if !ok {
			continue
		}

		bin := strings.Repeat(SYMBOL, count)
		entry := fmt.Sprintf("%-2s %-5d %s\r", string(char), count, bin)
		output += entry
	}

	return strings.TrimSuffix(output, "\r")
}
