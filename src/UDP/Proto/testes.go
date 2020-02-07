package main

import "fmt"

type Address struct{
	ip string
	port int
}

type CommandSender struct{
	Addr Address
}

type CommandReceiver struct{
	Addr Address
}

//getters and setters for address and port
func (s *CommandSender) SetAddr(newAddr string){
	s.Addr.ip = newAddr
}

func (s *CommandReceiver) SetAddr(newAddr string){
	s.Addr.ip = newAddr
}

func (s CommandSender) GetAddr() string{
	s.Addr.defaultValues()
	return s.Addr.ip
}

func (s CommandReceiver) GetAddr() string{
	s.Addr.defaultValues()
	return s.Addr.ip
}

func (s *CommandSender) SetPort(newPort int){
	s.Addr.port = newPort
}

func (s *CommandReceiver) SetPort(newPort int){
	s.Addr.port = newPort
}

func (s CommandSender) GetPort() int{
	s.Addr.defaultValues()
	return s.Addr.port
}

func (s CommandReceiver) GetPort() int{
	s.Addr.defaultValues()
	return s.Addr.port
}

//setting default values
func (addr *Address) defaultValues(){
	if addr.ip == ""{
		addr.ip = "127.0.0.1"
	}
	if addr.port == 0{
		addr.port = 5000
	}
}

func main(){
	var CS CommandSender
	//var CR CommandReceiver

	fmt.Println(CS.GetAddr(), CS.GetPort())

	var CR CommandReceiver

	fmt.Println(CR.GetAddr(), CR.GetPort())
}