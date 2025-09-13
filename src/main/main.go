package main

import (
	"fmt"
	projet_red_dnd "projet-red-dnd/character"
)

func main() {
	fmt.Println("Welcome to the RPG Game!")
	/*	for {
			UI.Menu()
		}
	*/
	p := projet_red_dnd.SetInfo()
	fmt.Println("============")
	projet_red_dnd.DisplayInfo(p)
}
