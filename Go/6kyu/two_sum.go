/*
Created on Sun Feb 26 13:01:02 2023
@author: Dmitry White
*/
package main

/*
	TODO: Write a function that takes an array of numbers (integers for the tests) and a target number.
	It should find two different items in the array that, when added together, give the target value.
	The indices of these items should then be returned in a tuple / list like so: (index1, index2).
	The input will always be valid:
	 - numbers will be an array of length 2 or greater, and all of the items will be numbers;
	 - target will always be the sum of two different items from that array.
*/

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{}
}
