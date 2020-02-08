package main

import "fmt"
import "UDP/Proto/Drivers"

func main(){
	var msg string
	fmt.Scan(&msg)

	Sender := communication.CommandSender()
	Sender.SetPort(3123)
	erro := Sender.createSocket()

	if erro != nil{
		fmt.Println()//terminar aqq
	}

	Received := communication.CommandReceiver()
	Received.SetPort(3124)
	Received.createSocket()

	fmt.Println()
}