package main

import (
	"fmt"
)

func main() {
	essareceberaoutrafuncao(issoseraumargumento)
}


func issoseraumargumento(){
	fmt.Println("Isso será um argumento...")
}

func essareceberaoutrafuncao(x func()) {
	fmt.Println("Recebendo outra função...")
	x()
}