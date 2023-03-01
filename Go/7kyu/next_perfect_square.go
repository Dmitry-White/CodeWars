/*
Created on Wed Mar 01 17:38:23 2023
@author: Dmitry White
*/
package kata

import "math"

/*
	TODO: Complete the method that finds the next integral perfect square
	after the one passed as a parameter.
	Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.
	If the parameter is itself not a perfect square then -1 should be returned.
	You may assume the parameter is non-negative.
*/

func FindNextSquare(sq int64) int64 {
	value := math.Sqrt(float64(sq))
	intValue := math.Trunc(value)

	if value != intValue {
		return -1
	}

	nextValue := int64(intValue + 1)

	return nextValue * nextValue
}
