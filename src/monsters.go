package dnd

import (
	"fmt"
	"math/rand"
	"strings"
)

type Monster struct {
	Name              string
	Class             string
	HPmax             int
	HP                int
	BaseDamage        int
	Level             int
	PoisonEffect      int
	MonsterUseHealPot bool
	DefenseTurnConter int
}

func ChooseMonster() Monster {
	// Choix de la classe
	var Class string
	ClassList := []string{"Warrior", "Mage", "Jester", "Piaf", "Devourer", "Boss"}
	RandomIndex := rand.Intn(len(ClassList))
	Class = ClassList[RandomIndex]

	//stats de classe(et Nom optionnel pour RP)
	var HPmax, HP, Level, BaseDamage, DefenseTurnConter int
	if Class == "Warrior" { //Compétence réduc dégats
		HPmax = 250
		HP = 250
		BaseDamage = 15
		Level = 1
	} else if Class == "Mage" { //potion heal et poison
		HPmax = 150
		HP = 150
		BaseDamage = 10
		Level = 1
	} else if Class == "Jester" { //potion affaiblissement et gel
		HPmax = 300
		HP = 300
		BaseDamage = 15
		Level = 1
	} else if Class == "Piaf" { // ou Sparrow //Résistance contre Player porté au sol
		HPmax = 200
		HP = 200
		BaseDamage = 10
		Level = 1
	} else if Class == "Devourer" { //Résistance contre Player Mage
		HPmax = 200
		HP = 200
		BaseDamage = 10
		Level = 1
	} else if Class == "Boss" {
		HPmax = 1000
		HP = 1000
		BaseDamage = 20
		Level = 1
	}
	//Autres infos
	var PoisonEffect int
	var MonsterUseHealPot = false

	//retour d'info
	return Monster{
		Class:             Class,
		HPmax:             HPmax,
		HP:                HP,
		Level:             Level,
		PoisonEffect:      PoisonEffect,
		BaseDamage:        BaseDamage,
		MonsterUseHealPot: MonsterUseHealPot,
		DefenseTurnConter: DefenseTurnConter,
	}
}

func DisplayMonsterInfo(m Monster) {
	fmt.Printf("\nHP \t\t: \033[1m%s%d\033[0m\t\t %s\n", MonsterHPColor(m), m.HP, MonsterHPBar(m))
	fmt.Println("Class \t\t:", m.Class)
	fmt.Println("Base damage \t:", m.BaseDamage)
	fmt.Println("Level \t\t:\033[36m\033[1m", m.Level, "\033[0m")
	if m.PoisonEffect > 0 {
		fmt.Println("Poison effect \t:", m.PoisonEffect, "\t turn(s) left")
	}
}

func MonsterHPBar(m Monster) string {
	barLength := 15
	if m.HPmax <= 0 {
		m.HPmax = 1
	}
	filled := int(float64(m.HP) / float64(m.HPmax) * float64(barLength))
	if filled < 0 {
		filled = 0
	}
	if filled > barLength {
		filled = barLength
	}

	bar := fmt.Sprintf("[%s%s]", strings.Repeat("█", filled), strings.Repeat(" ", barLength-filled))

	pct := m.HP * 100 / m.HPmax
	color := "\033[31m" //rouge
	if pct > 66 {
		color = "\033[32m" //vert
	} else if pct > 33 && pct <= 66 {
		color = "\033[33m" //jaune
	}
	return color + bar + "\033[0m"
}

// hpColor returns the ANSI color code for the current HP percentage.
func MonsterHPColor(m Monster) string {
	if m.HPmax < 0 {
		return "\033[31m" //rouge
	}
	pct := m.HP * 100 / m.HPmax
	if pct >= 66 {
		return "\033[32m" // vert
	} else if pct > 33 && pct <= 66 {
		return "\033[33m" //jaune
	}
	return "\033[31m" //rouge
}

func MonsterIsDead(p *Player, m Monster) {
	ClearScreen()
	var GoldWon, XPWon = 0, 0
	if ((p.Class == "Warrior" || p.Class == "Viking") && m.Class == "Piaf") || ((p.Class == "Mage" || p.Class == "Archer") && m.Class == "Devourer") {
		GoldWon = rand.Intn(20-15) + 15
		XPWon = rand.Intn(55-45) + 45
	} else if m.Class == "Boss" {
		GoldWon = rand.Intn(35-30) + 30
		XPWon = rand.Intn(95-80) + 80
	} else {
		GoldWon = rand.Intn(10-5) + 5
		XPWon = rand.Intn(33-27) + 27
	}
	m.PoisonEffect = 0
	fmt.Println("Congratulation, you killed that monster !")
	fmt.Println("You won :")
	fmt.Println("-", GoldWon, "golds")
	fmt.Println("-", XPWon, "XP")
	p.Gold += GoldWon
	p.XP += XPWon
	if p.XP >= p.XPmax {
		LevelUp(p)
	}
	fmt.Print("Press Enter to return to menu...")
	var _pause byte
	fmt.Scanf("%c", &_pause)
}
