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

var complementaryDNA = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

func DNAStrand(dna string) string {
	return complementaryDNA.Replace(dna)
}
