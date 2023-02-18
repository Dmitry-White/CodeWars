/*
Created on Fri Feb 17 20:01:01 2023
@author: Dmitry White
*/
package kata

/*
	TODO: Implement a difference function,
	which subtracts one list from another
	and returns the result.
*/
func ArrayDiff(a, b []int) []int {
	for _, bItem := range b {

		j := 0
		for i, aItem := range a {
			if aItem != bItem {
				a[j] = a[i]
				j++
			}
		}
		a = a[:j]
	}
	return a
}
