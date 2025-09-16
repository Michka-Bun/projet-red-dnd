package main

import (
	"dnd"
	"fmt"
)

func main() {
	dnd.ClearScreen()
	fmt.Println("Welcome to the Game!")
	p := dnd.SetInfo()
	dnd.ClearScreen()
	fmt.Println("-------------------")
	dnd.Menu(&p)

	/* // Test affichage stats monstre
	m := dnd.ChooseMonster()
	dnd.DisplayMonsterInfo(m)
	*/
}
