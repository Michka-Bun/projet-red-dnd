package main

import (
    "fmt"
    "dnd"
)

func main() {
	dnd.ClearScreen()
	fmt.Println("Welcome to the Game!")
	p := dnd.SetInfo()
	dnd.ClearScreen()
	fmt.Println("-------------------")
	dnd.Menu(p)
}
