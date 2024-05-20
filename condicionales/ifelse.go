package condicionales

import (
	"fmt"
	"runtime"
)

func RunIfElse() {
	os := runtime.GOOS
	if os == "linux" {
		fmt.Println("Your Operating System is Linux =)")
	} else if os == "darwin" {
		fmt.Println("Your Operating System is MacOs =)")
	} else {
		fmt.Println("Your Operating System is Windows or Another =)")
	}
}
