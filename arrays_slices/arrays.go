package arrays_slices

import "fmt"

var tabla [10]int = [10]int {10,23,204,}
var matriz [20][30]int
func MuestroArreglo(){
	tabla[7] = 12
	tabla[5] = 2

	tabla2 := [10]string {"Matias", "Moreno"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla) ; i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][7] = 2
	fmt.Println(matriz)
}