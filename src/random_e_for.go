package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano()) //deixar aleatorio mesmo
	var n, i int

	loop: //serve para parar o loop externo
	for{
		i = -1
		n = 1
		for n % 24 != 0{
			n = rand.Intn(4200)
			fmt.Println(n)
			i++
			if n % 11 == 0{
				break loop
			}
		}
		fmt.Println("Chegouu", i)
	}
	fmt.Println("Fim")
}