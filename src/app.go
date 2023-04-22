package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var opcion int
	fmt.Println("------------------------------------------------")
	fmt.Println("Bienvenido al Proyecto 2 de laboratorio Sopes 2")
	fmt.Println("***** Julio José Orellana Ruíz			")
	fmt.Println("*****		201908120					")
	fmt.Println("------------------------------------------------")

	// Ejecuta el comando "sudo su"
	cmd1 := exec.Command("sudo", "su")
	if err := cmd1.Run(); err != nil {
		fmt.Println("Error al ejecutar sudo su:", err)
		return
	}

	for opcion != 3 {
		fmt.Println("Escriba la opcion que desea elegir: ")
		fmt.Println("1. Dar permisos")
		fmt.Println("2. No dar permisos")
		fmt.Println("3. Salir")

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			cmd := exec.Command("sudo", "chmod", "777", "/media")
			output, err := cmd.Output()

			if err != nil {
				fmt.Println("Error al ejecutar el comando opcion 1:", err)
				return
			}

			fmt.Println("Has seleccionado la opción de dar permisos: --- ", string(output), "\n")

			// Observando la USB
			fmt.Println("Observando la USB...")
			fmt.Println("Combinacion de teclas 'ctrl + c' para detener la observación de la USB...")
			for {
				cmd2 := exec.Command("inotifywait", "-e", "create,delete,move", "/media/julio-or/UBUNTU 20_0")
				output, err := cmd2.Output()
				if err != nil {
					fmt.Println("Error al ejecutar inotifywait:", err)
					return
				}
				fmt.Println("Se detectó una acción en la USB: ", string(output))
			}

		case 2:
			cmd := exec.Command("sudo", "chmod", "000", "/media")
			output, err := cmd.Output()

			if err != nil {
				fmt.Println("Error al ejecutar el comando:", err)
				return
			}

			fmt.Println("Has seleccionado la opción de no dar permisos: ", string(output), "\n")
		case 3:
			fmt.Println("\nGracias por usar nuestro programa. ¡Hasta luego!")
		default:
			fmt.Println("Opción inválida. Por favor, seleccione una opción válida.\n")
		}
	}
}
