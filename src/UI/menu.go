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
			println("\033[31mExiting the game. Goodbye!\033[0m")
			os.Exit(0)
		default:
			clearScreen()
			println("\033[37mInvalid choice, please try again.\033[0m")
		}
	}
}

func accessShop() {
	for {
		const maxBackpackLevel = 3
		fmt.Println("You have \033[33m", Player.gold, "gold\033[0m.")
		fmt.Println("Here is what the shop has to offer:")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Health potion \t\t: \033[33m3 \tgold\033[0m")
		fmt.Println("2. Poison potion \t\t: \033[33m6 \tgold\033[0m")
		fmt.Println("3. Spell book \t: \033[31m\033[1mFireball\033[0m \t: \033[33m25 \tgold\033[0m")
		fmt.Println("4. Wolf fur \t\t\t: \033[33m4 \tgold\033[0m")
		fmt.Println("5. Troll leather \t\t: \033[33m7 \tgold\033[0m")
		fmt.Println("6. Boar leather \t\t: \033[33m3 \tgold\033[0m")
		fmt.Println("7. Crow feather \t\t: \033[33m1 \tgold\033[0m")
		if Player.BackpackLevel < maxBackpackLevel {
			fmt.Println("8. Backpack upgrade \t\t: \033[33m30 \tgold\033[0m")
		} else {
			fmt.Println("8. \033[9mBackpack upgrade\033[0m \t\t: \033[31mMAX CAPACITY REACHED\033[0m")
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
				fmt.Println("You bought a Health potion.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 2:
			if Player.gold >= 6 {
				Player.gold -= 6
				fmt.Println("You bought a Poison potion.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 3:
			if Player.gold >= 25 {
				Player.gold -= 25
				fmt.Println("You bought a Spell book : Fireball.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 4:
			if Player.gold >= 4 {
				Player.gold -= 4
				fmt.Println("You bought a Wolf fur.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 5:
			if Player.gold >= 7 {
				Player.gold -= 7
				fmt.Println("You bought a Troll leather.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 6:
			if Player.gold >= 3 {
				Player.gold -= 3
				fmt.Println("You bought a Boar leather.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 7:
			if Player.gold >= 1 {
				Player.gold -= 1
				fmt.Println("You bought a Crow feather.")
			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 8:
			if Player.BackpackLevel >= maxBackpackLevel {
				fmt.Println("Your backpack is already at maximum capacity.")
				break
			}
			if Player.gold >= 30 {
				Player.gold -= 30
				fmt.Println("You bought a Backpack upgrade.")
				Player.BackpackLevel++

			} else {
				fmt.Println("You don't have enough gold.")
			}
		case 9:
			clearScreen()
			fmt.Println("Exiting the shop.")
			return
		default:
			clearScreen()
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
