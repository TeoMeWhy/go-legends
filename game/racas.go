package game

import (
	"fmt"
	"go_legends/utils"
	"strconv"
	"strings"
)

type Raca struct {
	NomeRaca      string
	Modificadores map[string]int
}

func (r *Raca) GetNomeRaca() string {
	return r.NomeRaca
}

func (r *Raca) LoadMods() {

	raca := r.NomeRaca
	mods := map[string]int{}

	con := utils.ConectDB()
	defer con.Close()

	query := utils.ImportQuery("game/load_mods.sql")
	query = strings.ReplaceAll(query, "{raca}", raca)

	rows, err := con.Query(query)
	if err != nil {
		fmt.Println("Não foi possível executar a query")
	}

	var d, f, i string
	for rows.Next() {
		rows.Scan(&d, &f, &i)
	}

	mods["destreza"], _ = strconv.Atoi(d)
	mods["forca"], _ = strconv.Atoi(f)
	mods["inteligencia"], _ = strconv.Atoi(i)
	r.Modificadores = mods
}

func NewRaca(nomeRaca string) *Raca {
	raca := &Raca{NomeRaca: nomeRaca}
	raca.LoadMods()
	return raca
}
