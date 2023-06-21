package main

import "fmt"

func main() {
	var tamanho int
	fmt.Println("Digite o tamanho da slice:")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)
	var soma int

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor do elemento : %d", i)
		fmt.Scan(&slice[i])
		soma += slice[i]
	}
	fmt.Println("Slice:", slice)
	fmt.Println("Soma:", soma)

}
