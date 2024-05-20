package variables

import (
	"fmt"
)

func ShowBooleans(){
	var isTrue bool = true
	var isFalse bool = false
	UnoEsMayorQueCero := 1 > 0
	DosEsMayorQueCinco := 2 > 5
	fmt.Println("--------------------")
	fmt.Println("Booleans")
	fmt.Println("--------------------")
	fmt.Println("isTrue = ", isTrue)
	fmt.Println("isFalse = ", isFalse)
	fmt.Println("UnoEsMayorQueCero = ", UnoEsMayorQueCero)
	fmt.Println("DosEsMayorQueCinco = ", DosEsMayorQueCinco)
}