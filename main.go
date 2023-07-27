package main

import (
	"fmt"
	"go_legends/api"
	"net/http"
)

func main() {

	http.HandleFunc("/chars/", api.GetCharacter)

	http.HandleFunc("/create_char", api.CreateCharacter)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
