/*
Created on Mon Feb 20 21:34:22 2023
@author: Dmitry White
*/
package main

import (
	"sort"
	"strings"
)

/*
	TODO: Given two arrays of strings a1 and a2
 	return a sorted array r in lexicographical order
 	of the strings of a1 which are substrings of strings of a2.
*/

func deduplicateStr(strSlice []string) []string {
	keys := make(map[string]bool)
	set := []string{}
	for _, item := range strSlice {
		_, ok := keys[item]
		if !ok {
			keys[item] = true
			set = append(set, item)
		}
	}
	return set
}

func InArray(array1 []string, array2 []string) []string {
	result := []string{}

	for _, value1 := range array1 {
		for _, value2 := range array2 {
			if strings.Contains(value2, value1) {
				result = append(result, value1)
			}
		}
	}

	result = deduplicateStr(result)

	sort.Strings(result)

	return result
}
