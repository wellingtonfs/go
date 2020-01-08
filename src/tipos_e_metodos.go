package main

import(
	"fmt"
)

type ListaCompras []string

func main(){
	lista := make(ListaCompras, 6)

	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Laranja"
	lista[3] = "Bergamota"
	lista[4] = "Maminha"
	lista[5] = "Picanha"

	fmt.Println(lista.Tamanho())

	l, f, c := lista.Categorizar()

	fmt.Println(c, l, f)
}

func (lista ListaCompras) Tamanho() int{
	return len(lista)
}

func (lista ListaCompras) Categorizar() (legumes []string, frutas []string, carnes []string){
	for _, v := range lista{
		switch v{
			case "Alface", "Pepino":
				legumes = append(legumes, v)
			case "Laranja", "Bergamota":
				frutas = append(frutas, v)
			case "Maminha", "Picanha":
				carnes = append(carnes, v)
		}
	}
	return legumes, frutas, carnes
}