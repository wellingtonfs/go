package main

import "fmt"
import "UDP/Proto/Drivers"

func main(){
	Received := communication.CommandReceiver()
	Received.SetPort(3123)
	//Received.ChangeMode("tcp")
	var erro error = Received.CreateSocket()

	if erro != nil{
		fmt.Println("[Erro Receiver] ", erro)
	}

	erro = Received.AcceptConnection()

	defer Received.Close()

	if erro != nil{
		fmt.Println("[Erro Accept] ", erro)
	}

	fmt.Println("Esperando mensagens\n")
	for{
		msg, err := Received.Received()
		if err != nil{
			fmt.Println("[ERRO] ", err)
			break
		}
		fmt.Println("[Mensagem] ", string(msg))
		Received.Send(msg)
	}
}