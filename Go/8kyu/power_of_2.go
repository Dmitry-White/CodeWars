/*
Created on Wed Mar 01 17:56:15 2023
@author: Dmitry White
*/
package kata

import "math"

/*
	TODO: Complete the function that takes a non-negative integer n as input,
	and returns a list of all the powers of 2 with the exponent ranging from 0 to n (inclusive).
*/

func PowersOfTwo(n int) []uint64 {
	var result []uint64
	for i := 0; i <= n; i++ {
		result = append(result, uint64(math.Pow(2, float64(i))))
	}
	return result
}
