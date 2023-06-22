package main

import (
	"fmt"
)
func main() {
	array := [10]float64{2.5, 6.8, 1.2, 7.9, 3.4, 9.1, 4.8, 5.7, 8.2, 0.9}
	slice := make([]float64, 0)
	for _, value := range array {
		if value > 5 {
			slice = append(slice, value)
		}
	}
	fmt.Println("Novo slice:", slice)
}
