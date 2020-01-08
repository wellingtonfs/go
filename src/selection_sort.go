package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Precisa de, no mínimo, UM argumento!")
		os.Exit(1)
	}

	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, v := range entrada {
		valor, erro := strconv.Atoi(v)
		if erro != nil {
			fmt.Println("Todos os argumentos devem ser números inteiros!")
			os.Exit(1)
		}
		numeros[i] = valor
	}

	fmt.Println(selection_sort(numeros))
}

func selection_sort(n []int) []int{
	for i := 0; i < len(n); i++{
		for j := i; j < len(n); j++{
			if n[j] < n[i]{
				aux := n[i]
				n[i] = n[j]
				n[j] = aux
			}
		}
	}
	return n
}
