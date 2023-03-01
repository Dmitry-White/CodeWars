/*
Created on Wed Mar 01 17:55:28 2023
@author: Dmitry White
*/
package kata

/*
	TODO: In this game, the hero moves from left to right.
	The player rolls the dice and moves the number of spaces indicated by the dice two times.
	Create a function for the terminal game that takes the current position of the hero
	and the roll (1-6) and return the new position.
*/

func Move(position int, roll int) int {
	return position + roll*2
}
