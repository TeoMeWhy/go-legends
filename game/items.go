package game

import (
	"go_legends/utils"
	"strings"
)

type Item struct {
	Nome        string
	Descricao   string
	TipoItem    string
	Peso        int
	NivelMinimo int
	Atributos   map[string]int
}

func LoadItem(n string) *Item {

	con := utils.ConectDB()
	defer con.Close()

	query := utils.ImportQuery("game/load_item.sql")
	query = strings.ReplaceAll(query, "{item}", n)

	row := con.QueryRow(query)

	var nome, descricao, tipo string
	var att_forca, att_destreza, att_inteligencia int
	var peso_gramas, nivel_minimo int

	row.Scan(
		&nome, &descricao, &tipo,
		&att_forca, &att_destreza, &att_inteligencia,
		&peso_gramas, &nivel_minimo,
	)

	att := map[string]int{
		"forca":        att_forca,
		"destreza":     att_destreza,
		"inteligencia": att_inteligencia,
	}

	return &Item{
		Nome:        nome,
		Descricao:   descricao,
		TipoItem:    tipo,
		Atributos:   att,
		Peso:        peso_gramas,
		NivelMinimo: nivel_minimo,
	}
}
