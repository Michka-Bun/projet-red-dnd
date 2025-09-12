package UI

import (
	"fmt"
	"os"
)

// main menu function
func Menu() {
	for {
		println("Menu:")
		println("\033[36m1. Character info\033[0m")
		println("\033[32m2. \033[9mInventory\033[0m")
		println("\033[33m3. Shop\033[0m")
		println("\033[35m4. \033[9mBlacksmith\033[0m")
		println("\033[31m5. Exit\033[0m")
		print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			clearScreen()
			// displayInfo()
			println("\033[36mCharacter info selected.\033[0m")
		case 2:
			clearScreen()
			// accessInventory()
			println("\033[32mInventory selected.\033[0m")
		case 3:
			clearScreen()
			println("\033[33mShop selected.\033[0m")
			accessShop()
		case 4:
			clearScreen()
			// Blacksmith()
			println("\033[35mBlacksmith selected.\033[0m")
		case 5:
			clearScreen()
			println("\033[31m\033[1mExiting the game.\033[0m")
			os.Exit(0)
		default:
			clearScreen()
			println("\033[37mInvalid choice, please try again.\033[0m")
		}
	}
}

func accessShop() {
	lastMsg := ""
	for {
		clearScreen()
		if lastMsg != "" {
			fmt.Println(lastMsg)
			fmt.Println()
			lastMsg = ""
		}
		const maxBackpackLevel = 3
		fmt.Println("You have \033[33m\033[1m", Player.gold, "gold\033[0m.")
		fmt.Println("Here is what the shop has to offer:")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Health potion \t\t: \033[33m\033[1m3 \tgold\033[0m")
		fmt.Println("2. Poison potion \t\t: \033[33m\033[1m6 \tgold\033[0m")
		fmt.Println("3. Spell book \t: \033[31m\033[1mFireball\033[0m \t: \033[33m\033[1m25 \tgold\033[0m")
		fmt.Println("4. Wolf fur \t\t\t: \033[33m\033[1m4 \tgold\033[0m")
		fmt.Println("5. Troll leather \t\t: \033[33m\033[1m7 \tgold\033[0m")
		fmt.Println("6. Boar leather \t\t: \033[33m\033[1m3 \tgold\033[0m")
		fmt.Println("7. Crow feather \t\t: \033[33m\033[1m1 \tgold\033[0m")
		if Player.BackpackLevel < maxBackpackLevel {
			fmt.Println("8. Backpack upgrade \t\t: \033[33m30\033[1m \tgold\033[0m")
		} else {
			fmt.Println("8. \033[9mBackpack upgrade\033[0m \t\t: \033[31m\033[1mMAX CAPACITY REACHED\033[0m")
		}
		fmt.Println("-----------------------------------")
		fmt.Println("9. Exit shop")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if Player.gold >= 3 {
				Player.gold -= 3
				lastMsg = "You bought a Health potion."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 2:
			if Player.gold >= 6 {
				Player.gold -= 6
				lastMsg = "You bought a Poison potion."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 3:
			if Player.gold >= 25 {
				Player.gold -= 25
				lastMsg = "You bought a Spell book : Fireball."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 4:
			if Player.gold >= 4 {
				Player.gold -= 4
				lastMsg = "You bought a Wolf fur."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 5:
			if Player.gold >= 7 {
				Player.gold -= 7
				lastMsg = "You bought a Troll leather."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 6:
			if Player.gold >= 3 {
				Player.gold -= 3
				lastMsg = "You bought a Boar leather."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 7:
			if Player.gold >= 1 {
				Player.gold -= 1
				lastMsg = "You bought a Crow feather."
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 8:
			if Player.BackpackLevel >= maxBackpackLevel {
				lastMsg = "Your backpack is already at maximum capacity."
				break
			}
			if Player.gold >= 30 {
				Player.gold -= 30
				lastMsg = "You bought a Backpack upgrade."
				Player.BackpackLevel++

			} else {
				lastMsg = "You don't have enough gold."
			}
		case 9:
			clearScreen()
			fmt.Println("Exiting the shop.")
			return
		default:
			lastMsg = "Invalid choice, please try again."
		}
	}
}

func blacksmith() {
	lastMsg := ""
	for {
		clearScreen()
		if lastMsg != "" {
			fmt.Println(lastMsg)
			fmt.Println()
			lastMsg = ""
		}
		fmt.Println("You have :\n- \033[33m\033[1m", Player.gold, "gold\033[0m\n- \033[35m\033[1m", Player.CrowFeather, "Crow feather(s)\033[0m\n- \033[35m\033[1m", Player.BoarLeather, "Boar leather(s)\033[0m\n- \033[35m\033[1m", Player.WolfFur, "Wolf fur(s)\033[0m\n- \033[35m\033[1m", Player.TrollLeather, "Troll leather(s)\033[0m")
		fmt.Println("Welcome to the Blacksmith!")
		fmt.Println("Here is what the Blacksmith has to offer:")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Adventurer's hat \t\t: \033[33m\033[1m5 \tgold\033[0m | \033[35m1 \tCrow feather\033[0m | \033[35m1 \tBoar leather\033[0m |")
		fmt.Println("2. Adventurer's tunic \t\t: \033[33m\033[1m5 \tgold\033[0m | \033[35m2 \tWolf fur\033[0m | \033[35m1 \tTroll leather\033[0m |")
		fmt.Println("3. Adventurer's boots \t\t: \033[33m\033[1m5 \tgold\033[0m | \033[35m1 \tWolf fur\033[0m | \033[35m1 \tBoar leather\033[0m |")
		fmt.Println("-----------------------------------")
		fmt.Println("4. Exit Blacksmith")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if Player.gold >= 5 && Player.CrowFeather >= 1 && Player.BoarLeather >= 1 {
				Player.gold -= 5
				Player.CrowFeather -= 1
				Player.BoarLeather -= 1
				lastMsg = "You bought an Adventurer's hat."
			} else {
				lastMsg = "You don't have enough gold or materials."
			}
		case 2:
			if Player.gold >= 5 && Player.WolfFur >= 2 && Player.TrollLeather >= 1 {
				Player.gold -= 5
				Player.WolfFur -= 2
				Player.TrollLeather -= 1
				lastMsg = "You bought an Adventurer's tunic."
			} else {
				lastMsg = "You don't have enough gold or materials."
			}
		case 3:
			if Player.gold >= 5 && Player.WolfFur >= 1 && Player.BoarLeather >= 1 {
				Player.gold -= 5
				Player.WolfFur -= 1
				Player.BoarLeather -= 1
				lastMsg = "You bought an Adventurer's boots."
			} else {
				lastMsg = "You don't have enough gold or materials."
			}
		case 4:
			clearScreen()
			fmt.Println("Exiting the Blacksmith.")
			return
		default:
			lastMsg = "Invalid choice, please try again."
		}
	}
}

func fightMenu() {
	isFighting = true
	clearScreen()
	// displayEnemyInfo()
	// displayPlayerInfo()

	// fight options
	println("Fight Menu:")
	println("-------------------")
	println("1. \033[31mAttack\033[0m")
	println("2. \033[32mInventory\033[0m")
	print("Choose an option: ")

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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
