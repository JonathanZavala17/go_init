package main

import "fmt"

func main() {
	//definimos las variables
	a := 20
	b := 10

	//suma
	result := a + b
	fmt.Println("Suma: ", result)

	//resta
	result = a - b
	fmt.Println("Resta: ", result)

	//multiplicacion
	result = a * b
	fmt.Println("Multiplicacion: ", result)

	//division
	result = a / b
	fmt.Println("Division: ", result)

	//modulo
	result = a % b
	fmt.Println("Modulo: ", result)

	//division float
	var divi float32 = 3.0 / 2
	fmt.Println("el resultado es: ", divi)

}
