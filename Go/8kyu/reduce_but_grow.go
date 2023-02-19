/*
Created on Sun Feb 19 20:23:51 2023
@author: Dmitry White
*/
package kata

/*
	TODO: Given a non-empty array of integers,
	return the result of multiplying the values together in order.
*/

func Grow(arr []int) int {
	result := 1
	for _, value := range arr {
		result = result * value
	}
	return result
}
