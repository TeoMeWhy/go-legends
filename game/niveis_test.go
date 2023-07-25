package game

import "testing"

type testpair struct {
	exp    int
	nivel  int
	pontos int
}

var testsGetnivel = []testpair{
	{0, 1, 0},
	{400, 2, 1},
	{1000, 3, 1},
	{309000, 19, 5},
	{1000000, 20, 5},
}

func TestGetNivel(t *testing.T) {
	for _, pair := range testsGetnivel {
		nivel := GetNivel(pair.exp)
		if nivel != pair.nivel {
			t.Error("Esperado:", pair.exp,
				"Obtido:", nivel,
			)
		}
	}
}
