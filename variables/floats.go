package variables

import (
	"fmt"
)

func ShowFloats(){
	var salary float32 = 3200.0
	PI := 3.1415926535
	fmt.Println("--------------------")
	fmt.Println("Floats")
	fmt.Println("--------------------")
	fmt.Println("My salary is ",salary )
	fmt.Println("PI value is ", PI)
}