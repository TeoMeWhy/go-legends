package main

import (
	"fmt"
	"go_legends/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/chars/{nome}", api.GetCharacter).Methods("GET")
	r.HandleFunc("/create_char", api.CreateCharacter).Methods("POST")
	r.HandleFunc("/create_batalha", api.CreateBatalha).Methods("POST")
	r.HandleFunc("/batalha", api.GetBatalha).Methods("GET")
	r.HandleFunc("/pre_batalha", api.StartPreBatalha).Methods("POST")
	r.HandleFunc("/join_batalha/{nome}", api.JoinPreBatalha).Methods("POST")

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", r)

}
