/*
Created on Wed Mar 01 17:25:47 2023
@author: Dmitry White
*/
package kata

import "sort"

/*
	TODO: There is an array with some numbers.
	All numbers are equal except for one.
	It’s guaranteed that array contains at least 3 numbers.
	The tests contain some very huge arrays, so think about performance.
*/

func FindUniq(arr []float32) float32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	if arr[0] == arr[1] {
		return arr[len(arr)-1]
	}

	return arr[0]
}
