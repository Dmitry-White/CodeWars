/*
Created on Tue Feb 28 19:40:03 2023
@author: Dmitry White
*/
package kata

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
