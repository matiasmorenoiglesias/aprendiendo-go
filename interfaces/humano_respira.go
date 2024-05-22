package interfaces

import "fmt"

func HumanoRespirando(hu Humano) {
	hu.Respirar()
	fmt.Printf("soy un/a %s y estoy respirando.\n", hu.Sexo())
}