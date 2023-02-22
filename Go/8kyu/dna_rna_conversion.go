/*
Created on Wed Feb 22 17:15:31 2023
@author: Dmitry White
*/
package kata

import "strings"

/*
	TODO: Create a function which translates a given DNA string into RNA.
	DNA is composed of four nucleic acid bases:
	Guanine ('G'), Cytosine ('C'), Adenine ('A'), and Thymine ('T').
	RNA differs slightly from DNA its chemical structure and contains no Thymine.
	In RNA Thymine is replaced by another nucleic acid Uracil ('U').
*/

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, string('T'), string('U'))
}
