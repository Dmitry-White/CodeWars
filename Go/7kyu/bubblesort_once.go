/*
Created on Fri Feb 24 00:40:11 2023
@author: Dmitry White
*/
package kata

/*
	TODO: Given an array of integers, your function should return a new array
	equivalent to performing exactly 1 complete pass on the original array.
	Your function should be pure, i.e. it should not mutate the input array.
*/
func BubblesortOnce(numbers []int) []int {
	var arr []int

	arr = append(arr, numbers...)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	return arr
}

func Bubblesort(numbers []int) []int {
	var arr []int

	arr = append(arr, numbers...)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
