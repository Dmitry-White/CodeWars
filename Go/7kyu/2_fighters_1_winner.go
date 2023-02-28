/*
Created on Tue Feb 28 19:40:03 2023
@author: Dmitry White
*/
package main

import "fmt"

/*
	TODO: Create a function that returns the name of the winner in a fight between two fighters.
	Each fighter takes turns attacking the other and whoever kills the other first is victorious.
	Both health and damagePerAttack will be integers larger than 0.
	Death is defined as having health <= 0.
	Function also receives a third argument, a string, with the name of the fighter that attacks first.
*/

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	winner := firstAttacker

	fighter1Win := 0
	fighter2Win := 0

	for i := 0; fighter2.Health > 0; i++ {
		fighter2.Health -= fighter1.DamagePerAttack
		fighter1Win++
	}

	for i := 0; fighter1.Health > 0; i++ {
		fighter1.Health -= fighter2.DamagePerAttack
		fighter2Win++
	}

	if firstAttacker == fighter1.Name {
		fighter1Win--
		if fighter1Win < fighter2Win {
			winner = fighter1.Name
		} else {
			winner = fighter2.Name
		}
	} else if firstAttacker == fighter2.Name {
		fighter2Win--
		if fighter2Win < fighter1Win {
			winner = fighter2.Name
		} else {
			winner = fighter1.Name
		}
	}

	return winner
}

func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))        // Lew
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))      // Harry
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"))   // Harald
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald"))  // Harald
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry"))  // Harald
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald")) // Harald
}
