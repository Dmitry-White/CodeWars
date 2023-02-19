/*
Created on Sun Feb 19 20:30:28 2023
@author: Dmitry White
*/
package main

import "strings"

/*
	TODO: Complete the solution so that it splits
	the string into pairs of two characters.
	If the string contains an odd number of characters
	then it should replace the missing second character
	of the final pair with an underscore ('_').
*/
const SLICE_LEN int = 2

func prepareStrArr(str string) []string {
	strArr := strings.Split(str, "")

	modStrArr := len(strArr) % SLICE_LEN

	if modStrArr != 0 {
		for i := 0; i < modStrArr; i++ {
			strArr = append(strArr, "_")
		}
	}

	return strArr
}

func sliceStrArr(strArr []string) []string {
	iterations := len(strArr)

	result := []string{}
	for i := 0; i < iterations; i += SLICE_LEN {
		slicedStrArr := strArr[i : i+SLICE_LEN]
		slicedStr := strings.Join(slicedStrArr, "")
		result = append(result, slicedStr)
	}
	return result
}

func Solution(str string) []string {
	strArr := prepareStrArr(str)
	result := sliceStrArr(strArr)
	return result
}
