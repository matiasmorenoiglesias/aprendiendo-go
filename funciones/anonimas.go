package funciones

import "fmt"

func Calculos(){
	//var num1, num2 int
	calculo := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(calculo(10,10))
	calculo = func(n1 int, n2 int) int {
		return n1 * n2
	}

	fmt.Println(calculo(10,10))
}