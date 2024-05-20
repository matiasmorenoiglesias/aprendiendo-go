package main

import (
	"fmt"

	"github.com/matiasmorenoiglesias/aprendiendo-go/condicionales"
	"github.com/matiasmorenoiglesias/aprendiendo-go/funciones"
)


func main(){
	status, value := funciones.ConvertIntToStr(15)
	if status {
		fmt.Println("El valor es",value)
	}
	condicionales.RunIf()
	condicionales.RunIfElse()
}