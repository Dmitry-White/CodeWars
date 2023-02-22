/*
Created on Wed Feb 22 17:55:18 2023
@author: Dmitry White
*/
package kata

import "math"

/*
  TODO: Your task is to construct a building which will be a pile of n cubes.
  The cube at the bottom will have a volume of n^3,
  the cube above will have volume of (n-1)^3 and so on
  until the top which will have a volume of 1^3.
  You are given the total volume m of the building.
  Being given m can you find the number n of cubes you will have to build?
*/

/*
  Naive Approach: use Sum of Sequence of Cubes, m = (n(n+1)/2)^2
  Best Approach: Since n^3 + (n-1)^3...+ 1^3 = m,
  then iterating down through n substracting from m n^3 until it reaches 0.
*/

func getDiscriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}

func getRoots(a, b, discriminant float64) (float64, float64) {
	x1 := (-b + math.Sqrt(discriminant)) / 2
	x2 := (-b - math.Sqrt(discriminant)) / 2
	return x1, x2
}

func isWhole(n float64) bool {
	return n == math.Trunc(n)
}

func getPositive(x1, x2 float64) int {
	if x1 > 0 && isWhole(x1) {
		return int(x1)
	} else if x2 > 0 && isWhole(x2) {
		return int(x2)
	} else {
		return -1
	}
}

func FindNb(m int) int {
	subSqrt := math.Sqrt(float64(m))
	discriminant := getDiscriminant(1, 1, -2*subSqrt)

	x1, x2 := getRoots(1, 1, discriminant)

	return getPositive(x1, x2)
}

func BetterFindNb(m int) int {
	for n := 1; m > 0; n++ {
		m -= n * n * n
		if m == 0 {
			return n
		}
	}
	return -1
}
