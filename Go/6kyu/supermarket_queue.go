/*
Created on Tue Feb 21 18:13:04 2023
@author: Dmitry White
*/
package kata

import "sort"

/*
	TODO: There is a queue for the self-checkout tills at the supermarket.
	Write a function to calculate the total time required
	for all the customers to check out.
	Inputs:
	- customers: an array of positive integers representing the queue.
	Each integer represents a customer, and its value is the amount of time
	they require to check out.
	- n: a positive integer, the number of checkout tills.
*/

func QueueTime(customers []int, n int) int {
	tills := make([]int, n)
	for _, val := range customers {
		sort.Ints(tills)
		tills[0] += val
	}
	sort.Ints(tills)
	mostTime := tills[len(tills)-1]
	return mostTime
}
