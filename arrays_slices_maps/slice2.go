package main 

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoises de Bourgogne"
	cheeses = append(cheeses, "Camembert", "Reblochon", "Picodon")
	fmt.Println(cheeses)
}
