package main

import (
	"fmt"
	"go_legends/character"
)

func main() {

	// teo := character.NewChar("Téo", "humano", "guerreiro")
	// fmt.Println(teo)

	// character.SaveChar(teo)

	teo := character.LoadChar("Téo")
	fmt.Println(teo)
}
