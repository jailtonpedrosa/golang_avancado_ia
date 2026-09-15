package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex // não pode ser uma variável do mesmo escopo da atribuição ou leitura da goroutine
		wg sync.WaitGroup // eu digo para o wait group quantas goroutines eu quero que ele espere até que ele possa continuar 
		// a execução do meu código
	)

	wg.Add(100) //aqui estou setando 100 goroutines que meu waitgroup espere finalizar
	// para que ele continue a execução do meu programa

	for i := 0; i < 100; i++ {
		go func() {
			// posso fazer um defer aqui no wg.Done()
			defer wg.Done()
			
			mu.Lock() // trava a execução para outras goroutines
			m[i] = i
			mu.Unlock()

			//wg.Done() // aqui faz um decremento do meu contador de goroutines executadas
		}()
	}

	wg.Wait() // segura meu código até as 100 goroutines executar

	for chave, valor := range m {
		fmt.Println(chave, valor)
	}

	fmt.Println("fim da execução...")
}