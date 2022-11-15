package main

import "fmt"

func main() {
	//variables comunes
	var nombre1 string
	nombre1 = "Jonathan"

	var nombre2 string = "Rene"

	edad := 24

	fmt.Println(nombre1, nombre2, edad)

	//tipos de variables
	var a int = 27
	var b float32 = 2.77
	var bb float64 = 3.141592
	var c string = "Hi"
	var d bool = true

	fmt.Println(a, b, bb, c, d)

	//constantes definidas
	const pi = 3.14151992255
	fmt.Println(pi)

}
