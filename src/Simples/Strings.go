package main

import(
	"fmt"
)

type str string

func (s str) Tamanho() int{
	cont := 0
	for i, _ := range s{
		cont = i + 1
	}
	return cont
}

func (s str) contem(s2 string) bool{
	for _, v := range s2{
		tinha := true
		for _, v2 := range s{
			if v == v2{
				tinha = false
				break
			}
		}
		if tinha{
			return false
		}
	}
	return true
}

func main(){
	var ss str
	ss = "banana"

	os := "babas"

	fmt.Println("Tamanho: ", ss.Tamanho())
	fmt.Printf("'%s' contem as letras de '%s'?: %v\n", ss, os, ss.contem(os))
}