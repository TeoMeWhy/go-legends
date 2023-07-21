package main

import (
	"fmt"
	"go_legends/character"
)

func main() {

	// teo := character.NewChar("Téo", "anao", "guerreiro")

	// fmt.Println(teo)
	// teo.AddXP(7000)
	// fmt.Println(teo)

	// character.SaveChar(teo)

	teo := character.LoadChar("Téo")
	fmt.Println(teo)
}
