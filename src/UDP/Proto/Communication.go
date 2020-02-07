package communication

import "net"
import "errors"

/*
type Socket interface{
	createSocket error
	close nil
	SetAddr nil
	GetAddr string
	SetPort nil
	GetPort int
}*/

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

type Received struct{
	Addr Address
}

func (s interface{}) SetAddr(newAddr string){
	s.Addr.ip = newAddr
}

func (s interface{}) GetAddr(){
	return s.Addr.ip
}

func (s interface{}) SetPort(newPort int){
	s.Addr.port = newPort
}

func (s interface{}) GetPort(){
	return s.Addr.port
}