package dnd

import (
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name               string
	Class              string
	HPmax              int
	HP                 int
	BaseDamage         int
	Level              int
	XP                 int
	XPmax              int
	Gold               int
	BackpackLevel      int
	InventorySlot      int
	Inventory          map[string]int
	Skills             map[string]bool
	PoisonEffect       int
	WeakeningTrunCount int
	UseHealPot         bool
	Head               string
	Body               string
	Feet               string
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
	for !(Class == "Warrior" || Class == "Mage" || Class == "Viking" || Class == "Archer") {
		fmt.Println("Choose your character's class:")
		fmt.Println("1 - Warrior") //faible contre piaf -25%
		fmt.Println("2 - Mage")    //faible contre devourer
		fmt.Println("3 - Viking")  //faible contre piaf
		fmt.Println("4 - Archer")  //faible contre devoureur
		fmt.Scan(&inputClass)
		fmt.Println(strings.Repeat("\n", 21))

		if inputClass == "1" {
			Class = "Warrior"
		} else if inputClass == "2" {
			Class = "Mage"
		} else if inputClass == "3" {
			Class = "Viking"
		} else if inputClass == "4" {
			Class = "Archer"
		} else {
			fmt.Println("// Invalid Class // x", ComptInvalidClass)
			ComptInvalidClass++
		}
	}

	//stats en fonction de la classe
	var HPmax, HP, Level, XP, XPmax, BaseDamage int
	var Gold int
	var InventorySlot int
	var Inventory map[string]int
	var Skills map[string]bool
	if Class == "Warrior" {
		HPmax = 100
		HP = 100
		BaseDamage = 15
		InventorySlot = 5
	} else if Class == "Mage" {
		HPmax = 75
		HP = 75
		BaseDamage = 10
		InventorySlot = 10
	} else if Class == "Viking" {
		HPmax = 125
		HP = 125
		BaseDamage = 20
		InventorySlot = 3
	} else if Class == "Archer" {
		HPmax = 100
		HP = 100
		BaseDamage = 15
		InventorySlot = 5
	}
	Level = 1
	XP = 0
	XPmax = 100
	Gold = 100
	BackpackLevel := 1
	Inventory = map[string]int{}
	Skills = map[string]bool{"Punch": true}
	var PoisonEffect int
	var WeakeningTrunCount int
	var UseHealPot = false

	//retour d'infos
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Println("Informations on your new character:")
	fmt.Println("Name \t\t:", Name)
	fmt.Println("Class \t\t:", Class)
	fmt.Println("HP \t\t:", HP, "/", HPmax)
	fmt.Println("Base damage \t:", BaseDamage)
	fmt.Println("Level \t\t:", Level)
	fmt.Println("Gold \t\t:", Gold)
	fmt.Println("Backpack :",
		fmt.Sprintf("L%d (%d/%d)", BackpackLevel, CountItems(Inventory), InventorySlot))
	fmt.Println("Inventory items :", FormatInventory(Inventory))

	return Player{
		Name:               Name,
		Class:              Class,
		HPmax:              HPmax,
		HP:                 HP,
		BaseDamage:         BaseDamage,
		Level:              Level,
		XP:                 XP,
		XPmax:              XPmax,
		Gold:               Gold,
		BackpackLevel:      BackpackLevel,
		InventorySlot:      InventorySlot,
		Inventory:          Inventory,
		Skills:             Skills,
		PoisonEffect:       PoisonEffect,
		UseHealPot:         UseHealPot,
		WeakeningTrunCount: WeakeningTrunCount,
	}
}

func DisplayInfo(p Player) {
	fmt.Println("Informations on your character:")
	fmt.Println("------------------------------")
	fmt.Println("Name \t\t:\033[97m", p.Name, "\033[0m")
	fmt.Println("Class \t\t:\033[97m", p.Class, "\033[0m")
	fmt.Println("HP max \t\t:\033[97m", p.HPmax, "\033[0m")
	fmt.Println("Base damage \t:\033[97m", p.BaseDamage, "\033[0m")
	fmt.Println("Level \t\t:\033[36m\033[1m", p.Level, "\033[0m with\033[36m\033[1m", p.XP, "/", p.XPmax, "\033[0mXP")
	fmt.Println("Gold \t\t:\033[33m\033[1m", p.Gold, "gold(s)\033[0m")
	fmt.Println("Backpack \t:",
		fmt.Sprintf("\033[36m\033[1mL%d \033[0m\t\t(%d/%d)\033[0m", p.BackpackLevel, CountItems(p.Inventory), p.InventorySlot))
	fmt.Println("Skills \t\t:\033[34m", FormatSkills(p.Skills), "\033[0m")
	fmt.Printf("\nHP \t\t: \033[1m%s%d\033[0m\t\t %s\n", hpColor(p), p.HP, HPBar(p))
	fmt.Println("\nEquipment \t: \tHead:\t\033[35m", p.Head, "\033[0m,\n \t\t\tBody:\t\033[35m", p.Body, "\033[0m,\n \t\t\tFeet:\t\033[35m", p.Feet, "\033[0m,")
	fmt.Println("Poison effect \t:", p.PoisonEffect, "\t turn(s) left")
}

func LevelUp(p *Player) {
	p.Level++
	fmt.Println("Congratulations, you have \033[36m\033[1m leveled up \033[0m !")
	fmt.Println("Your sats have increassed by \033[36m\033[1m 10% \033[0m")
	p.HPmax += p.HPmax * 10 / 100
	p.HP = p.HPmax
	p.BaseDamage += p.BaseDamage * 10 / 100
	if p.XP == p.XPmax {
		p.XP = 0
	} else if p.XP > p.XPmax {
		p.XP -= p.XPmax
	}
	p.XPmax += p.XPmax * 15 / 100
	p.Gold += 30
	if p.Level == 100 {
		ClearScreen()
		fmt.Println("\033[34m\033[1mGo to sleep !\033[0m")
		os.Exit(0)
	}
}

