package main

import(
	"fmt"
	"time"
)

func Soma(i... int) (vTotal int){
	vTotal = 0

	op := func(a, b int) int{
		return a + b
	}

	for _, valor := range i{
		vTotal = op(vTotal, valor)
	}

	return
}

func GerarFibonacci(n int) func(){
	return func(){
		a, b := 0, 1

		fib := func() int{
			a, b = b, a+b
			return a
		}

		for i := 0; i < n; i++{
			fmt.Printf("%d ", fib())
		}
		fmt.Printf("\n")
	}
}

func Cronometrar(f func()){
	inicio := time.Now()

	f()

	fmt.Println("Tempo = ", time.Since(inicio))
}

func Cronometrar2(f int){
	inicio := time.Now()

	for i := 0; i < f; i++{
		for i := 0; i < f; i++{
			for i := 0; i < f; i++{
				
			}
		}	
	}

	fmt.Println("Tempo = ", time.Since(inicio))
}

func main(){
	fmt.Println(Soma())
	fmt.Println(Soma(1, 2, 3))
	fmt.Println(Soma(5, 6, 7, 4, 1))

	Cronometrar2(3000)
}