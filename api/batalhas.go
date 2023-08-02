package api

import (
	"encoding/json"
	"fmt"
	"go_legends/game"
	"net/http"

	"github.com/gorilla/mux"
)

type inimigosAPI struct {
	Nome   string `json:"nome"`
	Raca   string `json:"raca"`
	Classe string `json:"classe"`
}

var batalha game.Batalha
var chPreBatalha = make(chan string)

func CreateBatalha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inimigos []inimigosAPI

	err := json.NewDecoder(r.Body).Decode(&inimigos)
	if err != nil {
		fmt.Println("Não foi possível obter os inimigos")
		fmt.Println(err)
		return
	}

	inimigosGame := []game.Character{}
	for _, i := range inimigos {
		inimigosGame = append(inimigosGame, *game.NewChar(i.Nome, i.Raca, i.Classe))
	}

	batalha.Inimigos = inimigosGame
	batalha.Aliados = []game.Character{}
}

func GetBatalha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(batalha)
	if err != nil {
		fmt.Println("Não foi possível obter a batalha")
		fmt.Println(err)
	}

}

func StartPreBatalha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go batalha.PreBatalha(&chPreBatalha)
}

func JoinPreBatalha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	nome := params["nome"]
	chPreBatalha <- nome
}
