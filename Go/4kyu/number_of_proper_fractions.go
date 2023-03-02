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

	// fmt.Println("Cache: ", cache, len(cache))
}
