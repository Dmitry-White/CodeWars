/*
Created on Wed Mar 01 18:21:59 2023
@author: Dmitry White
*/
package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
	TODO: Complete the function which converts a binary number (given as a string) to a decimal number.
*/

func SourceToDecimal(source, input string) int {
	radix := len(source)
	remainder := make(map[int]int)
	quotient := input
	place := 0

	for len(quotient) != place {
		remainder[place] = sort.SearchStrings(strings.Split(source, ""), string(quotient[place]))

		place++
	}

	result := 0
	for index := 1; index <= place; index++ {
		result += remainder[place-index] * int(math.Pow(float64(radix), float64(index-1)))
	}

	return result
}

func BinToDec(bin string) int {
	value, _ := strconv.ParseInt(bin, 2, 64)
	return int(value)
}
