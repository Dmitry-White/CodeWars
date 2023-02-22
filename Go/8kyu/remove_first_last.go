/*
Created on Wed Feb 22 18:03:42 2023
@author: Dmitry White
*/

package kata

/*
	TODO: create a function that removes the first and last characters of a string.
	You're given one parameter, the original string.
	You don't have to worry with strings with less than two characters.
*/

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
