package files

import (
	"bufio"
	"fmt"
	"os"
	"github.com/matiasmorenoiglesias/aprendiendo-go/ejercicios"
)

var filePathName string = "./files/txt/tabla.txt"

func GrabarTabla(){
	var text string = ejercicios.RunEjercicio02()
	file, err := os.Create(filePathName)
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func append(fileFile string, text string) bool {
	file, err := os.OpenFile(fileFile, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error durante el Append %s.", err.Error())
		return false
	}
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error durante la escritura del archivo %s.", err.Error())
		return false
	}
	file.Close()
	return true
}

func SumarTabla(){
	var text string = ejercicios.RunEjercicio02()
	if !append(filePathName, text) {
		fmt.Println("Error al concatenar el texto")
	}
}

func ReadFile(){
	file, err := os.Open(filePathName)
	if err != nil {
		fmt.Printf("Hubo un error en la lectura del archivo %s.", err.Error())
		return
	}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		row := sc.Text()
		fmt.Printf("> %s\n",row)
	}
	file.Close()
}