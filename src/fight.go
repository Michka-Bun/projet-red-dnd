package dnd

import (
	"fmt"
)

func Fight() {
	ChooseMonster()
}

func Attack(p Player, m Monster) {
	var AttackDamage = p.BaseDamage // multiplié par un bonus de spell, potion ou autre etc...
	if ((p.Class == "Warrior" || p.Class == "Viking") && m.Class == "Piaf") || ((p.Class == "Mage" || p.Class == "Archer") && m.Class == "Devourer") {
		fmt.Println("Your class", p.Class, "is at a disadvantage against this monster, you inflict 20% less damage0")
		AttackDamage -= AttackDamage * 20 / 100
	}
	m.HP -= AttackDamage
	if m.HP <= 0 {
		MonsterIsDead(p, m)
	} else {
		Fight()
	}
}

func MonsterAttack() {

}
