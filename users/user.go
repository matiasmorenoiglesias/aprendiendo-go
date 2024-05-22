package users

import (
	"fmt"
	"github.com/matiasmorenoiglesias/aprendiendo-go/modelos"
)

func AltaUsuario(){
	u := new(modelos.User)
	u.AddUser(1,"Matias")
	fmt.Println(u)
}