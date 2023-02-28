/*
Created on Sun Feb 26 16:30:52 2023
@author: Dmitry White
*/
package kata

/*
	TODO: In a small town the population is `p0` at the beginning of a year.
	The population regularly increases by `percent`` per year
	and moreover `aug` new inhabitants per year come to live in the town.
	How many years does the town need to see its population greater or equal to `p` inhabitants?
*/

func NbYear(p0 int, percent float64, aug int, p int) int {
	var n int
	perYear := (100 + percent) / 100
	currP := p0

	for n = 0; currP < p; n++ {
		currP = int(float64(currP)*perYear) + aug
	}

	return n
}
