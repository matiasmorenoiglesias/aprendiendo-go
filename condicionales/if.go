package condicionales

import (
	"fmt"
	"runtime"
)

func RunIf() {
	if os := runtime.GOOS; os == "darwin" || os == "linux" {
		fmt.Println("Your Operating System is not Windows =), is ", os)
	}
}
