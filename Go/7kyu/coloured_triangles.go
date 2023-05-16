/*
Created on Tue May 16 20:56:02 2023
@author: Dmitry White
*/

package kata

/*
	TODO: A coloured triangle is created from a row of colours, each of which is red, green or blue.
	Successive rows, each containing one fewer colour than the last, are generated
	by considering the two touching colours in the previous row.
	If these colours are identical, the same colour is used in the new row.
	If they are different, the missing colour is used in the new row.
	This is continued until the final row, with only a single colour, is generated.
	Given the first row of the triangle as a string, return the final colour
	which would appear in the bottom row as a string.
*/

var colours = []rune{'R', 'G', 'B'}

func analyze(first, second rune) rune {
	if first == second {
		return first
	}

	var third rune

	for _, colour := range colours {
		if colour != first && colour != second {
			return colour
		}
	}

	return third
}

func Triangle(row string) rune {
	currentRow := []rune(row)

	for len(currentRow) > 1 {
		newRow := []rune{}

		for i := 0; i < len(currentRow)-1; i++ {
			first := currentRow[i]
			second := currentRow[i+1]

			third := analyze(first, second)

			newRow = append(newRow, third)
		}

		currentRow = newRow
	}

	return currentRow[0]
}
