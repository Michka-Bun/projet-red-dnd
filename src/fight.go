package dnd

import (
	"fmt"
	"math/rand"
)

func Fight() {
	ChooseMonster()
}

func Attack(p Player, m Monster) {
	var AttackDamage = p.BaseDamage // multiplié par un bonus de spell, potion ou autre etc...
	if ((p.Class == "Warrior" || p.Class == "Viking") && m.Class == "Piaf") || ((p.Class == "Mage" || p.Class == "Archer") && m.Class == "Devourer") {
		fmt.Println("Your class", p.Class, "is at a disadvantage against this monster, you inflict 20% less damage")
		AttackDamage -= AttackDamage * 20 / 100
	} else if m.DefenseTurnConter != 0 {
		fmt.Println(m.Class, "protected himself, you inflict 50% less damage")
		AttackDamage -= AttackDamage * 30 / 100
	} else if p.WeakeningTrunCount != 0 {
		AttackDamage -= AttackDamage * 20 / 100
	}

	m.HP -= AttackDamage
	if m.HP <= 0 {
		MonsterIsDead(p, m)
	} else {
		Fight()
	}
}

func MonsterAttack(p Player, m Monster) {
	switch m.Class {
	case "Warrior":
		m.DefenseTurnConter = 0
		RdmAttack := rand.Intn(2-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(&p)
			}
		case 2: //compétence protection 1 tour
			m.DefenseTurnConter = 1
		}
	case "Mage":
		RdmAttack := rand.Intn(3-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(&p)
			}
		case 2: //Potion de soins
			fmt.Println(m.Class, "use a heal potion, he recovers 10% HP")
			if m.HP+m.HPmax*10/100 <= m.HPmax {
				m.HP += m.HPmax * 10 / 100
			} else {
				m.HP = m.HPmax
			}
		case 3: //Potion de poison
			fmt.Println(m.Class, "throws a poison potion, you are poisoned for 3 turn")
			p.PoisonEffect = 3
		}
	case "Jester":
		RdmAttack := rand.Intn(3-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(&p)
			}
		case 2: //Potion d'affaiblissement
			fmt.Println(m.Class, "throws a weakness potion, you are weakened for 1 turn")
			p.WeakeningTrunCount = 1
		case 3: //Potion de gel
			fmt.Println(m.Class, "throws a freezing potion, you can't play for 1 turn")
			MonsterAttack(p, m)
		}
	case "Piaf":
		p.HP -= m.BaseDamage
		if p.HP <= 0 {
			IsDead(&p)
		}
	case "Devourer":
		p.HP -= m.BaseDamage
		if p.HP <= 0 {
			IsDead(&p)
		}
	case "Boss":
		RdmAttack := rand.Intn(5-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(&p)
			}
		case 2: //Potion de soins
			fmt.Println(m.Class, "use a heal potion, he recovers 10% HP")
			if m.HP+m.HPmax*5/100 <= m.HPmax {
				m.HP += m.HPmax * 5 / 100
			} else {
				m.HP = m.HPmax
			}
		case 3: //Potion de poison
			fmt.Println(m.Class, "throws a poison potion, you are poisoned for 3 turn")
			p.PoisonEffect = 3
		case 4:
			fmt.Println(m.Class, "throws a weakness potion, you are weakened for 1 turn")
			p.WeakeningTrunCount = 1
		case 5:
			fmt.Println(m.Class, "throws a freezing potion, you can't play for 1 turn")
			MonsterAttack(p, m)
		}
	}
}
