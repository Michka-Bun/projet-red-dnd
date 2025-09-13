package projet_red_dnd

//package main

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
	fmt.Print("Écris le nom de ton nouveau personnage : ")
	fmt.Scan(&inputName)

	// Nettoyage console
	fmt.Println(strings.Repeat("\n", 21))

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
			fmt.Println("Choissis un autre nom :")
			fmt.Scan(&inputName)
			ComptInvalidName++
			fmt.Println(strings.Repeat("\n", 21))
		}
	}

	//Classe (fonctionnel)
	var inputClass string
	var Class string
	ComptInvalidClass := 1
	for !(Class == "Guerrier" || Class == "Mage") {
		fmt.Println("Choisis la classe de ton personnage :")
		fmt.Println("1 - Guerrier")
		fmt.Println("2 - Mage")
		fmt.Scan(&inputClass)
		fmt.Println(strings.Repeat("\n", 21))

		if inputClass == "1" {
			Class = "Guerrier"
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
	if Class == "Guerrier" {
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

	//retour d'infos
	fmt.Println(strings.Repeat("\n", 21))
	fmt.Println("Informations de votre nouveau personnage :")
	fmt.Println("Name :", Name)
	fmt.Println("Class :", Class)
	fmt.Println("HP max :", HPmax)
	fmt.Println("HP :", HP)
	fmt.Println("Level :", Level)
	fmt.Println("Inventory slots :", InventorySlot)
	fmt.Println("Inventory items : Empty")

	return Player{
		Name:          Name,
		Class:         Class,
		HPmax:         HPmax,
		HP:            HP,
		Level:         Level,
		InventorySlot: InventorySlot,
		Inventory:     Inventory,
	}
}

func DisplayInfo(p Player) {
	fmt.Println("Informations de votre nouveau personnage :")
	fmt.Println("Name :", p.Name)
	fmt.Println("Class :", p.Class)
	fmt.Println("HP max :", p.HPmax)
	fmt.Println("HP :", p.HP)
	fmt.Println("Level :", p.Level)
	fmt.Println("Inventory slots :", p.InventorySlot)
	fmt.Println("Inventory items :", p.Inventory)
}

func AccessInventory(p Player) {
	fmt.Println(p.Inventory)

}
