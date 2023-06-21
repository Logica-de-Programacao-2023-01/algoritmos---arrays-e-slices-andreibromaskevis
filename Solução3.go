package main

import "fmt"

func main() {

	num1 := [4]float64{2.0, 3.0, 6.0, 4.0}

	produto := 1.0

	for _, valor := range num1 {
		produto *= valor
	}

	fmt.Printf("Resultado : %.2f", produto)

}
