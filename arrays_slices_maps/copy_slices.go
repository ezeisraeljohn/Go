package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses[1:])
	fmt.Println(smellyCheeses)
}
