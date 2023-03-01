/*
Created on Wed Mar 01 18:10:47 2023
@author: Dmitry White
*/
package kata

import "strings"

/*
	TODO: In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G".
	Your function receives one side of the DNA;
	you need to return the other complementary side.
	DNA strand is never empty or there is no DNA at all.
*/

func DNAStrand(dna string) string {
	strArr := strings.Split(dna, "")

	for i := range strArr {
		switch strArr[i] {
		case "A":
			strArr[i] = "T"
		case "T":
			strArr[i] = "A"
		case "C":
			strArr[i] = "G"
		case "G":
			strArr[i] = "C"
		}
	}

	return strings.Join(strArr, "")
}
