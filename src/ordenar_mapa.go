package main

import(
	"fmt"
	"sort"
)

func main(){
	mapa := make(map[int]int)
	for i := 1; i <= 15; i++{
		mapa[i] = i*i
	}

	fmt.Println(mapa) //sem ordem

	num := make([]int, 0, len(mapa))
	for chave, _ := range mapa{ //poderia ser for chave := range mapa
		num = append(num, chave)
	}

	sort.Ints(num)

	for _, n := range num{
		fmt.Printf("%d:%d ", n, mapa[n])
	}
	fmt.Printf("\n")

}