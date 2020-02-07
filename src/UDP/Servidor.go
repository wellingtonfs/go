package main

import (
	"fmt"
	"net"
	"log"
	"github.com/golang/protobuf/proto"
	"UDP/Proto/Equipe"
)

func main(){
	Cnn, erro := net.ListenPacket("udp", ":5000")

	if erro != nil{
		log.Fatal(erro)
	}

	defer Cnn.Close()

	for{
		buffer := make([]byte, 1024)

		_, addr, _ := Cnn.ReadFrom(buffer)

		var recebido equipe.Equipe

		err := proto.Unmarshal(buffer, &recebido)

		if err != nil{
			log.Fatal(err)
		}

		fmt.Println("...........IP: ", addr, ":")

		fmt.Println("Nome Equipe:", recebido.GetNomeEquipe())
		fmt.Println("Lider:", recebido.GetNomeLider())
		fmt.Println("Integrantes:", recebido.GetNumComponentes())
		fmt.Println("Cor:", recebido.GetCorEquipe())
		fmt.Println("------------------\n")

		Cnn.WriteTo([]byte("ok"), addr)
	}
}