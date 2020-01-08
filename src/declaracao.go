package main

import(
	"fmt"
)

type Cidade struct{
	Estado string
	Codigo int
}

func main(){
	//tipos de declarações de variaveis (arrays)
	var a [3]int
	b := [3]int{1, 2, 3}
	c := [...]int{2, 3, 4, 2, 1, 2}
	d := [2]string{}
	var e [2]string
	var f [2][2]int

	f[0][1] = 2

	g := [2][3][2]int{{{1, 2}, {1, 2}, {1, 2}}, {{1, 2}, {1, 2}, {1, 2}}}

	fmt.Println(a, b, c, d, e, f, g)

	//tipos de declarações de variaveis (slices) OBS: só muda os []

	var h []int
	i := []int{1, 2, 3}
	j := []int{2, 3, 4, 2, 1, 2}
	k := []string{}

			//lista := make([]int, tam, capacidade) //capacidade é opcional, ela limita o tamanho de memoria reservada para o slice

	lista1 := make([]int, 10, 20)
	lista2 := make([]int, 10)

	fmt.Println(h, i, j, k, "\n", lista1, len(lista1), cap(lista1), "\n", lista2, len(lista2), cap(lista2))

	//tipos de declarações de variaveis (maps)

	dicio := make(map[int]string)
	dicio2 := map[int]string{}
	dicio3 := map[string]string{
		"BR" : "Brasil",
		"USA" : "Estados Unidos da America",
		"OI" : "Tchau"}

	dicio3["Ei"] = "ou"
	
	fmt.Println(dicio, dicio2, dicio3)

	delete(dicio3, "OI")

	fmt.Println(dicio, dicio2, dicio3, "\n\n")

	cidades := make(map[string]Cidade)
	cidades["Rio Grande"] = Cidade{"RS", 53}
	cidades["POA"] = Cidade{"RS", 51}
	cidades["Camboriu"] = Cidade{"SC", 48}

	for chave, valor := range cidades{
		fmt.Println(chave, "Estado:", valor.Estado, "Codigo de area:", valor.Codigo)
	}
}
