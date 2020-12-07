package main

import (
	"fmt"
)

func main() {

	esportefavorito := "tenis"

	switch esportefavorito {
	case "tenis":
		fmt.Println("Zero")
	case "volei":
		fmt.Println("Um")
	default:
		fmt.Println("Ahhhh")
	}
}