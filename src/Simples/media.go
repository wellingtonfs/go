package main

import(
	"fmt"
)

func main(){
	var Num1, Num2 float32
	fmt.Printf("Digite um numero: ")
	fmt.Scan(&Num1)
	fmt.Printf("Digite outro numero: ")
	fmt.Scan(&Num2)

	media := (Num1 + Num2) / 2

	fmt.Println("\nResultado: ", media)
}