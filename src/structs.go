package main

import(
	"fmt"
	"errors"
)

type Pilha struct{
	valores []interface{}
}

type Arquivo struct{
	nome string
	tamanho int
	palavras int
	linhas int
}

func (pilha Pilha) Tamanho() int{
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool{
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}){
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error){
	if pilha.Vazia(){
		return nil, errors.New("Pilha Vazia!")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = append(pilha.valores[:pilha.Tamanho()-1])
	return valor, nil
}

func main(){
	duracell := Pilha{}

	arq1 := &Arquivo{nome: "structs.go", linhas: 6}
	arq2 := Arquivo{nome: "structs2.go", linhas: 7}

	fmt.Println("Teste: ", arq1, arq2)

	fmt.Println("Tamanho: ", duracell.Tamanho())

	duracell.Empilhar("Wellington")
	duracell.Empilhar(997054573)
	duracell.Empilhar("wellingtonsilveira99@gmail.com")
	duracell.Empilhar("trento")

	fmt.Println("Tamanho: ", duracell.Tamanho())

	for !duracell.Vazia(){
		v, _ := duracell.Desempilhar()
		fmt.Println("Desempilhando ", v)
		fmt.Println("Tamanho ", duracell.Tamanho())
	}
	_, erro := duracell.Desempilhar()
	if erro != nil{
		fmt.Println("Erro: ", erro)
	}
}
