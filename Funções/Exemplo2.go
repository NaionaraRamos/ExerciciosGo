package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func main() {
	mamis := pessoa{"Mamãe", 50}
	mamis.oibomdia()

}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz 'Bom dia!' e tem", p.idade, "anos.")
}
