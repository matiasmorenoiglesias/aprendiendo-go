package funciones

import "fmt" 

func tabla(value int) func() int {
	num := value // 2, 2
	sec := 0 //0, 1

	return func() int {
		sec++ // 0 + 1 => 1, 1 + 1 => 2
		return num * sec // 2 * 1 => 2, 2 * 2 => 4
	}
}

func LlamarClosure(){
	myTabla := tabla(2)
	for i:=1; i < 11; i++ {
		fmt.Println(myTabla())
	}
}