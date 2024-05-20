package main

import (
	"fmt"
	"strconv"

	"github.com/matiasmorenoiglesias/aprendiendo-go/condicionales"
	"github.com/matiasmorenoiglesias/aprendiendo-go/ejercicios"
	"github.com/matiasmorenoiglesias/aprendiendo-go/funciones"
)


func main(){
	status, value := funciones.ConvertIntToStr(15)
	if status {
		fmt.Println("El valor es",value)
	}
	condicionales.RunIf()
	condicionales.RunIfElse()
	{
		num, message := ejercicios.RunEjercicio01("101")
		fmt.Printf("num: %s, message: %s.\n",strconv.Itoa(num), message)
	}
	{
		num, message := ejercicios.RunEjercicio01("100")
		fmt.Printf("num: %s, message: %s.\n",strconv.Itoa(num), message)
	}
	{
		num, message := ejercicios.RunEjercicio01("99")
		fmt.Printf("num: %s, message: %s.\n",strconv.Itoa(num), message)
	}
}