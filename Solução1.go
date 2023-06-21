package main

import "fmt"

func main() {

	var numeros = [3]int{5, 10, 20}
	fmt.Println("Os valores no array são", numeros)

	soma := 0

	for _, valor := range numeros {
		soma += valor
	}
	fmt.Printf("A soma dos valores no array é %d", soma)

}
