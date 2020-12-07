package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	v1 := caminhonete{veiculo{4, "Preto"}, true}
	v2 := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "Branco"},
		modeloLuxo: false}

	fmt.Println(v1.veiculo.portas)
	fmt.Println(v1.veiculo.cor)
	fmt.Println(v1.tracaoNasRodas)
	fmt.Println(v2.veiculo.cor)
	fmt.Println(v2.veiculo.portas)
	fmt.Println(v2.modeloLuxo)
}