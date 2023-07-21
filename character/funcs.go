package character

import (
	"fmt"
	"go_legends/utils"
	"sort"
	"strconv"
	"strings"
)

func loadAtributos(classe string) map[string]int {

	att := map[string]int{}

	con := utils.ConectDB()
	defer con.Close()

	query := utils.ImportQuery("character/load_att.sql")
	query = strings.ReplaceAll(query, "{classe}", classe)

	rows, err := con.Query(query)
	if err != nil {
		fmt.Println("Não foi possível executar a query")
	}

	var att1, att2, att3 string
	for rows.Next() {
		rows.Scan(&att1, &att2, &att3)
	}
	atts := []string{att1, att2, att3}

	dices := utils.RollsDiceN(5, 6)
	sort.Sort(sort.Reverse(sort.IntSlice(dices)))
	dices = dices[1 : len(dices)-1]

	for i := range atts {
		att[atts[i]] = dices[i] + 8
	}

	return att
}

func loadMods(raca string) map[string]int {
	mods := map[string]int{}

	con := utils.ConectDB()
	defer con.Close()

	query := utils.ImportQuery("character/load_mods.sql")
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
	return mods
}
