package main

import "fmt"
import "UDP/Proto/Drivers"
import "bufio"
import "os"

func main(){
	Sender := communication.CommandSender()
	Sender.SetPort(3123)
	//Sender.ChangeMode("tcp")
	var erro error = Sender.CreateSocket()

	if erro != nil{
		fmt.Println("[Erro Sender] ", erro)
	}

	defer Sender.Close()

	for{
		fmt.Printf("Digite algo para mandar: ")

		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		b := []byte(msg)
		Sender.Send(b[0:len(b)-1])


		rec, er := Sender.Received()
		if er == nil{
			fmt.Println("[ECO] ", string(rec))
		}else{
			fmt.Println("[Erro] ", er)
		}
	}
}