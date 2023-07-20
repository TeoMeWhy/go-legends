package utils

import (
	"math/rand"
	"time"
)

func RollDiceN(n int) int {
	randSrc := rand.NewSource(time.Now().UnixMicro())
	random := rand.New(randSrc)
	return random.Intn(n) + 1
}

func RollsDiceN(N, n int) []int {
	dices := []int{}
	for i := 1; i <= N; i++ {
		dices = append(dices, RollDiceN(n))
	}
	return dices
}
