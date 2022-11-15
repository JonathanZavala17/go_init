package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola,", nombre)

}

func sumar(a int, b int) {
	resultado := a + b
	fmt.Println("el resultado es: ", resultado)

}

func restar() {
	var f int
	var g int
	var result int
	fmt.Printf("iNGRESE UN NUMERO PRIMERO: ")
	fmt.Scanln(&f)
	fmt.Printf("iNGRESE UN NUMERO SEGUNDO: ")
	fmt.Scanln(&g)
	result = f - g
	fmt.Println("el resultado es: ", result)

}

func multi(h, j int) int {
	return h * j
}

func division() {
	var k int
	var l int
	var divisione int
	fmt.Printf("Ingrese un numero entero: ")
	fmt.Scanln(&k)
	fmt.Printf("Ingrese un numero entero 2: ")
	fmt.Scanln(&l)
	divisione = k / l
	fmt.Println("el resultado es: ", divisione)

}

func main() {
	saludar("Jonathan")
	sumar(2, 3)
	restar()
	r := multi(4, 5)
	fmt.Println("La multiplicacion es: ", r)
	division()

}
