package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunEjercicio02(){
	var num int
	var err error
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("ingrese un numero: ")
		if sc.Scan() {
			num, err = strconv.Atoi(sc.Text())
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Tabla de multiplicacion del numero %d\n", num)
				for i:=1;i<11;i++ {
					fmt.Printf("%d x %d = %d\n", num, i, (num*i))
				}
				break
			}
		}
	}
}