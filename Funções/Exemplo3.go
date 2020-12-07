package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()

	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrução)
	}
}

func main() {

	dente := dentista{
		pessoa: pessoa{
			nome:      "Dentista",
			sobrenome: "Dente",
			idade:     29,
		},
		dentesarrancados: 3,
		salario:          3000,
	}

	casa := arquiteto{
		pessoa: pessoa{
			nome:      "Arquiteto",
			sobrenome: "Casa",
			idade:     35,
		},
		tipodeconstrução: "casa",
		tamanhodaloucura: "muita",
	}

	dente.oibomdia()
	casa.oibomdia()

	serhumano(dente)
	serhumano(casa)
}
