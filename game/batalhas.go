package game

import (
	"fmt"
	"time"
)

type Batalha struct {
	Inimigos []Character
	Aliados  []Character
}

func (b *Batalha) PreBatalha(c *chan string) {
	fmt.Println("Iniciando pré batalha!")
	for {
		select {
		case nome := <-*c:
			fmt.Println(nome)
			char := LoadChar(nome)
			b.Aliados = append(b.Aliados, *char)
		case <-time.After(time.Second * 10):
			fmt.Println("Pre batalha encerrada!")
			fmt.Println(b.Aliados)
			return
		}
	}

}

func CriaBatalha(inimigos []Character) *Batalha {
	return &Batalha{Inimigos: inimigos}
}
