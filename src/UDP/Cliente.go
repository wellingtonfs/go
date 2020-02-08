package main

import (
	//"fmt"
	"net"
	"log"
)

func main(){
	Cnn, erro := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{127,0,0,1}, Port:4050, Zone:""})

	if erro != nil{
		log.Fatal(erro)
	}

	defer Cnn.Close()

	Cnn.Write([]byte("Hello World!"))
/*
	buffer := make([]byte, 1024)

	nBytes, addr, _ := Cnn.ReadFrom(buffer)
	
	fmt.Println("Received: ", string(buffer[0:nBytes]), ", From: ", addr)*/
}