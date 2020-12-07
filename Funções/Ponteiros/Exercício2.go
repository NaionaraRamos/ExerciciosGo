package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa){
	(*p).nome = "Sobrenome"
	p.idade	= 1
}

func main() {

	x := pessoa{
		nome: "Nome",
		idade: 0,
	}
	
	fmt.Println(x)
	mudeMe(&x)
	fmt.Println(x)
}