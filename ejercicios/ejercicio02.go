package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunEjercicio02() (string){
	var num int
	var err error
	var texto string
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("ingrese un numero: ")
		if sc.Scan() {
			num, err = strconv.Atoi(sc.Text())
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			} else {
				break
			}
		}
	}
	for i:=1;i<11;i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", num, i, (num*i))
	}
	return texto
}