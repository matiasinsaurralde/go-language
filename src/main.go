package main

import(	"fmt"
	"language" )

func main() {

	// language.LoadText("example")

	letters := language.ExtractLetters("texto de prueba")

	var k string
	var v float64

	for k,v = range letters {
		fmt.Printf("%s: %f\n", k, v)
	}

}
