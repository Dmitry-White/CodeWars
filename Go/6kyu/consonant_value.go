/*
Created on Tue Feb 21 18:21:20 2023
@author: Dmitry White
*/
package kata

import (
	"sort"
	"strings"
)

/*
	TODO: Given a lowercase string that has alphabetic characters only and no spaces,
	return the highest value of consonant substrings.
	Consonants are any letters of the alphabet except "aeiou".
*/

const VOWELS = "aeiou"

func buildAlphabetMap() map[rune]int {
	mapChar := map[rune]int{}

	for i, r := 1, 'a'; r <= 'z'; r++ {
		mapChar[r] = i
		i++
	}

	return mapChar
}

func getConsonantValues(str string, mapChar map[rune]int) []int {
	arrChar := []int{}
	acc := 0

	for i, char := range str {
		if !strings.ContainsRune(VOWELS, char) {
			acc += mapChar[char]
		} else if acc != 0 {
			arrChar = append(arrChar, acc)
			acc = 0
		}
		if i == len(str)-1 {
			arrChar = append(arrChar, acc)
			acc = 0
		}
	}

	return arrChar
}

func solve(str string) int {
	mapChar := buildAlphabetMap()
	consonantValues := getConsonantValues(str, mapChar)

	if len(consonantValues) == 0 {
		return 0
	}

	sort.Ints(consonantValues)

	return consonantValues[len(consonantValues)-1]
}
