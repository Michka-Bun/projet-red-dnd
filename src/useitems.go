package dnd

import (
	"fmt"
)

func TakeHealPot(p Player) {
	if p.HP < p.HPmax/2 { // issue de la structure Player
		p.HP += p.HPmax / 2
	} else {
		p.HP = p.HPmax
	}
}

func PoisonPot(p Player, m Monster) {
	if p.PoisonEffect > 0 { //Affecter une valeur de poison directement pendant le combat // issue de la structure Player
		p.HP -= 15
		p.PoisonEffect -= 1
		fmt.Println("You took 15 HP damage from the poison, you have", p.PoisonEffect, "turn(s) of poison left")
		if p.HP < 1 {
			//	dnd.IsDead() //IsDead pas encore définie
		}
	}

	if m.PoisonEffect > 0 { //Affecter une valeur de poison directement pendant le combat // issue de la structure Monster (non faite)
		m.HP -= 15
		m.PoisonEffect -= 1
		fmt.Println("The monster took 15 HP damage from the poison, he has", m.PoisonEffect, "turn(s) of poison left")
		if m.HP < 1 {
			//	dnd.MonsterIsDead() //MonsterIsDead pas encore définie
		}
	}
}

func TakePoisonPot(p Player) {
	fmt.Println("You did self-sabotage and drank a poison potion. You take 30 damage.")
	p.HP -= 30
	if p.HP < 1 {
		//	dnd.IsDead() //IsDead pas encore définie
	}
}
