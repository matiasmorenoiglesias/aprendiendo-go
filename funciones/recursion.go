package funciones

import "fmt"


func Exponencia(num int){
	if num > 1000000000 {
		return
	}
	fmt.Println(num)
	Exponencia(num * 2)
}