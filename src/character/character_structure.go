package projet_red_dnd

import (
	"fmt"
)

func SetInfo() {
	//Nom (Fonctionnel)
	inputName := ""
	fmt.Print("Écris le nom de ton nouveau personnage : ")
	fmt.Scan(&inputName)
	for i := 0; i <= 20; i++ {
		fmt.Println()
	}
	ComptInvalidName := 1
	var Name string
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
	inputClass := ""
	var Class string
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
	var HPmax int
	var HP int
	var Level int
	var InventorySlot []string
	switch Class {
	case "Guerrier":
		HPmax = 100
		HP = 100
		Level = 1
		InventorySlot = []string{}
	case "Mage":
		HPmax = 75
		HP = 75
		Level = 1
		InventorySlot = []string{}
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
}

func DisplayInfo(
	Name string,
	Class string,
	HPmax int,
	HP int,
	Level int,
	InventorySlot int) {

	fmt.Println("Informations de votre nouveau personnage :")
	fmt.Println("Name :", Name)
	fmt.Println("Class :", Class)
	fmt.Println("HP max :", HPmax)
	fmt.Println("HP :", HP)
	fmt.Println("Level :", Level)
	fmt.Println("Inventory slots :", InventorySlot)
}

type Player struct {
	name          string
	Class         string
	HPmax         int
	HP            int
	Level         int
	InventorySlot int
}
