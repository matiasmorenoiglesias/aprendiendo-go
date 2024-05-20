package main

import (
	"fmt"

	"github.com/matiasmorenoiglesias/aprendiendo-go/conditionals"
	"github.com/matiasmorenoiglesias/aprendiendo-go/functions"
)


func main(){
	status, value := functions.ConvertIntToStr(15)
	if status {
		fmt.Println("El valor es",value)
	}
	conditionals.RunIf()
	conditionals.RunIfElse()
}