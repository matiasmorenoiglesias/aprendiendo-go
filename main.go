package main

import (
	//"fmt"

	//"github.com/matiasmorenoiglesias/aprendiendo-go/ejercicios"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/files"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/arrays_slices"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/mapas"
	"github.com/matiasmorenoiglesias/aprendiendo-go/interfaces"
	"github.com/matiasmorenoiglesias/aprendiendo-go/users"
	//"github.com/matiasmorenoiglesias/aprendiendo-go/funciones"
)


func main(){
	//arrays_slices.MuestroArreglo()
	//arrays_slices.MuestroSlices()
	//arrays_slices.Capacidad()
	//mapas.MostrarMapas()
	users.AltaUsuario()

	Pedro := new(users.Hombre)
	Maria := new(users.Mujer)

	interfaces.HumanoRespirando(Pedro)
	interfaces.HumanoRespirando(Maria)
}