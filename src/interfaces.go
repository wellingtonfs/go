package main

import(
	"fmt"
)

type Operacao interface{
	Calcular() int
}

type Soma struct{
	operando1, operando2 int
}

type Subtracao struct{
	operando1, operando2 int
}

func (s Soma) Calcular() int{
	return s.operando1 + s.operando2
}

func (s Subtracao) Calcular() int{
	return s.operando1 - s.operando2
}

type Num int

func (s Num) Zerar() int{
	return 0
}

func main(){
	op := Soma{2, 4}
	po := Subtracao{5, 4}
	
	var oo Operacao
	oo = Soma{1, 1}


	fmt.Println(op.Calcular(), po.Calcular(), oo.Calcular())

}
