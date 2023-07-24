package character

import (
	"fmt"
	"go_legends/utils"
	"sort"
	"strings"
)

type Mago struct {
}

func (m *Mago) AtaqueEspecial() int {
	return 1
}

type Guerreiro struct {
}

func (g *Guerreiro) AtaqueEspecial() int {
	return 2
}

type Arqueiro struct {
}

func (a *Arqueiro) AtaqueEspecial() int {
	return 3
}

type Ladino struct {
}

func (m *Ladino) AtaqueEspecial() int {
	return 4
}

type Clerigo struct {
}

func (m *Clerigo) AtaqueEspecial() int {
	return 5
}

type ClassePoder interface {
	AtaqueEspecial() int
}

type Classe struct {
	NomeClasse  string
	CargasPoder int
	Atributos   map[string]int
	ClassePoder
}

func (c *Classe) LoadAtributos() {

	att := map[string]int{}

	con := utils.ConectDB()
	defer con.Close()

	query := utils.ImportQuery("character/load_att.sql")
	query = strings.ReplaceAll(query, "{classe}", c.NomeClasse)

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

	c.Atributos = att
}

func NewClasse(classe string) *Classe {
	c := Classe{NomeClasse: classe, CargasPoder: 2}

	switch classe {
	case "mago":
		c.ClassePoder = &Mago{}
	case "guerreiro":
		c.ClassePoder = &Guerreiro{}
	case "arqueiro":
		c.ClassePoder = &Arqueiro{}
	case "ladino":
		c.ClassePoder = &Ladino{}
	case "clerigo":
		c.ClassePoder = &Clerigo{}
	}

	c.LoadAtributos()

	return &c
}
