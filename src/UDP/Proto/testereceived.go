package main

import "fmt"
import "UDP/Proto/Drivers"

func main(){
	Sender := communication.CommandSender()
	Sender.SetPort(3124)
	var erro error = Sender.CreateSocket()

	if erro != nil{
		fmt.Println("[Erro Sender] ", erro)
	}

	Received := communication.CommandReceiver()
	Received.SetPort(3123)
	erro = Received.CreateSocket()

	if erro != nil{
		fmt.Println("[Erro Receiver] ", erro)
	}

	for{
		fmt.Println("Esperando mensagem...")
		msg := string(Received.Received())
		fmt.Println(msg)
		Sender.Send([]byte(msg))
	}
}