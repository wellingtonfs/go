package main

//Pelo que testei, a conversao foi feita automaticamente

import(
	"fmt"
)

type Lista []string

func main(){
	lista := Lista{"Ola", "amigos"}
	slice := []string{"Hello", "friends"}

	imprimir1([]string(lista))
	imprimir2(Lista(slice))

	imprimir1(lista) //funciona de qualquer maneira
	imprimir2(slice)
}

func imprimir1(slice []string){
	fmt.Println("Slice: ", slice)
}

func imprimir2(lista Lista){
	fmt.Println("Lista: ", lista)
}