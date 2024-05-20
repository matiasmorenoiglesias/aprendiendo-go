package variables

import "fmt"

func ShowIntegers(){
	var integer int
	intde32 := int32(10)
	intde64 := int64(100)
	//intde64 = integer // Error ya que son diferentes enteros
	// intde64 = int64(integer)
	fmt.Println("integer comun = ", integer)
	fmt.Println("integer de 32 = ", intde32)
	fmt.Println("integer de 64 = ", intde64)
}