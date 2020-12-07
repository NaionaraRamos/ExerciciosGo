package main

import (
	"fmt"
)

func main() {

	anoemqueeunasci := 1996
	anoatual := 2020

	for {
		if anoemqueeunasci > anoatual {
			break
		}

		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}
