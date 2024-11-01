package main

import "fmt"

func main(){
	var htmlElements = make(map[string]string)

	htmlElements["p"] = "Paragraph"
	htmlElements["img"] = "Image"
	htmlElements["h1"] = "Heading One"
	htmlElements["h2"] = "Heading Two"
	fmt.Println(htmlElements)
}
