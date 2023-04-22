package main

import (
	"fmt"
)

func main() {
	var opcion int
	fmt.Println("------------------------------------------------")
	fmt.Println("Bienvenido al Proyecto 2 de laboratorio Sopes 2")
	fmt.Println("***** Julio José Orellana Ruíz			")
	fmt.Println("*****		201908120					")
	fmt.Println("------------------------------------------------")
	for opcion != 3 {
		fmt.Println("Escriba la opcion que desea elegir: ")
		fmt.Println("1. Dar permisos")
		fmt.Println("2. No dar permisos")
		fmt.Println("3. Salir")

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Has seleccionado la opción de dar permisos")
		case 2:
			fmt.Println("Has seleccionado la opción de no dar permisos")
		case 3:
			fmt.Println("\nGracias por usar nuestro programa. ¡Hasta luego!")
		default:
			fmt.Println("Opción inválida. Por favor, seleccione una opción válida.\n")
		}
	}
}
