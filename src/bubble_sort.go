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

	fmt.Println(bubble_sort(numeros))
}

func bubble_sort(n []int) []int{
	tam := len(n) - 1
	for i := 0; i < len(n); i++{
		for j := 0; j < tam; j++{
			if n[j] > n[j+1]{
				aux := n[j]
				n[j] = n[j+1]
				n[j+1] = aux
			}
		}
		tam--
	}
	return n
}