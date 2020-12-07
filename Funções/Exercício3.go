package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Aparece depois...")
	fmt.Println("Aparece antes.")
}