package arrays_slices

import "fmt"

var table []int = []int{1,2,3}
var vector [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
func MuestroSlices(){
	fmt.Println(table)
	//fmt.Println(len(table), cap(table))
	fmt.Println("vector[3:]: ",vector[3:])
	fmt.Println("vector[:5]: ",vector[:5])
	fmt.Println("vector[6:7]: ",vector[6:7])
}

func Capacidad(){
	elements := make([]int, 0, 1)
	fmt.Printf("len: %d, cap: %d.\n", len(elements), cap(elements))
	fmt.Println("Datos:", elements)

	for i:=1; i<=100; i++ {
		elements = append(elements, i)
	}
	fmt.Printf("len: %d, cap: %d.\n", len(elements), cap(elements))
}