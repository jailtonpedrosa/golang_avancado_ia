package main

import "fmt"

type mySlice []int

func main() {
	// as 3 funções map, filter e reduce trabalham com listas

	// map(aplica uma transformação em cada um dos itens da lista) -> [1,2,3,4,5] -> aplica por exemplo uma multiplicação
	// por 2 e devolve uma nova lista -> [2,4,6,8,10]

	// filter(filtra os dados) -> [1,2,3,4,5,6] -> aplica por exemplo filtro em numeros pares e devolve uma nova lista
	// [2,4,6]

	// reduce(retorna um dado final como somar os dados da lista) -> [1,2,3,4,5,6] -> 21

	var lista = mySlice{1, 2, 3, 4, 5, 6}
	lista = lista.Filter(func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(lista)

	lista = lista.Map(func(i int) int {
		return i * 2
	})

	fmt.Println(lista)

	soma := lista.Reduce(func(i1, i2 int) int {
		return i1 + i2
	}, 0)

	fmt.Println(soma)

	subtracao := lista.Reduce(func(i1, i2 int) int {
		return i1 - i2
	}, 0)

	fmt.Println(subtracao)
	
	multiplicacao := lista.Reduce(func(i1, i2 int) int {
		return i1 * i2
	}, 1)

	fmt.Println(multiplicacao)

	// encadeamento
	multiplicacao2 := lista.Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 10
	}).Reduce(func(i1, i2 int) int {
		return i1 * i2
	}, 1)
	
	fmt.Println(multiplicacao2)
}

func (m mySlice) Filter(condicao func(int) bool) mySlice {
	var resultado mySlice

	for _, numero := range m {
		if condicao(numero) {
			resultado = append(resultado, numero)
		}
	}

	return resultado
}

func (m mySlice) Map(transformacao func(int) int) mySlice {
	var resultado mySlice

	for _, numero := range m {
		resultado = append(resultado, transformacao(numero))
	}

	return resultado
}

func (m mySlice) Reduce(acumulador func(int, int) int, inicial int) int {
	var resultado = inicial

	for _, numero := range m {
		resultado = acumulador(resultado, numero)
	}

	return resultado
}
