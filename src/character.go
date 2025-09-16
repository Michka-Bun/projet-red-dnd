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
	Gold          int
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
	var Gold int
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
	Gold = 0

	//retour d'infos
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Println("Informations on your new character:")
	fmt.Println("Name \t\t:", Name)
	fmt.Println("Class \t\t:", Class)
	fmt.Println("HP max \t\t:", HPmax)
	fmt.Println("HP \t\t:", HP)
	fmt.Println("Level \t\t:", Level)
	fmt.Println("Gold \t\t:", Gold)
	fmt.Println("Inventory slots :", InventorySlot)
	fmt.Println("Inventory items :", Inventory)

	return Player{
		Name:          Name,
		Class:         Class,
		HPmax:         HPmax,
		HP:            HP,
		Level:         Level,
		Gold:          Gold,
		InventorySlot: InventorySlot,
		Inventory:     Inventory,
	}
}

func DisplayInfo(p Player) {
	fmt.Println("Informations on your character:")
	fmt.Println("------------------------------")
	fmt.Println("Name \t\t:", p.Name)
	fmt.Println("Class \t\t:", p.Class)
	fmt.Println("HP max \t\t:", p.HPmax)
	fmt.Println("HP \t\t:", p.HP, "\t\t", HPBar(p))
	fmt.Println("Level \t\t:", p.Level)
	fmt.Println("Gold \t\t:\033[33m\033[1m", p.Gold, "gold\033[0m")
	fmt.Println("Inventory slots :", p.InventorySlot)
}

func AccessInventory(p Player) {
	if len(p.Inventory) == 0 {
		fmt.Println("\033[31m\033[1mYour inventory is empty.\033[0m")
		return
	}
	fmt.Println("Your inventory contains:")
	for _, item := range p.Inventory {
		fmt.Println("-", item)
	}
}

func HPBar(p Player) string {
	barLength := 15
	if p.HPmax <= 0 {
		p.HPmax = 1
	}
	filled := int(float64(p.HP) / float64(p.HPmax) * float64(barLength))
	if filled < 0 {
		filled = 0
	}
	if filled > barLength {
		filled = barLength
	}

	bar := fmt.Sprintf("[%s%s]", strings.Repeat("█", filled), strings.Repeat(" ", barLength-filled))

	pct := p.HP * 100 / p.HPmax
	color := "\033[31m"
	if pct > 66 {
		color = "\033[32m"
	} else if pct > 33 && pct < 66 {
		color = "\033[33m"
	}
	return color + bar + "\033[0m"
}
