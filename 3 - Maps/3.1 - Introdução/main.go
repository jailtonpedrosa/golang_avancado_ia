package main

import "fmt"

func main() {
	// estrutura de chave-valor
	// tem vantagem sobre o struct pois não preciso declarar explicitamente os campos

	//map[tipo da chave] tipo do valor
	var meuMap = map[string]string{
		"chave1": "valor1",
		"chave2": "valor2",
	}
	fmt.Println(meuMap)

	/*
		IMPORTANTE -> SINTAXE COMUM DE SE VER COM TIPO ANY
		ONDE ANY É A MESMA COISA QUE interface{} que é um tipo que pode ser utilizado como qualquer tipo válido do go
		ISSO É A MELHOR REPRESENTAÇÃO DE UM JSON NO GO
		var meuMap = map[string]any{
			"chave1": "valor1",
			"chave2": "valor2",
		}
	*/

	var meuMap1 = map[string]any{
		"chave1": 1,
		"chave2": "valor",
		"chave3": false,
	}

	fmt.Println(meuMap1)

	var meuMap2 = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	// *** QUANDO VC ITERA SOBRE UM MAP NÃO TEM GARANTIA DE ORDENAÇÃO ***
	for chave, valor := range meuMap2 {
		fmt.Println(chave, valor)
	}

	// CENÁRIOS QUE PODEM GERAR PANIC
	// 1). Declarando map sem sinal de igual e sem chaves
	//var meuMap3 map[string]any // declarando o map assim seus valores zero são nil
	// panic: assignment to entry in nil map
	//meuMap3["chave"] = "valor"

	// É NECESSÁRIO INICIALIZAR O MAP COLOCANDO O SINAL DE IGUAL(=) E ABRINDO E FECHANDO CHAVES {}
	var meuMap4 = map[string]any{}
	meuMap4["chave"] = "valor"

	// OUTRA FORMA DE INICIALIZAR O MAP É USANDO A FUNÇÃO MAKE
	var meuMap5 = make(map[string]any)
	meuMap5["chave"] = "valor"

	// 2). Tentar acessar um campo que não existe dentro do seu map
	fmt.Println(meuMap4["chaave____"])

	//var meuValorStr = meuMap4["campo_que_nao_existe"].(string) // tentando converter nil pra string, não dá certo
	//var meuValorInt = meuMap4["campo_que_nao_existe"].(int) // tentando converter nil pra string, não dá certo
	//fmt.Println(meuValorStr) //panic: interface conversion: interface {} is nil, not string
	//fmt.Println(meuValorInt) //panic: interface conversion: interface {} is nil, not string

	// FORMA DE ACESSAR O VALOR DA CHAVE PARA VER SE A CHAVE ESTÁ OK, DECLARE O "ok" COMO ABAIXO
	// funciona com ou sem casting(é a conversão do .(string))
	var meuValor, ok = meuMap4["campo_que_nao_existe"].(string)

	if ok {
		fmt.Println(meuValor)
	} else {
		fmt.Println("valor não encontrado")
	}
}
