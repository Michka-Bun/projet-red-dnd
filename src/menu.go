package dnd

import (
	"fmt"
	"os"
)

// main menu function
func Menu(p *Player) {
	lastMsg := ""
	for {
		ClearScreen()
		if lastMsg != "" {
			fmt.Println(lastMsg)
			fmt.Println()
			lastMsg = ""
		}
		fmt.Println("-------------------")
		fmt.Println("Menu:")
		fmt.Println("\033[36m1. Character info\033[0m")
		fmt.Println("\033[32m2. Inventory\033[0m")
		fmt.Println("\033[33m3. Shop\033[0m")
		fmt.Println("\033[35m4. Blacksmith\033[0m")
		fmt.Println("\033[31m5. Exit\033[0m")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ClearScreen()
			DisplayInfo(*p)
			fmt.Println()
			fmt.Print("Press Enter to return to menu...")
			var _pause byte
			fmt.Scanf("%c", &_pause)
			lastMsg = "\033[36mCharacter info displayed.\033[0m"
		case 2:
			ClearScreen()
			fmt.Println("Inventory:")
			fmt.Println("----------------")
			AccessInventory(*p)
			fmt.Println()
			fmt.Print("Press Enter to return to menu...")
			var _pause2 byte
			fmt.Scanf("%c", &_pause2)
			lastMsg = "\033[32mInventory displayed.\033[0m"
		case 3:
			lastMsg = "\033[33mShop selected.\033[0m"
			AccessShop(p)
		case 4:
			lastMsg = "\033[35mBlacksmith selected.\033[0m"
			Blacksmith(p)
		case 5:
			ClearScreen()
			fmt.Println("\033[31m\033[1mExiting the game.\033[0m")
			os.Exit(0)
		default:
			lastMsg = "\033[37mInvalid choice, please try again.\033[0m"
		}
	}
}

func AccessShop(p *Player) {
	lastMsg := ""
	for {
		ClearScreen()
		if lastMsg != "" {
			fmt.Println(lastMsg)
			fmt.Println()
			lastMsg = ""
		}
		// const maxBackpackLevel = 3
		fmt.Println("You have \033[33m\033[1m", p.Gold, "gold\033[0m.")
		fmt.Println("Here is what the shop has to offer:")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Health potion \t\t: \033[33m\033[1m3 \tgold\033[0m")
		fmt.Println("2. Poison potion \t\t: \033[33m\033[1m6 \tgold\033[0m")
		fmt.Println("3. Spell book \t: \033[31m\033[1mFireball\033[0m \t: \033[33m\033[1m25 \tgold\033[0m")
		fmt.Println("4. Wolf fur \t\t\t: \033[33m\033[1m4 \tgold\033[0m")
		fmt.Println("5. Troll leather \t\t: \033[33m\033[1m7 \tgold\033[0m")
		fmt.Println("6. Boar leather \t\t: \033[33m\033[1m3 \tgold\033[0m")
		fmt.Println("7. Crow feather \t\t: \033[33m\033[1m1 \tgold\033[0m")
		// if Player.BackpackLevel < maxBackpackLevel {
		//     fmt.Println("8. Backpack upgrade \t\t: \033[33m30\033[1m \tgold\033[0m")
		// } else {
		//     fmt.Println("8. \033[9mBackpack upgrade\033[0m \t\t: \033[31m\033[1mMAX CAPACITY REACHED\033[0m")
		// }
		fmt.Println("8. Backpack upgrade \t\t: \033[33m\033[1m30 \tgold\033[0m")
		fmt.Println("-----------------------------------")
		fmt.Println("9. Exit shop")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if p.Gold >= 3 {
				p.Gold -= 3
				lastMsg = "You bought a Health potion."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 2:
			if p.Gold >= 6 {
				p.Gold -= 6
				lastMsg = "You bought a Poison potion."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 3:
			if p.Gold >= 25 {
				p.Gold -= 25
				lastMsg = "You bought a Spell book: Fireball."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 4:
			if p.Gold >= 4 {
				p.Gold -= 4
				lastMsg = "You bought a Wolf fur."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 5:
			if p.Gold >= 7 {
				p.Gold -= 7
				lastMsg = "You bought a Troll leather."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 6:
			if p.Gold >= 3 {
				p.Gold -= 3
				lastMsg = "You bought a Boar leather."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 7:
			if p.Gold >= 1 {
				p.Gold -= 1
				lastMsg = "You bought a Crow feather."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 8:
			if p.Gold >= 30 {
				p.Gold -= 30
				lastMsg = "You bought a Backpack upgrade."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 9:
			ClearScreen()
			fmt.Println("Exiting the shop.")
			return
		default:
			lastMsg = "Invalid choice, please try again."
		}
	}
}

func Blacksmith(p *Player) {
	lastMsg := ""
	for {
		ClearScreen()
		if lastMsg != "" {
			fmt.Println(lastMsg)
			fmt.Println()
			lastMsg = ""
		}
		fmt.Println("You have :\n- \033[33m\033[1m", p.Gold, "gold\033[0m")
		fmt.Println("Welcome to the Blacksmith!")
		fmt.Println("Here is what the Blacksmith has to offer:")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Adventurer's hat \t\t: \033[33m\033[1m5 \tgold\033[0m | \033[35m1 Crow feather\033[0m \t| \033[35m1 Boar leather\033[0m \t|")
		fmt.Println("2. Adventurer's tunic \t\t: \033[33m\033[1m5 \tgold\033[0m | \033[35m2 Wolf fur\033[0m \t| \033[35m1 Troll leather\033[0m \t|")
		fmt.Println("3. Adventurer's boots \t\t: \033[33m\033[1m5 \tgold\033[0m | \033[35m1 Wolf fur\033[0m \t| \033[35m1 Boar leather\033[0m \t|")
		fmt.Println("-----------------------------------")
		fmt.Println("4. Exit Blacksmith")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if p.Gold >= 5 {
				p.Gold -= 5
				lastMsg = "Crafted Adventurer's hat (gold only)."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 2:
			if p.Gold >= 5 {
				p.Gold -= 5
				lastMsg = "Crafted Adventurer's tunic (gold only)."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 3:
			if p.Gold >= 5 {
				p.Gold -= 5
				lastMsg = "Crafted Adventurer's boots (gold only)."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 4:
			ClearScreen()
			fmt.Println("Exiting the Blacksmith.")
			return
		default:
			lastMsg = "Invalid choice, please try again."
		}
	}
}

func FightMenu() {
	// isFighting := true
	ClearScreen()
	// displayEnemyInfo()
	// displayPlayerInfo()

	fmt.Println("Fight Menu:")
	fmt.Println("-------------------")
	fmt.Println("1. \033[31mAttack\033[0m")
	fmt.Println("2. \033[32mInventory\033[0m")
	fmt.Println("Choose an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		//attack()
	case 2:
		// accessInventory()
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}

func ClearScreen() {
	fmt.Print("\033[2J\033[3J\033[H")
}
