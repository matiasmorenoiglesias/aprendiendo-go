package mapas

import "fmt"

func MostrarMapas(){
	countries := make(map[string]string)
	countries["Chile"] = "Santiago"
	countries["Argmentina"] = "Buenos Aires"
	fmt.Println(countries)
	fmt.Println(countries["Chile"])
	fmt.Println(countries["Undefined"])

	championship := map[string] int {
		"Barcelona": 39,
		"Real Madrid": 38,
		"Boca Jrs": 30,
	}

	fmt.Println(championship)

	for index, value := range championship {
		fmt.Printf("{ index: %s, value: %d }\n", index, value)
	}

	delete(championship, "Boca Jrs")
	fmt.Println(championship)

	championship["U de Chile"] = 36

	value, exists := championship["U de Chile"]

	if exists {
		fmt.Printf("El Equipo U de Chile clasifico con %d puntos.", value)
	}
}