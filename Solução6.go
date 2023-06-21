package main

import "fmt"

func main() {

	array := [10]int{7, 9, 10, 2, 5, 11, 6, 8, 3, 15}

	var valor int
	fmt.Println("Digite um valor: ")
	fmt.Scanln(&valor)

	encontrado := false
	for _, v := range array {
		if v == valor {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Printf("O valor %d está no array", valor)
	} else {
		fmt.Printf("O valor %d não está no array", valor)
	}
}