func AddItem(p *Player, name string, count int) bool {
	if name == "" || count <= 0 {
		return false
	}
	if p.Inventory == nil {
		p.Inventory = map[string]int{}
	}
	if CountItems(p.Inventory)+count > p.InventorySlot {
		return false
	}
	p.Inventory[name] += count
	return true
}

func RemoveItem(p *Player, name string, count int) bool {
	if p == nil || name == "" || count <= 0 {
		return false
	}
	if p.Inventory == nil {
		return false
	}
	current, ok := p.Inventory[name]
	if !ok || current < count {
		return false
	}
	current -= count
	if current == 0 {
		delete(p.Inventory, name)
	} else {
		p.Inventory[name] = current
	}
	return true
}

func FormatInventory(inv map[string]int) string {
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

func FormatSkills(sk map[string]bool) string {
	if len(sk) == 0 {
		return "[]"
	}
	s := "["
	first := true
	for name := range sk {
		if !first {
			s += ", "
		}
		s += name
		first = false
	}
	s += "]"
	return s
}

func CountItems(inv map[string]int) int {
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

func hpColor(p Player) string {
	pct := p.HP * 100 / p.HPmax
	if pct > 66 {
		return "\033[32m\033[1m" //vert
	} else if pct > 33 && pct <= 66 {
		return "\033[33m\033[1m" //jaune
	} else {
		return "\033[31m\033[1m" //rouge
	}
}

func IsDead(p *Player) {
	var Revive int
	ClearScreen()
	fmt.Println("\033[31m\033[1m Game Over !\033[0m")
	fmt.Println()
	fmt.Println("\033[31m\033[1m Give up\033[0m \t➔  Enter 1")
	fmt.Println("\033[32m\033[1m Revive \033[0m \t➔  Enter any other key")
	fmt.Println("\033[38;5;208m ⚠ Reviving will delete your inventory and reduce your \033[33m\033[1mgolds\033[38;5;208m by 75%\033[0m")
	fmt.Println("\033[38;5;208m ⚠ Your new balance will be :\033[1m", p.Gold-p.Gold*75/100, "\033[33m\033[1mgold(s)\033[0m")
	fmt.Println("\033[32m You will keep all the rest\033[0m")
	fmt.Scan(&Revive)
	if Revive == 1 {
		ClearScreen()
		var suregiveup int
		fmt.Println("\033[38;5;820m ⚠ Are you sure you are giving up ?\033[0m")
		fmt.Println("\033[38;5;208m ⚠ Your progress will be deleted !\033[0m")
		fmt.Println("\033[31m\033[1m Give up\033[0m \t➔  Enter 1")
		fmt.Println("\033[32m\033[1m Back\033[0m \t\t➔  Enter any other key")
		fmt.Scan(&suregiveup)
		if suregiveup == 1 {
			ClearScreen()
			fmt.Println("\033[31m\033[1m Giving up the game.\033[0m")
			os.Exit(0)
		} else {
			IsDead(p)
		}
	} else {
		p.Gold -= p.Gold * 75 / 100
		p.HP = p.HPmax / 2
		p.Inventory = map[string]int{}
		if p.Head != "" {
			switch p.Head {
			case "Adventurer's hat":
				p.HPmax -= 10
			}
		}
		if p.Body != "" {
			switch p.Body {
			case "Adventurer's tunic":
				p.HPmax -= 25
			}
		}
		if p.Feet != "" {
			switch p.Feet {
			case "Adventurer's boots":
				p.HPmax -= 15
			}
		}
		p.PoisonEffect = 0
		Menu(p)
		fmt.Println("\033[33m\033[1m You were resurrected.\033[0m")
	}
}

type EquipmentStats struct {
	Slot    string
	HPBonus int
}

var equipmentData = map[string]EquipmentStats{
	"Adventurer's hat":   {Slot: "Head", HPBonus: 10},
	"Adventurer's tunic": {Slot: "Body", HPBonus: 25},
	"Adventurer's boots": {Slot: "Feet", HPBonus: 15},
}

func EquipItem(p *Player, item string) (string, bool) {
	data, ok := equipmentData[item]
	if !ok || p.Inventory[item] <= 0 {
		return "", false
	}
	prevItem := ""
	switch data.Slot {
	case "Head":
		prevItem = p.Head
		p.Head = item
	case "Body":
		prevItem = p.Body
		p.Body = item
	case "Feet":
		prevItem = p.Feet
		p.Feet = item
	}
	if prevItem != "" {
		removeEquipmentStats(p, prevItem)
		p.Inventory[prevItem]++
	}
	RemoveItem(p, item, 1)
	applyEquipmentStats(p, item)
	return prevItem, true
}

func UnequipItem(p *Player, slot string) bool {
	var item string
	switch slot {
	case "Head":
		item = p.Head
		p.Head = ""
	case "Body":
		item = p.Body
		p.Body = ""
	case "Feet":
		item = p.Feet
		p.Feet = ""
	default:
		return false
	}
	if item == "" {
		return false
	}
	removeEquipmentStats(p, item)
	p.Inventory[item]++
	return true
}

func applyEquipmentStats(p *Player, item string) {
	if data, ok := equipmentData[item]; ok {
		p.HPmax += data.HPBonus
		if p.HP > p.HPmax {
			p.HP = p.HPmax
		}
	}
}

func removeEquipmentStats(p *Player, item string) {
	if data, ok := equipmentData[item]; ok {
		p.HPmax -= data.HPBonus
		if p.HP > p.HPmax {
			p.HP = p.HPmax
		}
	}
}
