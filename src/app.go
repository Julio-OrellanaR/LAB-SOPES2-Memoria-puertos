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
			cmd := exec.Command("chmod", "+x" ,"/home/julio-or/Documentos/USAC/9no Semestre/Lab Sopes 2/Proyecto 2/LAB-SOPES2-Memoria-puertos/src/Scripts/permisos.sh")
			output, err := cmd.Output()

			if err != nil {
				fmt.Println("Error al ejecutar el comando opcion 1:", err)
				return
			}

			fmt.Println("Has seleccionado la opción de dar permisos: --- ", string(output), "\n")

			// Observando la USB
			fmt.Println("Observando la USB...")
			for {
				cmd2 := exec.Command("inotifywait", "-e", "create,delete,move", "/media/usb/")
				output, err := cmd2.Output()
				if err != nil {
					fmt.Println("Error al ejecutar inotifywait:", err)
					return
				}
				fmt.Println("Se detectó una acción en la USB: ", string(output))
			}


		case 2:
			cmd := exec.Command("chmod", "+x" ,"/home/julio-or/Documentos/USAC/9no Semestre/Lab Sopes 2/Proyecto 2/LAB-SOPES2-Memoria-puertos/src/Scripts/noPermisos.sh")
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
