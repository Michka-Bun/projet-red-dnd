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

func PoisonPot(p Player) {
	if p.PoisonEffect > 0 { //Affecter une valeur de poison directement pendant le combat // issue de la structure Player
		p.HP -= 15
		p.PoisonEffect -= 1
		fmt.Println("Tu as pris 15 HP à cause du poison, il te reste", p.PoisonEffect, "tours de poison")
	}
	if p.HP < 1 {
		//	projet_red_dnd.IsDead() //IsDead pas encore définie
	}
}
