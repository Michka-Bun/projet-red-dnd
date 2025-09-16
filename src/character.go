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
	BaseDamage    int
	Level         int
	Gold          int
	BackpackLevel int
	InventorySlot int
	Inventory     map[string]int
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
	var HPmax, HP, Level, BaseDamage int
	var Gold int
	var InventorySlot int
	var Inventory map[string]int
	if Class == "Warrior" {
		HPmax = 100
		HP = 100
		BaseDamage = 15
		InventorySlot = 5
		Inventory = map[string]int{}
	} else if Class == "Mage" {
		HPmax = 75
		HP = 75
		BaseDamage = 10
		InventorySlot = 10
		Inventory = map[string]int{}
	}
	Level = 1
	Gold = 100
	BackpackLevel := 1
	//Autres infos
	var PoisonEffect int

	//retour d'infos
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Println("Informations on your new character:")
	fmt.Println("Name \t\t:", Name)
	fmt.Println("Class \t\t:", Class)
	fmt.Println("HP \t\t:", HP, "/", HPmax)
	fmt.Println("Base damage \t\t:", BaseDamage)
	fmt.Println("Level \t\t:", Level)
	fmt.Println("Gold \t\t:", Gold)
	fmt.Println("Backpack :",
		fmt.Sprintf("L%d (%d/%d)", BackpackLevel, countItems(Inventory), InventorySlot))
	fmt.Println("Inventory items :", formatInventory(Inventory))

	return Player{
		Name:          Name,
		Class:         Class,
		HPmax:         HPmax,
		HP:            HP,
		BaseDamage:    BaseDamage,
		Level:         Level,
		Gold:          Gold,
		BackpackLevel: BackpackLevel,
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
	fmt.Println("Base damage \t:", p.BaseDamage)
	fmt.Println("Level \t\t:\033[36m\033[1m", p.Level, "\033[0m")   //bleu clair
	fmt.Println("Gold \t\t:\033[33m\033[1m", p.Gold, "gold\033[0m") //jaune
	fmt.Println("Backpack \t:",
		fmt.Sprintf("\033[36m\033[1mL%d \033[0m\t\t(%d/%d)\033[0m", p.BackpackLevel, countItems(p.Inventory), p.InventorySlot))
	fmt.Printf("\nHP \t\t: \033[1m%s%d\033[0m\t\t %s\n", HPColor(p), p.HP, HPBar(p))
	fmt.Println("Poison effect \t:", p.PoisonEffect, "\t turn(s) left")
}

func AddItem(p *Player, name string, count int) bool {
	if name == "" || count <= 0 {
		return false
	}
	if p.Inventory == nil {
		p.Inventory = map[string]int{}
	}
	if countItems(p.Inventory)+count > p.InventorySlot {
		return false
	}
	p.Inventory[name] += count
	return true
}

func formatInventory(inv map[string]int) string {
	if len(inv) == 0 {
		return "[]"
	}
	s := "["
	first := true
	for name, count := range inv {
		if !first {
			s += ", "
		}
		s += fmt.Sprintf("%s x%d", name, count)
		first = false
	}
	s += "]"
	return s
}

// countItems returns the total number of item units in the inventory.
func countItems(inv map[string]int) int {
	total := 0
	for _, c := range inv {
		total += c
	}
	return total
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
	color := "\033[31m" //rouge
	if pct > 66 {
		color = "\033[32m" //vert
	} else if pct > 33 && pct <= 66 {
		color = "\033[33m" //jaune
	}
	return color + bar + "\033[0m"
}

// hpColor returns the ANSI color code for the current HP percentage.
func HPColor(p Player) string {
	if p.HPmax <= 0 {
		return "\033[31m" //rouge
	}
	pct := p.HP * 100 / p.HPmax
	if pct > 66 {
		return "\033[32m" //vert
	} else if pct > 33 && pct <= 66 {
		return "\033[33m" //jaune
	}
	return "\033[31m" //rouge
}
