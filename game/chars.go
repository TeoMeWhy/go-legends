package game

import (
	"fmt"
	"go_legends/utils"
	"strings"
)

type Character struct {
	Nome string
	Classe
	Raca
	Exp        int
	Nivel      int
	Pontos     int
	Vitalidade int
	Items      map[string]Item
}

func (c *Character) String() string {
	text := `
	%s: %s(%s)
	Forca: %d(%d) // Dest.: %d(%d) // Int.: %d(%d)
	`
	text = fmt.Sprintf(text,
		c.Nome, c.NomeRaca, c.NomeClasse,
		c.Atributos["forca"], c.Modificadores["forca"],
		c.Atributos["destreza"], c.Modificadores["destreza"],
		c.Atributos["inteligencia"], c.Modificadores["inteligencia"])

	return text
}

func (c *Character) SetNivel() {
	c.Nivel = GetNivel(c.Exp)
}

func (c *Character) SetVitalidade() {
	c.Vitalidade = 10 + c.Modificadores["forca"]
	for _, i := range c.Items {
		c.Vitalidade += i.Atributos["forca"]
	}
}

func (c *Character) SetItemInicial() {

	items := map[string]Item{}

	switch c.NomeClasse {
	case "guerreiro":
		items["arma"] = *LoadItem("Espada Curta")
	case "arqueiro":
		items["arma"] = *LoadItem("Arco Simples")
	case "clerigo":
		items["arma"] = *LoadItem("Cetro Iniciante")
	case "mago":
		items["arma"] = *LoadItem("Cetro Iniciante")
	case "ladino":
		items["arma"] = *LoadItem("Adaga da Cozinha")
	}
	c.Items = items
}

func (c *Character) AddXP(xp int) {
	c.Exp += xp
	for novoNivel := GetNivel(c.Exp); novoNivel > c.Nivel; novoNivel-- {
		c.Pontos += nivelPonto[novoNivel]
	}
	c.SetNivel()
}

func NewChar(nome, raca, classe string) *Character {

	char := Character{
		Nome:   nome,
		Raca:   *NewRaca(raca),
		Classe: *NewClasse(classe),
	}

	char.SetNivel()
	char.SetItemInicial()
	char.SetVitalidade()
	return &char
}

func SaveCharItens(c *Character) {

	con := utils.ConectDB()
	defer con.Close()

	stmt, err := con.Prepare("DELETE FROM chars_itens WHERE nome_char = ?;")
	if err != nil {
		fmt.Println("Erro em preparar o SQL para deleção de char/item")
		fmt.Println(err)
	}
	stmt.Exec(c.Nome)

	insert, err := con.Prepare("INSERT INTO chars_itens VALUES (?,?);")
	if err != nil {
		fmt.Println("Erro em preparar o SQL para inserir char/item")
	}

	for _, item := range c.Items {
		insert.Exec(c.Nome, item.Nome)
	}

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

	_, err = stmt.Exec(
		c.Nome, c.NomeRaca, c.NomeClasse,
		c.Exp, c.Nivel, c.Pontos,
		c.Vitalidade,
		attForca, attDestreza, attInteligencia,
		modForca, modDestreza, modInteligencia)

	if err != nil {
		fmt.Println("Não foi possível executar o statement.")
	}

	SaveCharItens(c)
}

func LoadCharItens(nome string) map[string]Item {

	itens := map[string]Item{}

	var nomeItem, descricao, tipoItem string
	var attForca, attDestreza, attInteligencia int
	var pesoGramas, nivelMinimo int

	con := utils.ConectDB()
	defer con.Close()

	query := utils.ImportQuery("game/load_char_itens.sql")
	query = strings.ReplaceAll(query, "{nome}", nome)

	rows, err := con.Query(query)
	if err != nil {
		fmt.Println("Erro aotentar buscar itens de char")
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(
			&nomeItem, &descricao, &tipoItem,
			&attForca, &attDestreza, &attInteligencia,
			&pesoGramas, &nivelMinimo,
		)

		itens[tipoItem] = Item{
			Nome:      nomeItem,
			Descricao: descricao,
			TipoItem:  tipoItem,
			Atributos: map[string]int{
				"forca":        attForca,
				"destreza":     attDestreza,
				"inteligencia": attInteligencia,
			},
			Peso:        pesoGramas,
			NivelMinimo: nivelMinimo,
		}
	}

	return itens
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
			&exp, &nivel, &pontos, &vitalidade,
			&attForca, &attDestreza, &attInteligencia,
			&modForca, &modDestreza, &modInteligencia,
		)
	}

	c := &Character{
		Nome:   nome,
		Raca:   *NewRaca(raca),
		Classe: *NewClasse(classe),

		Exp:        exp,
		Nivel:      nivel,
		Pontos:     pontos,
		Vitalidade: vitalidade,
	}

	c.Atributos = map[string]int{
		"forca":        attForca,
		"destreza":     attDestreza,
		"inteligencia": attInteligencia,
	}

	c.Modificadores = map[string]int{
		"forca":        modForca,
		"destreza":     modDestreza,
		"inteligencia": modInteligencia,
	}

	c.Items = LoadCharItens(nome)

	return c
}
