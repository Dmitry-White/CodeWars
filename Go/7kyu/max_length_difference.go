package main

import (
	"fmt"
	"math"
)

func diffLength(s1 string, s2 string) int {
	dif := len(s1) - len(s2)
	return int(math.Abs(float64(dif)))
}

func maxer(init int) func(int) int {
	max := init
	return func(x int) int {
		if x > max {
			max = x
		}
		return max
	}
}

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	getMax := maxer(0)
	var result int

	for _, v1 := range a1 {
		for _, v2 := range a2 {
			diff := diffLength(v1, v2)
			result = getMax(diff)
		}
	}

	return result
}

func main() {
	fmt.Println(MxDifLg(
		[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
		[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
	))

	fmt.Println(MxDifLg(
		[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
		[]string{"bbbaaayddqbbrrrv"},
	))
}
