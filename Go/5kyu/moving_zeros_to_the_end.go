/*
Created on Mon Feb 20 22:14:53 2023
@author: Dmitry White
*/
package kata

/*
	TODO: Write an algorithm that takes an array
	and moves all of the zeros to the end,
	preserving the order of the other elements.
*/

func MoveZeros(arr []int) []int {
	zeros := []int{}
	nonzeros := []int{}
	for _, value := range arr {
		if value == 0 {
			zeros = append(zeros, 0)
		} else {
			nonzeros = append(nonzeros, value)
		}
	}
	return append(nonzeros, zeros...)
}
