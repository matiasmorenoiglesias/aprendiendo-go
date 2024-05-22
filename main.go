package main

import (
	//"fmt"

	//"github.com/matiasmorenoiglesias/aprendiendo-go/ejercicios"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/files"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/arrays_slices"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/mapas"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/defers"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/goroutines"
	"github.com/matiasmorenoiglesias/aprendiendo-go/webserver"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/interfaces"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/users"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/funciones"
)

func main() {
	/* chan1 := make(chan bool)
	go goroutines.MiNombreLento("Matias", chan1)
	defer func(){
		<- chan1
	}()
	println("Estoy Aqui!") */

	webserver.MyWebServer()
}
