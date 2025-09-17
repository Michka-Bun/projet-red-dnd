package dnd

import (
	"fmt"
	"os"
	"sort"
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
			if len(p.Inventory) == 0 {
				lastMsg = "\033[31m\033[1mYour inventory is empty.\033[0m"
				continue
			}
			ClearScreen()
			fmt.Println("Inventory:")
			fmt.Println("----------------")
			AccessInventory(p)
		case 3:
			lastMsg = "\033[33mShop selected.\033[0m"
			AccessShop(p)
		case 4:
			lastMsg = "\033[35mBlacksmith selected.\033[0m"
			Blacksmith(p)
		case 5:
			ClearScreen()
			var surexit int
			fmt.Println("\033[33m\033[1m ⚠ Are you sure to exiting the game.\033[0m")
			fmt.Println("\033[31m\033[1mExit\033[0m ➔  Enter 1")         //\033[31m\033[1m1\033[0m 	//\033[31m\033[1mExit.\033[0m")
			fmt.Println("\033[32m\033[1mResume\033[0m ➔  Any other key") //\033[32m\033[1many other key\033[0m 	//\033[32m\033[1mResume.\033[0m")
			fmt.Scan(&surexit)
			switch surexit {
			case 1:
				ClearScreen()
				fmt.Println("\033[31m\033[1mExiting the game.\033[0m")
				os.Exit(0)

			case 2:
				Menu(p)
			default:
				fmt.Println("\033[37mInvalid choice, please try again.\033[0m")
			}
		default:
		}
	}
}

