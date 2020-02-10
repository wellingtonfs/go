package main

import "fmt"
import "UDP/Proto/Drivers"

func main(){
	Sender := communication.CommandSender()
	Sender.SetPort(3123)
	var erro error = Sender.CreateSocket()

	if erro != nil{
		fmt.Println("[Erro Sender] ", erro)
	}

	Received := communication.CommandReceiver()
	Received.SetPort(3124)
	erro = Received.CreateSocket()

	if erro != nil{
		fmt.Println("[Erro Receiver] ", erro)
	}

	var msg string
	for{
		fmt.Printf("Digite algo para mandar: ")
		fmt.Scan(&msg)

		Sender.Send([]byte(msg))

		fmt.Println("Esperando resposta...")
		fmt.Println(string(Received.Received()))
	}
}