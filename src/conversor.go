package main

import(
	"fmt"
	"strconv"
	"os"
)

func main(){
	if len(os.Args) < 3{
		fmt.Println("Uso: <valor> <valor> ... <valor> <unidade>")
		os.Exit(1)
	}

	var unidadeFinal string

	if os.Args[len(os.Args) - 1] != "celsius" && os.Args[len(os.Args) - 1] != "quilometros"{
		fmt.Println("Ultimo argumento deve ser a unidade: 'celsius' ou 'quilometros'")
		os.Exit(1)
	}else{
		if os.Args[len(os.Args) - 1] == "celsius"{
			unidadeFinal = "fahrenheit"
		}else{
			unidadeFinal = "milhas"
		}
	}

	valores := os.Args[1 : len(os.Args)-1] 
	
	for i, v := range valores{
		valor, erro := strconv.ParseFloat(v, 64)
		if erro != nil{
			fmt.Println("O valor de entrada '%s' [%d] está incorreto", v, i)
			os.Exit(1)
		}

		if unidadeFinal == "fahrenheit"{
			valor = valor*1.8 + 32
			fmt.Printf("%s celsius = %.2f %s\n", v, valor, unidadeFinal)
		}else{
			valor = valor/1.60934
			fmt.Printf("%s quilometros = %.2f %s\n", v, valor, unidadeFinal)
		}
	}
}