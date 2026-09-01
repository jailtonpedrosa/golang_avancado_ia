package main

import "fmt"

func main() {
	var (
		// diferença do slice para o array é que o array tem tamanho fixo
		// o slice pode aumentar de tamanho
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println("Tamanho do array:", len(array))
	fmt.Println("Capacidade do array:", cap(array))

	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Capacidade do slice:", cap(slice))

	fmt.Println("----------")
	// adicionando um item no slice
	slice = append(slice, 6) //append só suporta slice
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Capacidade do slice:", cap(slice))
	
	// array[5] = 6 //invalid argument: index 5 out of bounds [0:5]
	printSlice(slice)
	
	fmt.Println("----------")
	slice = append(slice, 7, 8, 9, 10, 11, 12) // posso passar vários valores de uma vez no append
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Capacidade do slice:", cap(slice))
}

func printSlice(slice []int) {
	for indice, valor := range slice {
		fmt.Println("Indice:", indice, " - Valor:", valor)
	}
}
