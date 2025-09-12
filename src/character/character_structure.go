package main

import (
	"fmt"
)

func main() {
	setInfo()

}

func setInfo() string {
	//Nom (Fonctionnel)
	var inputName string
	fmt.Print("Écris le nom de ton nouveau personnage : ")
	fmt.Scan(&inputName)
	for i := 0; i <= 20; i++ {
		fmt.Println()
	}
	var ComptInvalidName = 1
	Name := ""
	for VerifName := 0; VerifName < len(inputName); VerifName++ {
		if (inputName[VerifName] >= 'A' && inputName[VerifName] <= 'Z') || (inputName[VerifName] >= 'a' && inputName[VerifName] <= 'z') || inputName[VerifName] == ' ' {
			Name = inputName
		} else {
			fmt.Println("// Invalid Name // x", ComptInvalidName)
			fmt.Println("Choissis un autre nom :")

			fmt.Scan(&inputName)
			VerifName = -1
			ComptInvalidName = ComptInvalidName + 1
			for i := 0; i <= 20; i++ {
				fmt.Println()
			}
		}
	}

	//Classe (fonctionnel)
	var inputClass string
	Class := ""
	for ComptInvalidClass := 1; !(Class == "Guerrier" || Class == "Mage"); ComptInvalidClass++ {
		fmt.Println("Choisis la classe de ton personnage :")
		fmt.Println("1 - Guerrier")
		fmt.Println("2- Mage")
		fmt.Scan(&inputClass)
		for i := 0; i <= 20; i++ {
			fmt.Println()
		}
		switch inputClass {
		case "1":
			Class = "Guerrier"
		case "2":
			Class = "Mage"
		default:
			fmt.Println("// Invalid Class // x", ComptInvalidClass)

		}
	}

	//stats en fonction de la classe
	HPmax := 0
	HP := 0
	Level := 1
	InventorySlot := 0
	switch Class {
	case "Guerrier":
		HPmax = 100
		HP = 100
		Level = 1
		InventorySlot = 5
	case "Mage":
		HPmax = 75
		HP = 75
		Level = 1
		InventorySlot = 10
	}

	//retour d'infos
	for i := 0; i <= 20; i++ {
		fmt.Println()
	}
	fmt.Println("Informations de votre nouveau personnage :")
	fmt.Println("Name :", Name)
	fmt.Println("Class :", Class)
	fmt.Println("HP max :", HPmax)
	fmt.Println("HP :", HP)
	fmt.Println("Level :", Level)
	fmt.Println("Inventory slots :", InventorySlot)

	return ""
}

/*
func displayInfo() {

}
*/
