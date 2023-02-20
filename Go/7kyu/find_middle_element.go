/*
Created on Mon Feb 20 21:08:41 2023
@author: Dmitry White
*/
package kata

import "sort"

/*
	TODO: Create a function that when provided with a triplet,
	returns the index of the numerical element that lies between the other two elements.
	The input to the function will be an array of three distinct numbers
*/

func Gimme(array [3]int) int {
	s := make([]int, len(array))
	copy(s, array[:])

	sort.Ints(s)

	var middle int
	for i, value := range array {
		if value == s[1] {
			middle = i
		}
	}
	return middle
}
