package character

import (
	"fmt"
	"go_legends/utils"
	"strings"
)

type Character struct {
	Nome          string
	Raca          string
	Classe        string
	Atributos     map[string]int
	Modificadores map[string]int
	Exp           int
	Nivel         int
	Pontos        int
	Vitalidade    int
}

func (c *Character) SetNivel() {
	c.Nivel = GetNivel(c.Exp)
}

func (c *Character) SetVitalidade() {
	c.Vitalidade = 10 + c.Modificadores["forca"]
}

func (c *Character) AddXP(xp int) {
	c.Exp += xp
	for novoNivel := GetNivel(c.Exp); novoNivel > c.Nivel; novoNivel-- {
		c.Pontos += nivelPonto[novoNivel]
	}
	c.SetNivel()
}

func NewChar(nome, raca, classe string) *Character {

	mods := loadMods(raca)
	att := loadAtributos(classe)

	char := Character{
		Nome:          nome,
		Raca:          raca,
		Classe:        classe,
		Atributos:     att,
		Modificadores: mods,
	}

	char.SetNivel()
	char.SetVitalidade()
	return &char
}

func SaveChar(c *Character) {

	con := utils.ConectDB()
	defer con.Close()

	attForca := c.Atributos["forca"]
	attDestreza := c.Atributos["destreza"]
	attInteligencia := c.Atributos["inteligencia"]
	modForca := c.Modificadores["forca"]
	modDestreza := c.Modificadores["destreza"]
	modInteligencia := c.Modificadores["inteligencia"]

	query := "DELETE FROM chars WHERE nome = '{nome}';"
	query = strings.ReplaceAll(query, "{nome}", c.Nome)
	con.Exec(query)

	query = `INSERT INTO chars (
		nome,
		raca,
		classe,
		exp,
		nivel,
		pontos,
		vitalidade,
		att_forca,
		att_destreza,
		att_inteligencia,
		mod_forca,
		mod_destreza,
		mod_inteligencia
	)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`

	stmt, err := con.Prepare(query)
	if err != nil {
		fmt.Println("Não foi possível preparar a query")
	}

	_, err = stmt.Exec(c.Nome,
		c.Raca,
		c.Classe,
		c.Exp,
		c.Nivel,
		c.Pontos,
		c.Vitalidade,
		attForca,
		attDestreza,
		attInteligencia,
		modForca,
		modDestreza,
		modInteligencia)

	if err != nil {
		fmt.Println("Não foi possível executar o statement.")
	}

}

func LoadChar(nome string) *Character {
	con := utils.ConectDB()
	defer con.Close()

	query := "SELECT * FROM chars WHERE nome = '{nome}'"
	query = strings.ReplaceAll(query, "{nome}", nome)

	rows, err := con.Query(query)
	if err != nil {
		fmt.Println("Erro ao executar a query de char")
	}

	var raca, classe string
	var exp, nivel, pontos, vitalidade int
	var attForca, attDestreza, attInteligencia int
	var modForca, modDestreza, modInteligencia int

	for rows.Next() {
		rows.Scan(
			&nome,
			&raca,
			&classe,
			&exp,
			&nivel,
			&pontos,
			&vitalidade,
			&attForca,
			&attDestreza,
			&attInteligencia,
			&modForca,
			&modDestreza,
			&modInteligencia,
		)
	}

	return &Character{
		Nome:   nome,
		Raca:   raca,
		Classe: classe,
		Atributos: map[string]int{
			"forca":        attForca,
			"destreza":     attDestreza,
			"inteligencia": attInteligencia,
		},
		Modificadores: map[string]int{
			"forca":        modForca,
			"destreza":     modDestreza,
			"inteligencia": modInteligencia,
		},
		Exp:        exp,
		Nivel:      nivel,
		Pontos:     pontos,
		Vitalidade: vitalidade,
	}

}
