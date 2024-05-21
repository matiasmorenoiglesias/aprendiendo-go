package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var num1 int
var num2 int
var leyenda string
var err error

func MultiplicacionNumeros(){
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1:")
	if sc.Scan() {
		num1, err = strconv.Atoi(sc.Text())
		if err != nil {
			panic("Error en el ingreso del numero " + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2:")
	if sc.Scan() {
		num2, err = strconv.Atoi(sc.Text())
		if err != nil {
			panic("Error en el ingreso del numero " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda:")
	if sc.Scan() {
		leyenda = sc.Text()
	}

	fmt.Println(leyenda, num1 * num2)
}