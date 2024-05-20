package conditionals

import (
	"fmt"
	"runtime"
)

func RunSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS is Darwin")
	case "linux":
		fmt.Println("Os is Linux")
	default:
		fmt.Printf("%s \n", os)
	}
}
