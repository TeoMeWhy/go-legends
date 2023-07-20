package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConectDB() *sql.DB {
	con, err := sql.Open("sqlite3", "data/data.db")
	if err != nil {
		fmt.Println("Não foi possível abrir conexão com o banco de dados.")
	}
	return con
}

func ImportQuery(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Não foi possível abrir o arquivo da query.")
	}
	return string(content)
}
