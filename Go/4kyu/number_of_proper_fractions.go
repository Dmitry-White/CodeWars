/*
Created on Wed Mar 01 20:32:16 2023
@author: Dmitry White
*/
package main

import (
	"fmt"
	"time"
)

var cache = make(map[int]int)

type memoized struct {
	f     func(int, int) int
	cache map[int]int
}

func (m *memoized) call(a, b int) int {
	key := getCommutativeCantorPair(a, b)
	if v, ok := m.cache[key]; ok {
		return v
	}

	newValue := m.f(b, a%b)
	newKey := getCommutativeCantorPair(b, a%b)
	m.cache[newKey] = newValue

	return newValue
}

func memoize(f func(int, int) int) *memoized {
	return &memoized{f: f, cache: make(map[int]int)}
}

// https://gist.github.com/jorimvanhove/cc706a8b9dd4d98350476d28c1ccf9c4
// https://en.wikipedia.org/wiki/Pairing_function
func getCommutativeCantorPair(a, b int) int {
	var x int
	var y int

	if a <= b {
		x = a
		y = b
	} else {
		x = b
		y = a
	}

	return (((x + y) * (x + y + 1)) / 2) + y
}

// Remainder Recursion Memoized
func GCDRecursionMemoized(a, b int) int {
	key := getCommutativeCantorPair(a, b)
	value, ok := cache[key]
	if ok {
		return value
	}

	if b == 0 {
		return a
	} else {
		newValue := GCDRecursionMemoized(b, a%b)
		newKey := getCommutativeCantorPair(b, a%b)
		cache[newKey] = newValue

		return newValue
	}
}

// Remainder Recursion
func GCDRecursion(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCDRecursion(b, a%b)
	}
}

// Euclidean algorithm: Subtraction
func GCDLoop(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

// Euclidean algorithm: Swap
func GCDSwapLoop(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ProperFractions(n int) int {
	denominator := n
	count := 0

	for numerator := 1; numerator < denominator; numerator++ {
		if GCDLoop(numerator, denominator) == 1 {
			count++
		}
	}

	return count
}

func ProperFractionsMemo(n int) int {
	result := memoize(GCDRecursion)
	denominator := n
	count := 0

	for numerator := 1; numerator < denominator; numerator++ {
		if result.call(numerator, denominator) == 1 {
			count++
		}
	}

	return count
}

func with(algorithm func(int, int) int) func(int) int {
	return func(n int) int {
		denominator := n
		count := 0

		for numerator := 1; numerator < denominator; numerator++ {
			if algorithm(numerator, denominator) == 1 {
				count++
			}
		}

		return count
	}
}

/*
Euler's totient function

	Euler’s totient function counts the total numbers between 1 to N, which are coprime to `N`.
	Two numbers are said to be coprime if the GCD(Greatest Common Divisor) of both the numbers is 1.

	https://en.wikipedia.org/wiki/Euler%27s_totient_function
	https://www.codingninjas.com/codestudio/library/euler-s-totient-function


	There are generally two approaches to this function:
	1. Iteratively counting the numbers `k ≤ n` such that `gcd(n,k) = 1`.
	2. Using the Euler product formula.

	Usage of the Euler product formula:
	This is an explicit formula for calculating `φ(n)` depending on the prime divisor of n:
	`φ(n) = n * Product (1 - 1/p)` where the product is taken over the primes `p ≤ n` that divide `n`.
	For example: `φ(36) = 36 * (1 - 1/2) * (1 - 1/3) = 36 * 1/2 * 2/3 = 12`.

	Time Complexity: The time complexity is `O(√N)` because we are running an `O(√N)` for loop only.
	Space Complexity: The space complexity is `O(1)` because we are using constant auxiliary space.
*/
func Phi(n int64) int64 {
	// Initialize result as n
	result := n

	// Remove all factors p of n where p < sqrt(n)
	for p := int64(2); p*p <= n; p++ {
		// Check if p is a factor of n.
		if n%p == 0 {
			// Remove factor p and all its multiples from n
			for n%p == 0 {
				n /= p
			}
			// Count multiples of p and remove from result
			result -= (result / p)
		}
	}

	// If a factor still exists, this must be a relative prime
	// (since all factors < sqrt(n) has been removed)
	if n > 1 {
		// Remove all its multiples from result
		result -= (result / n)
	}

	return result
}

func ProperFractionsEuler(n int) int {
	if n == 1 {
		return 0
	}

	count := Phi(int64(n))

	return int(count)
}

func main() {
	fmt.Println(ProperFractions(1))  // 0
	fmt.Println(ProperFractions(2))  // 1
	fmt.Println(ProperFractions(5))  // 4
	fmt.Println(ProperFractions(15)) // 8
	fmt.Println(ProperFractions(25)) // 20

	start1 := time.Now()
	fmt.Println(with(GCDLoop)(9999999))
	elapsed1 := time.Since(start1)
	fmt.Println("Time 1: ", elapsed1)

	start2 := time.Now()
	fmt.Println(with(GCDSwapLoop)(9999999))
	elapsed2 := time.Since(start2)
	fmt.Println("Time 2: ", elapsed2)

	start3 := time.Now()
	fmt.Println(with(GCDRecursion)(9999999))
	elapsed3 := time.Since(start3)
	fmt.Println("Time 3: ", elapsed3)

	start4 := time.Now()
	fmt.Println(with(GCDRecursionMemoized)(9999999))
	elapsed4 := time.Since(start4)
	fmt.Println("Time 4: ", elapsed4)

	start5 := time.Now()
	fmt.Println(ProperFractionsMemo(9999999))
	elapsed5 := time.Since(start5)
	fmt.Println("Time 5: ", elapsed5)

	start6 := time.Now()
	fmt.Println(ProperFractionsEuler(9999999))
	elapsed6 := time.Since(start6)
	fmt.Println("Time 6: ", elapsed6)

	fmt.Println("Cache Size: ", len(cache))
}
