package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, chan1 chan bool){
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	chan1 <- true
}