package main

import(
	"fmt"
)

func main(){
	var n int

	loop: //serve para parar o loop externo
	for n = 0; n < 7; n++{
		switch n{
			case 3:
				fmt.Println("Parando..")
				break
			case 5:
				fmt.Println("Finalizando..")
				break loop
			default: //opcional
				fmt.Println(n)
				break
		}
	}
}