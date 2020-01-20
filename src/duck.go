package main

import(
	"fmt"
	"time"
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

type Idade struct{
	anoIni int
}

func (i Idade) Calcular() int{
	return time.Now().Year() - i.anoIni
}

func (i Idade) String() string{
	return fmt.Sprintf("Idade desde %d", i.anoIni)
}

func (s Soma) String() string{
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}

func (s Subtracao) String() string{
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2)
}

func acumular(oprs []Operacao) int{
	acumulador := 0

	for _, op := range oprs{
		valor := op.Calcular()
		fmt.Printf("%v = %d\n", op, valor)
		acumulador += valor
	}

	return acumulador
}

func main(){
	ops := make([]Operacao, 4)
	ops[0] = Soma{2, 3}
	ops[1] = Subtracao{10, 7}
	ops[2] = Subtracao{2, 3}
	ops[3] = Soma{3, 3}

	fmt.Println("Valor acumulado =", acumular(ops))

	idades := make([]Operacao, 4)
	idades[0] = Idade{1972}
	idades[1] = Idade{1973}
	idades[2] = Idade{1993}
	idades[3] = Idade{1999}

	fmt.Println("Idades acumuladas =", acumular(idades))
}