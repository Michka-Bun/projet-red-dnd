package dnd

import (
	"fmt"
)

func Attack(p *Player, m *Monster, skill string) int {
	var AttackDamage = p.BaseDamage // multiplié par un bonus de spell, potion ou autre etc...
	switch skill {
	case "Fireball":
		AttackDamage += 10
	}
	if ((p.Class == "Warrior" || p.Class == "Viking") && m.Class == "Piaf") || ((p.Class == "Mage" || p.Class == "Archer") && m.Class == "Devourer") {
		fmt.Println("Your class", p.Class, "is at a disadvantage against this monster, you inflict 20% less damage0")
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
		return
	}
	skill := skills[idx-1]
	dmg := Attack(p, m, skill)
	m.HP -= dmg
	fmt.Printf("You used %s and dealt %d damage.\n", skill, dmg)
}

func MonsterTurn(p *Player, m *Monster) {
	dmg := m.BaseDamage
	p.HP -= dmg
	fmt.Printf("The %s hits you for %d damage.\n", m.Class, dmg)
}

func endOfRound(p *Player, m *Monster) {
	if p.PoisonEffect > 0 {
		p.HP -= 10
		p.PoisonEffect--
		fmt.Println("You take 10 poison damage.")
	}
	if m.PoisonEffect > 0 {
		m.HP -= 10
		m.PoisonEffect--
		fmt.Println("Monster takes 10 poison damage.")
	}
	fmt.Print("Press Enter to continue...")
	var _pause byte
	fmt.Scanf("%c", &_pause)
}
