package main

import (
	"fmt"
	"go_legends/game"
)

func main() {

	teo := game.NewChar("Téo", "elfo", "arqueiro")
	fmt.Println(teo)

	teo.Items["armadura"] = *game.LoadItem("Armadura de Couro")

	fmt.Println(teo)
	game.SaveChar(teo)

	teo = game.LoadChar("Téo")
	fmt.Println(teo)

}
