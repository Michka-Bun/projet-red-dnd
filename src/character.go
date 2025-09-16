package dnd

import (
	"fmt"
	"strings"
)

type Player struct {
	Name          string
	Class         string
	HPmax         int
	HP            int
	Level         int
	InventorySlot int
	Inventory     []string
	PoisonEffect  int
}

/*
func main() {
    player := SetInfo()
    DisplayInfo(player)
}
*/

func SetInfo() Player {
	//Nom (Fonctionnel)
	var inputName string
	fmt.Print("Write the name of your new character: ")
	fmt.Scan(&inputName)

	// Nettoyage console
	fmt.Print("\033[2J\033[3J\033[H")

	ComptInvalidName := 1
	var Name string
	valid := false
	for !valid {
		valid = true
		for i := 0; i < len(inputName); i++ {
			c := inputName[i]
			if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == ' ') {
				valid = false
				break
			}
		}
		if valid {
			Name = inputName
		} else {
			fmt.Println("// Invalid Name // x", ComptInvalidName)
			fmt.Println("Choose another name:")
			fmt.Scan(&inputName)
			ComptInvalidName++
			fmt.Println(strings.Repeat("\n", 21))
		}
	}

	//Classe (fonctionnel)
	var inputClass string
	var Class string
	ComptInvalidClass := 1
	for !(Class == "Warrior" || Class == "Mage") {
		fmt.Println("Choose your character's class:")
		fmt.Println("1 - Warrior")
		fmt.Println("2 - Mage")
		fmt.Scan(&inputClass)
		fmt.Println(strings.Repeat("\n", 21))

		if inputClass == "1" {
			Class = "Warrior"
		} else if inputClass == "2" {
			Class = "Mage"
		} else {
			fmt.Println("// Invalid Class // x", ComptInvalidClass)
			ComptInvalidClass++
		}
	}

	//stats en fonction de la classe
	var HPmax, HP, Level int
	var InventorySlot int
	var Inventory []string
	if Class == "Warrior" {
		HPmax = 100
		HP = 100
		Level = 1
		InventorySlot = 5
		Inventory = []string{}
	} else if Class == "Mage" {
		HPmax = 75
		HP = 75
		Level = 1
		InventorySlot = 10
		Inventory = []string{}
	}
	//Autres infos
	var PoisonEffect int

	//retour d'infos
	fmt.Println(strings.Repeat("\n", 21))
	fmt.Println("Informations on your new character:")
	fmt.Println("Name \t\t:", Name)
	fmt.Println("Class \t\t:", Class)
	fmt.Println("HP max \t\t:", HPmax)
	fmt.Println("HP \t\t:", HP)
	fmt.Println("Level \t\t:", Level)
	fmt.Println("Inventory slots :", InventorySlot)
	fmt.Println("Inventory items :", Inventory)

	return Player{
		Name:          Name,
		Class:         Class,
		HPmax:         HPmax,
		HP:            HP,
		Level:         Level,
		InventorySlot: InventorySlot,
		Inventory:     Inventory,
		PoisonEffect:  PoisonEffect,
	}
}

func DisplayInfo(p Player) {
	fmt.Println("Informations on your character:")
	fmt.Println("------------------------------")
	fmt.Println("Name \t\t:", p.Name)
	fmt.Println("Class \t\t:", p.Class)
	fmt.Println("HP max \t\t:", p.HPmax)
	fmt.Println("HP \t\t:", p.HP)
	fmt.Println("Level \t\t:", p.Level)
	fmt.Println("Inventory slots :", p.InventorySlot)
	fmt.Println("Inventory items :", p.Inventory)
	fmt.Println("Poison effect for :", p.PoisonEffect, "turn(s)")
}

func AccessInventory(p Player) {
	fmt.Println(p.Inventory)

}
