package main

import (
	"fmt"
)

func main() {
	setInfo()

}

func setInfo() string {
	//Nom
	var inputName string
	fmt.Print("Écris le nom de ton nouveau personnage : ")
	fmt.Scan(&inputName)
	Name := ""
	for VerifName := 0; VerifName < len(inputName); VerifName++ {
		if (inputName[VerifName] >= 'A' && inputName[VerifName] <= 'Z') || (inputName[VerifName] >= 'a' && inputName[VerifName] <= 'z') || inputName[VerifName] == ' ' {
			Name = inputName
		} else {
			fmt.Println("// Invalid Name //")
			fmt.Println("Choissis un autre nom :")

			fmt.Scan(&inputName)
			VerifName = -1
		}
	}

	//Classe
	var inputClass string
	fmt.Println("Choisis la classe de ton personnage :")
	fmt.Println("1 - Guerrier")
	fmt.Println("2- Mage")
	fmt.Scan(&inputClass)
	Class := ""
	for Class == "Guerrier" || Class == "Mage" {
		switch inputClass {
		case "1":
			Class = "Guerrier"
		case "2":
			Class = "Mage"
		default:
			fmt.Println("Invalid Class")
			fmt.Scan(&inputClass)
		}
	}

	//retours
	fmt.Println()
	fmt.Println("Informations de votre nouveau personnage :")
	fmt.Println("Nom :", Name)
	fmt.Println("Classe :", Class)
	return ""
}

/*
func displayInfo() {

}
*/
