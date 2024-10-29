package main

import (
	"fmt"
)

func verificarEdadParaLicencia(edad int) string {
	if edad >= 18 {
		return "Puedes sacar el carnet de conducir."
	}
	return "No puedes sacar el carnet de conducir."
}

func main() {
	var edad int
	fmt.Print("Introduce tu edad: ")
	_, err := fmt.Scan(&edad)
	if err != nil {
		fmt.Println("Por favor, introduce un número válido.")
		return
	}
	fmt.Println(verificarEdadParaLicencia(edad))
}