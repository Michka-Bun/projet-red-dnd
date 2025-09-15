package main

import (
	"fmt"
	UI "projet-red-dnd/UI"
	character "projet-red-dnd/character"
)

func main() {
	UI.ClearScreen()
	fmt.Println("Welcome to the Game!")
	p := character.SetInfo()
	UI.ClearScreen()
	fmt.Println("-------------------")
	UI.Menu(p)
}
