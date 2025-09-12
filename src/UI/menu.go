package UI

import (
	"fmt"
	"os"
)

// main menu function
func Menu() {
	println("Menu:")
	println("\033[32m1. Character info\033[0m")
	println("\033[32m2. Inventory\033[0m")
	println("\033[32m3. Shop\033[0m")
	println("\033[32m4. Blacksmith\033[0m")
	println("\033[31m5. Exit\033[0m")
	print("Choose an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// displayInfo()
		println("Character info selected.")
	case 2:
		// accessInventory()
		println("Inventory selected.")
	case 3:
		// accessShop()
		println("Shop selected.")
	case 4:
		// Blacksmith()
		println("Blacksmith selected.")
	case 5:
		println("Exiting the game. Goodbye!")
		os.Exit(0)
	default:
		println("Invalid choice, please try again.")
	}
}
