/*
Funcoes:

Cadastrar
Editar
Procurar
Visualizar

*/

package main

import(
	"fmt"
	"os"
	"os/exec"
)

func Menu() int{
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Bem vindos a minha biblioteca haha\n")

	fmt.Println("(1) Cadastrar")
	fmt.Println("(2) Editar")
	fmt.Println("(3) Procurar")
	fmt.Println("(4) Visualizar")
	fmt.Println("(5) Sair")

	var resposta int
	fmt.Scan(&resposta)

	switch resposta{
		case 1, 2, 3, 4, 5:
			return resposta
		default:
			return Menu()
	}
}

func main(){
	Menu()
}