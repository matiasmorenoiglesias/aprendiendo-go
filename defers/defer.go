package defers

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic(){
	defer func(){
		reco := recover()
		if reco != nil {
			log.Printf("ocurrio un error que genero panic.\n %v", reco)
			//log.Fatalf("ocurrio un error que genero panic.\n %v", reco)
		}
	}()
	a:=1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}