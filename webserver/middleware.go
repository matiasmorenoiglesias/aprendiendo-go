package webserver

import "fmt"

func sumar(num1 int, num2 int ) int {
	return num1 + num2
}

func restar(num1 int, num2 int ) int {
	return num1 - num2
}

func multi(num1 int, num2 int ) int {
	return num1 * num2
}

func div(num1 int, num2 int ) int {
	return num1 / num2
}

func operacion(f func(int, int) int) func(int, int) int {
	return func (a, b int) int {
		fmt.Println("Inicio de Operacion")
		return f(a,b)
	}
}

func MyMiddleware(){
	fmt.Println("Inicio de Middleware")
	fmt.Println(operacion(sumar)(1,2))
	fmt.Println(operacion(restar)(5,6))
	fmt.Println(operacion(multi)(5,5))
	fmt.Println(operacion(div)(5,15))
}