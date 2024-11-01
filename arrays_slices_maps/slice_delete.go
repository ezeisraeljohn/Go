package main

import "fmt"

func main(){
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgonene"
	cheeses = append(cheeses, "Camembert")
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
