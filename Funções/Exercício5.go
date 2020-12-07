package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Área do quadrado:", resultado)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("Área do círculo:", resultado)
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}

func main() {
	q := quadrado{
		lado: 4.0,
	}

	c := circulo{
		raio: 4.0,
	}

	q.area()
	c.area()

	info(q)
	info(c)
}