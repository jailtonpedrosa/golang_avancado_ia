package main

import "fmt"

func main() {
	fmt.Println("---------- Copia Insegura ----------")
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice1[0] = 10

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	// aqui já cria um novo slice
	slice2 = append(slice2, 6)
	slice2[0] = 99

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	fmt.Println("---------- Copia Segura ----------")

	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, len(slice3)) // make instancia um tipo especifico pra mim com valores em branco
	copy(                                 // destino dst
		slice4,
		// source
		slice3,
	)

	slice3[0] = 10 // não afeta o slice 4
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Slice 4:", slice4)

	// no argumento de tamanho da função make posso colocar o tamanho que eu quiser, tipo criar um slice
	// menor que o 3 ele ao copiar copia até o tamanho q foi definido sem dar erro
	// e caso criar um slice maior que o 3 ele ao copiar preenche com 0(zeros) os espaços vazios
}
