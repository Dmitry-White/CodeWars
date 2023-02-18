/*
Created on Fri Feb 17 20:00:01 2023
@author: Dmitry White
*/

package kata

import "sort"

/*
	TODO: Implement a function that take an array of numbers as its argument
	and return the two highest numbers within the array.
*/
func TwoOldestAges(ages []int) [2]int {
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] > ages[j]
	})

	return [2]int{ages[1], ages[0]}
}
