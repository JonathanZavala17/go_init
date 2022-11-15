package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Jonathan"
	edad := 26

	fmt.Println("Mi noombre es", nombre, "y tengo ", edad, "años de edad")
	fmt.Printf("Hola 1, %s y su edad es %d\n", nombre, edad)
	fmt.Printf("Hola 2, %v y su edad es %v\n", nombre, edad)

	mensaje := fmt.Sprintf("Hola 3, %s y su edad es %d\n", nombre, edad)
	fmt.Println(mensaje)

	fmt.Printf("NOMBRE: %T \n", nombre)
	fmt.Printf("EDAD: %T \n", edad)

	fmt.Printf("Ingrese su otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre: ", nombre)

}
