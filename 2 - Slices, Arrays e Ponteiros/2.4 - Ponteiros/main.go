package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a // * é o ponteiro que aponta para o endereço de memória onde está o valor armazenado

	fmt.Println("variável a:", a)
	fmt.Println("variável b:", b)

	fmt.Println("----------")

	fmt.Println("variável a:", &a)
	fmt.Println("variável b(DESREFERENCIAMENTO):", *b)

	p := NovaPessoa("Jailton", 40)

	fmt.Println(p.Nome, p.Idade, p.telefone)
	fmt.Println(p.Nome, p.Idade, p.Telefone())
	p.AtualizarIdade(39)
	fmt.Println(p.Nome, p.Idade, p.Telefone())
	p.AtualizarTelefone("4002-8922")
	fmt.Println(p.Nome, p.Idade, p.Telefone())
}

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome: nome,
		Idade: idade,
	}
}

func (p Pessoa) Telefone() string {
	// (p Pessoa) é o receiver está atrelado amarrado a uma struct
	if p.telefone == nil {
		return ""
	}
	return *p.telefone
}

func (p *Pessoa) AtualizarIdade(idade int) {
	p.Idade = idade
}

func (p *Pessoa) AtualizarTelefone(telefone string) {
	p.telefone = &telefone
}