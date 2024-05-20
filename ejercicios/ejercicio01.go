package ejercicios

import "strconv"

func RunEjercicio01(value string) (int, string){
	switch intvalue, _ := strconv.Atoi(value); {
	case intvalue > 100:
		return intvalue, ("Es mayor que 100")
	case intvalue == 100:
		return intvalue, ("Es igual a 100")
	default:
		return intvalue, ("Es menor a 100")
	}
}