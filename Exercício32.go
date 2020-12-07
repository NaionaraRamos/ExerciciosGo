package main

import (
	"fmt"
)

func main() {
	ss := map[string][]string{
		"nai_ramos": []string{
			"ler", "programar",
		},
		"ilan_vitor": []string{
			"jogar", "torrar a paciencia",
		},
		"roni_ramos":[]string{
			"sair com os amigos", "jogar futebol",
		},
	}

	fmt.Println(ss)
	
	ss["mel_doguinho"] = []string{"correr", "brincar"}
	
	for k, v := range ss {
		fmt.Println(k, v)
	}
	
	delete(ss, "nai_ramos")
	
	for k, v := range ss {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}