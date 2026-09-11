package main

import "fmt"

func main() {
	fmt.Println("----------Shallow Copy----------")
	matriz1 := [][]int{
		{1, 2},
		{3, 4},
	}
	
	matriz2 := make([][]int, len(matriz1))
	
	copy(matriz2, matriz1)
	
	matriz2[0][0] = 99 // se fizer assim altera as duas matrizes
	// matriz2[0] = []int{99,99} // se fizer assim altera somente a matriz 2
	
	fmt.Println("matriz 1", matriz1)
	fmt.Println("matriz 2", matriz2)
	
	fmt.Println("----------Deep Copy----------")
	matriz3 := [][]int{
		{1, 2},
		{3, 4},
	}
	
	matriz4 := deepCopy(matriz3)
	matriz4[0][0] = 99

	fmt.Println("matriz 3", matriz3)
	fmt.Println("matriz 4", matriz4)
}

// é para que todas as copias que eu fizer pare de copiar o array interno
func deepCopy(matriz [][]int) [][]int {
	destino := make([][]int, len(matriz))

	for i, slice := range matriz {
		destino[i] = make([]int, len(slice))
		copy(destino[i], slice)
	}

	return destino
}