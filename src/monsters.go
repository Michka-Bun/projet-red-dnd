package dnd

import "fmt"

type Monster struct {
	Name         string
	Class        string
	HPmax        int
	HP           int
	Level        int
	PoisonEffect int
}

func ChooseMonster() Monster {
	//Classe et stats de classe(et Nom optionnel pour RP)
	var Class string
	var HPmax, HP, Level int
	if Class == "Warrior" {
		HPmax = 250
		HP = 250
		Level = 1
	} else if Class == "Mage" {
		HPmax = 150
		HP = 150
		Level = 1
	} else if Class == "Boss" {
		HPmax = 1000
		HP = 1000
		Level = 1
	}
	//Autres infos
	var PoisonEffect int

	//retour d'info
	return Monster{
		Class:        Class,
		HPmax:        HPmax,
		HP:           HP,
		Level:        Level,
		PoisonEffect: PoisonEffect,
	}
}

func DisplayMonsterInfo(m Monster) {
	fmt.Println("Informations on the monster:")
	fmt.Println("------------------------------")
	fmt.Println("Class \t\t:", m.Class)
	fmt.Println("HP max \t\t:", m.HPmax)
	fmt.Println("HP \t\t:", m.HP, " / ", m.HPmax)
	fmt.Println("Level \t\t:", m.Level)
	fmt.Println("Poison effect for :", m.PoisonEffect, "turn(s)")
}
