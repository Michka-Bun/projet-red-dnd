package main

import (
	"fmt"
	"projet-red-dnd/UI"
	projet_red_dnd "projet-red-dnd/character"
)

func main() {
	fmt.Println("Welcome to the RPG Game!")
	for {
		UI.Menu()
	}
	projet_red_dnd.SetInfo()
	projet_red_dnd.DisplayInfo()
}
