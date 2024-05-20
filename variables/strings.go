package variables

import (
	"fmt"
	"strings"
)

func ShowStrings(){
	var myName string = "Matias Moreno Iglesias"
	githubUsername := "matiasmorenoiglesias"
	firstName := strings.Split(myName, " ")[0]
	fmt.Println("--------------------")
	fmt.Println("Strings")
	fmt.Println("--------------------")
	fmt.Println("My Name is ", myName)
	fmt.Println("My Github @",githubUsername)
	fmt.Println("My FirstName is", firstName)
}