func AccessInventory(p *Player) {
	lastMsg := ""
	for {
		ClearScreen()
		fmt.Println("Inventory:")
		fmt.Println("----------------")
		if lastMsg != "" {
			fmt.Println(lastMsg)
			fmt.Println()
			lastMsg = ""
		}

		if len(p.Inventory) == 0 {
			fmt.Println("\033[31m\033[1mYour inventory is empty.\033[0m")
			return
		}

		names := make([]string, 0, len(p.Inventory))
		for name, count := range p.Inventory {
			if count <= 0 {
				delete(p.Inventory, name)
				continue
			}
			names = append(names, name)
		}
		if len(names) == 0 {
			fmt.Println("\033[31m\033[1mYour inventory is empty.\033[0m")
			return
		}
		sort.Strings(names)
		for i, name := range names {
			fmt.Printf("%d. %s x%d\n", i+1, name, p.Inventory[name])
		}
		fmt.Println("----------------")
		fmt.Println("1) \033[32mUse/Equip an item\033[0m")
		fmt.Println("2) \033[36mItem info\033[0m")
		fmt.Println("3) \033[35mDrop an item\033[0m")
		fmt.Println("4) \033[31mBack\033[0m")
		fmt.Println("----------------")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var idx int
			fmt.Print("Enter item number to use/equip: ")
			fmt.Scanln(&idx)
			if idx < 1 || idx > len(names) {
				lastMsg = "Invalid item index."
				continue
			}
			name := names[idx-1]
			switch name {
			case "Health potion":
				lastMsg = TakeHealPot(p)
			case "Poison potion":
				lastMsg = TakePoisonPot(p)
			case "Spell book: Fireball":
				if HasSkill(p, "Fireball") {
					lastMsg = "You already know Fireball."
				} else {
					if AddSkill(p, "Fireball") {
						if RemoveItem(p, name, 1) {
							lastMsg = "You learned the skill: \033[34mFireball!\033[0m"
						} else {
							lastMsg = "The spell book seems to be missing."
						}
					} else {
						lastMsg = "Failed to learn Fireball."
					}
				}
			case "Adventurer's hat", "Adventurer's tunic", "Adventurer's boots":
				prevItem, ok := EquipItem(p, name)
				if ok {
					if prevItem != "" {
						lastMsg = fmt.Sprintf("Equipped %s, unequipped %s.", name, prevItem)
					} else {
						lastMsg = fmt.Sprintf("Equipped %s.", name)
					}
				} else {
					lastMsg = "Failed to equip item."
				}
			case "Wolf fur", "Boar leather", "Troll leather", "Crow feather":
				lastMsg = "This item cannot be used or equipped."
			default:
				lastMsg = "This item cannot be used or equipped."
			}
		case 2:
			var idx int
			fmt.Print("Enter item number for info: ")
			fmt.Scanln(&idx)
			if idx < 1 || idx > len(names) {
				lastMsg = "Invalid item index."
				continue
			}
			name := names[idx-1]
			desc := "No description available."
			switch name {
			case "Health potion":
				desc = "Restores 50% of HP when used."
			case "Poison potion":
				desc = "A toxic brew. Does 10 damage each second for 3 seconds to an enemy."
			case "Spell book: Fireball":
				desc = "Learn the Fireball spell for combat."
			case "Adventurer's hat":
				desc = "A sturdy hat for adventurers. Add 10 max HP."
			case "Adventurer's tunic":
				desc = "A durable tunic for adventurers. Add 25 max HP."
			case "Adventurer's boots":
				desc = "Comfortable boots for long journeys. Add 15 max HP."
			case "Wolf fur", "Boar leather", "Troll leather", "Crow feather":
				desc = "\033[35mCrafting material.\033[0m"
			}
			lastMsg = fmt.Sprintf("%s — %s", name, desc)
		case 3:
			var idx int
			fmt.Print("Enter item number to drop: ")
			fmt.Scanln(&idx)
			if idx < 1 || idx > len(names) {
				lastMsg = "Invalid item index."
				continue
			}
			name := names[idx-1]
			if RemoveItem(p, name, 1) {
				lastMsg = fmt.Sprintf("Dropped one %s.", name)
			} else {
				lastMsg = "You don't have that item."
			}
		case 4:
			return
		default:
			lastMsg = "Invalid choice, please try again."
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
		const maxBackpackLevel = 3
		fmt.Println("You have \033[33m\033[1m", p.Gold, "gold\033[0m.")
		fmt.Println("Here is what the shop has to offer:")
		fmt.Println("-----------------------------------")
		if p.Inventory["Health potion"] > 0 {
			fmt.Println("1. Health potion \t\t: \033[33m\033[1m3 \tgold\033[0m")
		} else {
			fmt.Println("1. Health potion \t\t: \033[33m\033[1m \tFREE\033[0m")
		}
		fmt.Println("2. Poison potion \t\t: \033[33m\033[1m6 \tgold\033[0m")
		if p.Inventory["Spell book: Fireball"] == 0 && !HasSkill(p, "Fireball") {
			fmt.Println("3. Spell book \t: \033[34m\033[1mFireball\033[0m \t: \033[33m\033[1m25 \tgold\033[0m")
		}
		fmt.Println("4. Wolf fur \t\t\t: \033[33m\033[1m4 \tgold\033[0m")
		fmt.Println("5. Troll leather \t\t: \033[33m\033[1m7 \tgold\033[0m")
		fmt.Println("6. Boar leather \t\t: \033[33m\033[1m3 \tgold\033[0m")
		fmt.Println("7. Crow feather \t\t: \033[33m\033[1m1 \tgold\033[0m")
		if p.BackpackLevel < maxBackpackLevel {
			fmt.Println("8. Backpack upgrade \t\t: \033[33m\033[1m30 \tgold\033[0m")
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
			if p.Inventory["Health potion"] > 0 {
				if p.Gold >= 3 {
					if AddItem(p, "Health potion", 1) {
						p.Gold -= 3
						lastMsg = "You bought a Health potion."
					} else {
						lastMsg = "Your backpack is full."
					}
				} else {
					lastMsg = "You don't have enough gold."
				}
			} else {
				if AddItem(p, "Health potion", 1) {
					lastMsg = "You received a free Health potion."
				} else {
					lastMsg = "Your backpack is full."
				}
			}
		case 2:
			if p.Gold >= 6 {
				if AddItem(p, "Poison potion", 1) {
					p.Gold -= 6
					lastMsg = "You bought a Poison potion."
				} else {
					lastMsg = "Your backpack is full."
				}
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 3:
			if HasSkill(p, "Fireball") {
				lastMsg = "You already know this spell."
				continue
			}
			if p.Inventory["Spell book: Fireball"] > 0 {
				lastMsg = "You already have this spell book."
				continue
			}
			if p.Gold >= 25 {
				if AddItem(p, "Spell book: Fireball", 1) {
					p.Gold -= 25
					lastMsg = "You bought a Spell book: Fireball."
				} else {
					lastMsg = "Your backpack is full."
				}
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 4:
			if p.Gold >= 4 {
				if AddItem(p, "Wolf fur", 1) {
					p.Gold -= 4
					lastMsg = "You bought a Wolf fur."
				} else {
					lastMsg = "Your backpack is full."
				}
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 5:
			if p.Gold >= 7 {
				if AddItem(p, "Troll leather", 1) {
					p.Gold -= 7
					lastMsg = "You bought a Troll leather."
				} else {
					lastMsg = "Your backpack is full."
				}
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 6:
			if p.Gold >= 3 {
				if AddItem(p, "Boar leather", 1) {
					p.Gold -= 3
					lastMsg = "You bought a Boar leather."
				} else {
					lastMsg = "Your backpack is full."
				}
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 7:
			if p.Gold >= 1 {
				if AddItem(p, "Crow feather", 1) {
					p.Gold -= 1
					lastMsg = "You bought a Crow feather."
				} else {
					lastMsg = "Your backpack is full."
				}
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 8:
			if p.BackpackLevel >= maxBackpackLevel {
				lastMsg = "Your backpack is already at maximum capacity."
			} else if p.Gold >= 30 {
				p.Gold -= 30
				p.BackpackLevel++
				p.InventorySlot += 5
				lastMsg = fmt.Sprintf("Backpack upgraded to level %d (capacity %d).", p.BackpackLevel, p.InventorySlot)
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
				if p.Inventory["Crow feather"] < 1 || p.Inventory["Boar leather"] < 1 {
					lastMsg = "You don't have enough material."
					continue
				}
				if !RemoveItem(p, "Crow feather", 1) || !RemoveItem(p, "Boar leather", 1) {
					AddItem(p, "Crow feather", 1)
					AddItem(p, "Boar leather", 1)
					lastMsg = "You don't have enough material."
					continue
				}
				p.Gold -= 5
				lastMsg = "Crafted Adventurer's hat."
				p.Inventory["Adventurer's hat"]++
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 2:
			if p.Gold >= 5 {
				if p.Inventory["Wolf fur"] < 2 || p.Inventory["Troll leather"] < 1 {
					lastMsg = "You don't have enough material."
					continue
				}
				if !RemoveItem(p, "Wolf fur", 2) || !RemoveItem(p, "Troll leather", 1) {
					AddItem(p, "Wolf fur", 2)
					AddItem(p, "Troll leather", 1)
					lastMsg = "You don't have enough material."
					continue
				}
				p.Gold -= 5
				lastMsg = "Crafted Adventurer's tunic."
				p.Inventory["Adventurer's tunic"]++
			} else {
				lastMsg = "You don't have enough gold."
			}
		case 3:
			if p.Gold >= 5 {
				if p.Inventory["Wolf fur"] < 1 || p.Inventory["Boar leather"] < 1 {
					lastMsg = "You don't have enough material."
					continue
				}
				if !RemoveItem(p, "Wolf fur", 1) || !RemoveItem(p, "Boar leather", 1) {
					AddItem(p, "Wolf fur", 1)
					AddItem(p, "Boar leather", 1)
					lastMsg = "You don't have enough material."
					continue
				}
				p.Gold -= 5
				lastMsg = "Crafted Adventurer's boots."
				p.Inventory["Adventurer's boots"]++
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
