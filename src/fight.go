package dnd

import (
	"fmt"
	"math/rand"
)

func Attack(p *Player, m *Monster, skill string) int {
	var AttackDamage = p.BaseDamage // multiplié par un bonus de spell, potion ou autre etc...
	switch skill {
	case "Fireball":
		AttackDamage += AttackDamage * 33 / 100
	}
	if ((p.Class == "Warrior" || p.Class == "Viking") && m.Class == "Piaf") || ((p.Class == "Mage" || p.Class == "Archer") && m.Class == "Devourer") {
		fmt.Println("Your class", p.Class, "is at a disadvantage against this monster, you inflict 20% less damage")
		AttackDamage -= AttackDamage * 20 / 100
	} else if m.DefenseTurnConter != 0 {
		AttackDamage -= AttackDamage * 30 / 100
	} else if p.WeakeningTrunCount != 0 {
		AttackDamage -= AttackDamage * 20 / 100
	}
	if AttackDamage < 1 {
		AttackDamage = 1
	}
	return AttackDamage
}

func ChooseAttack(p *Player, m *Monster) {
	// Build the list of available attacks from learned skills
	skills := []string{"Punch"}
	if HasSkill(p, "Fireball") {
		skills = append(skills, "Fireball")
	}
	fmt.Println("Choose your attack:")
	for i, s := range skills {
		fmt.Printf("%d) %s\n", i+1, s)
	}
	fmt.Print("Select: ")
	var idx int
	fmt.Scanln(&idx)
	if idx < 1 || idx > len(skills) {
		fmt.Println("Invalid attack.")
		ChooseAttack(p, m)
		return

	} else if idx == 2 {
		if !(p.Mana >= 30) {
			fmt.Println("Insufficient mana.")
			ChooseAttack(p, m)
		} else {
			p.Mana -= 40
		}

	}
	skill := skills[idx-1]
	dmg := Attack(p, m, skill)
	m.HP -= dmg
	fmt.Printf("You used %s and dealt %d damage.\n", skill, dmg)
}

func MonsterTurn(p *Player, m *Monster) {
	switch m.Class {
	case "Warrior":
		m.DefenseTurnConter = 0
		RdmAttack := rand.Intn(4-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(p)
			}
		case 2: //compétence protection 1 tour
			fmt.Println(m.Class, "protected himself, you inflict 50% less damage for 1 turn")
			m.DefenseTurnConter = 1
		case 3: // réduction du mana du joueur
			fmt.Println(m.Class, "sucks 7 of your mana")
			p.Mana -= 7
			if p.Mana < 0 {
				p.Mana = 0
			}
		}
	case "Mage":
		RdmAttack := rand.Intn(4-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(p)
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
		RdmAttack := rand.Intn(4-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(p)
			}
		case 2: //Potion d'affaiblissement
			fmt.Println(m.Class, "throws a weakness potion, you are weakened for 1 turn")
			p.WeakeningTrunCount = 2
		case 3: //Potion de gel
			fmt.Println(m.Class, "throws a freezing potion, you can't play for 1 turn and take 10 of damage")
			p.HP -= 10
			if p.HP <= 0 {
				IsDead(p)
			}
			MonsterTurn(p, m)
		}
	case "Piaf":
		fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
		p.HP -= m.BaseDamage
		if p.HP <= 0 {
			IsDead(p)
		}
	case "Devourer":
		fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
		p.HP -= m.BaseDamage
		if p.HP <= 0 {
			IsDead(p)
		}
	case "Boss":
		RdmAttack := rand.Intn(7-1) + 1
		switch RdmAttack {
		case 1: //attaque classique
			fmt.Println(m.Class, "hit you, you lost", m.BaseDamage, "HP")
			p.HP -= m.BaseDamage
			if p.HP <= 0 {
				IsDead(p)
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
			p.WeakeningTrunCount = 2
		case 5:
			fmt.Println(m.Class, "throws a freezing potion, you can't play for 1 turn and take 10 of damage")
			p.HP -= 10
			if p.HP <= 0 {
				IsDead(p)
			}
			MonsterTurn(p, m)
		case 6:
			fmt.Println(m.Class, "sucks 15 of your mana")
			p.Mana -= 15
			if p.Mana < 0 {
				p.Mana = 0
			}
		}
	}
	//fmt.Printf("The %s hits you for %d damage.\n", m.Class, dmg)
}

func endOfRound(p *Player, m *Monster) {
	if p.PoisonEffect > -1 {
		p.HP -= 10
		p.PoisonEffect--
		fmt.Println("You take 10 poison damage.")
	}
	if m.PoisonEffect > 0 {
		m.HP -= 15
		m.PoisonEffect--
		fmt.Println("Monster takes 15 poison damage.")
	}
	if p.WeakeningTrunCount > -1 {
		p.WeakeningTrunCount--
		fmt.Println("You dealt 20% less damage due to the weakness potion")
	}
	if p.Class == "Mage" {
		if p.Mana+20 <= p.Manamax {
			p.Mana += 20
		} else {
			p.Mana = p.Manamax
		}
	} else if p.Class == "Archer" {
		if p.Mana+15 <= p.Manamax {
			p.Mana += 15
		} else {
			p.Mana = p.Manamax
		}
	} else {
		if p.Mana+10 <= p.Manamax {
			p.Mana += 10
		} else {
			p.Mana = p.Manamax
		}
	}
	fmt.Print("Press Enter to continue...")
	var _pause byte
	fmt.Scanf("%c", &_pause)
}
