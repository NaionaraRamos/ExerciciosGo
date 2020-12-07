package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (x pessoa) nomeEsobrenome (){
	fmt.Println("Nome completo:", x.nome, x.sobrenome, "\nIdade:", x.idade)
}

func main() {
	eu := pessoa {
		nome: "Eu",
		sobrenome: "de Novo",
		idade: 19,
		}
	eu.nomeEsobrenome()
}