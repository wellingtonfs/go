package main

import(
	"fmt"
	"github.com/golang/protobuf/proto"
	"UDP/Proto/Equipe"
	"net"
)

//ou *net.UDPConn
func Enviar(nome, nomeL, cor string, qtd int64, cnn net.Conn) {
	eq := &equipe.Equipe{
		NomeEquipe: nome,
		NomeLider: nomeL,
		NumComponentes: qtd,
		CorEquipe: cor}

	binario, erro := proto.Marshal(eq)

	if erro != nil{
		fmt.Errorf("Erro: %v", erro)
	}

	cnn.Write(binario)
}

/*
func Resposta(cnn *net.UDPConn) bool{
	fmt.Printf("\nEnviando...... ")

	buffer := make([]byte, 1024)

	nBytes, _, _ := cnn.ReadFrom(buffer)

	if string(buffer[0:nBytes]) == "ok"{
		return true
	}
	return false
}*/

func main(){
	//Cnn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{127,0,0,1}, Port:5000, Zone:""})
	Cnn, _ := net.Dial("udp", "127.0.0.1:5000")

	defer Cnn.Close()

	var nome, nomeL, cor string
	var qtd int64

	for{
		fmt.Printf("Digite o nome da equipe: ")
		fmt.Scan(&nome)
		fmt.Printf("Digite o nome do lider da equipe: ")
		fmt.Scan(&nomeL)
		fmt.Printf("Quantidade de pessoas: ")
		fmt.Scan(&qtd)
		fmt.Printf("Cor da equipe: ")
		fmt.Scan(&cor)

		Enviar(nome, nomeL, cor, qtd, Cnn)
		/*if Resposta(Cnn){
			fmt.Println("OK\n")
		}else{
			fmt.Println("Erro!\n")
		}*/
	}
}