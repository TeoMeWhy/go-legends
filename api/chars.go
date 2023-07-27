package api

import (
	"encoding/json"
	"go_legends/game"
	"net/http"
	"strings"
)

type CharBaseAPI struct {
	Nome   string `json:"nome"`
	Raca   string `json:"raca"`
	Classe string `json:"classe"`
}

type CharAPI struct {
	Nome          string         `json:"nome"`
	Raca          string         `json:"raca"`
	Classe        string         `json:"classe"`
	Nivel         int            `json:"nivel"`
	Vitalidade    int            `json:"vitalidade"`
	Atributos     map[string]int `json:"atributos"`
	Modificadores map[string]int `json:"modificadores"`
}

func CharGameToAPI(c *game.Character) *CharAPI {
	return &CharAPI{
		Nome:          c.Nome,
		Raca:          c.NomeRaca,
		Classe:        c.NomeClasse,
		Nivel:         c.Nivel,
		Vitalidade:    c.Vitalidade,
		Atributos:     c.Atributos,
		Modificadores: c.Modificadores,
	}
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	splitURL := strings.Split(r.URL.Path, "/")
	nome := splitURL[len(splitURL)-1]

	char := game.LoadChar(nome)
	charAPI := CharGameToAPI(char)
	json.NewEncoder(w).Encode(charAPI)

}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {

	var params CharBaseAPI

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido para esta rota", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Deu merda na captura dos parâmetros", http.StatusBadRequest)
		return
	}

	char := game.NewChar(params.Nome, params.Raca, params.Classe)
	game.SaveChar(char)

	charAPI := CharGameToAPI(char)
	json.NewEncoder(w).Encode(charAPI)
}
