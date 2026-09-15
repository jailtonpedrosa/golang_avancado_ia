package main

import (
	"fmt"
	"time"
)

func main() {
	// race condition
	// 1). só ocorre quando temos go routines envolvidas, quando tenho duas ou mais partes do meu código
	// executando ao mesmo tempo

	// -> duas ou mais goroutines acessando o mesmo dado ao mesmo tempo
	// -> pelo menos um desses acessos tem que ser uma escrita

	m := make(map[int]int)
	
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	
	fmt.Println(m)
	
	m2 := make(map[int]int)
	
	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i
		}
	}()
	
	go func() {
		for i := 1000; i < 2000; i++ {
			//m2[i] = i //fatal error: concurrent map writes

			// para leitura daria o mesmo erro tbm 
			fmt.Println(m2[i]) //fatal error: concurrent map read and map write
		}
	}()

	time.Sleep(time.Second * 5)
}
