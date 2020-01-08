package main

import(
	"fmt"
	"os"
	"strings"
)

func main(){
	palavras := os.Args[1:]
	dicionario := pass_dicio(palavras)
	imprimir(dicionario)
}

func pass_dicio(palavras []string) map[string]int{
	dicionario := make(map[string]int)

	for _, valor := range palavras{
		letra := strings.ToUpper(string(valor[0]))

		_ , existe := dicionario[letra]
		if existe{
			dicionario[letra]++
		}else{
			dicionario[letra] = 1
		}
	}
	return dicionario
}

func imprimir(dicionario map[string]int){
	fmt.Println("Lista:\n")
	for letras, valor := range dicionario{
		fmt.Printf("%s = %d\n", letras, valor)
	}
}