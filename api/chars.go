package api

import (
	"encoding/json"
	"go_legends/game"
	"net/http"

	"github.com/gorilla/mux"
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

	params := mux.Vars(r)
	nome := params["nome"]

	char := game.LoadChar(nome)
	charAPI := CharGameToAPI(char)

	if charAPI.Raca == "" {
		http.Error(w, "Char não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(charAPI)
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var params CharBaseAPI

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
