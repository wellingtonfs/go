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

	fmt.Println(quicksort(numeros))
}

func particionar(nn []int, p int) (menores []int, maiores []int){
	for _, v := range nn{
		if v <= p{
			menores = append(menores, v)
		}else{
			maiores = append(maiores, v)
		}
	}
	return menores, maiores
}

func quicksort(n []int) []int{
	if len(n) <= 1{
		return n
	}

	i_pivo := len(n) / 2
	n_pivo := n[i_pivo]
	n = append(n[:i_pivo], n[i_pivo+1:]...)

	menores, maiores := particionar(n, n_pivo)

	return append( append(quicksort(menores), n_pivo), quicksort(maiores)... )
}