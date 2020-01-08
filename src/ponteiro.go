package main

import(
	"fmt"
)

type ListaGen []interface{}

func (lista *ListaGen) RemoverIndice(indice int) interface{}{
	removido := lista[indice]
	*lista = append(*lista[0:indice], *lista[indice+1:]...)
	return removido
}

func main(){
	list := ListaGen{
		1, "cafe", 123, false, "olho", 3, true}

	fmt.Println(list)
	list.RemoverIndice(0)

	fmt.Println(list)
}