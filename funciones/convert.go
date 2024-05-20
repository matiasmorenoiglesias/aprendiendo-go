package funciones

import "strconv"

func ConvertIntToStr(number int) (bool, string){
	return true,strconv.Itoa(number)
}