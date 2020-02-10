package communication

import "net"
import "strconv" 

type AddressSend struct{
	ip string
	port int
	mode string
	cnn net.Conn
}

type AddressRec struct{
	ip string
	port int
	mode string
	cnnUDP net.PacketConn
}

type Socket interface{
	SetAddr(newAddr string)
	SetPort(newPort int)
	GetAddr() string
	GetPort() int
	CreateSocket() error
	Close() error
	Send(bytes []byte)
	Received() []byte
}

//function init
func CommandSender() Socket{
	var conn Socket
	conn = &AddressSend{"127.0.0.1", 5000, "udp", nil}
	return conn
}

func CommandReceiver() Socket{
	var conn Socket
	conn = &AddressRec{"127.0.0.1", 5000, "udp", nil}
	return conn
}

//getters and setters for address
func (a *AddressSend) SetAddr(newAddr string){
	a.ip = newAddr
	if newAddr == "localhost"{
		a.ip = "127.0.0.1"	
	}
}

func (a *AddressRec) SetAddr(newAddr string){
	a.ip = newAddr
	if newAddr == "localhost"{
		a.ip = "127.0.0.1"	
	}
}

func (a *AddressSend) SetPort(newPort int){
	a.port = newPort
}

func (a *AddressRec) SetPort(newPort int){
	a.port = newPort
}

func (a AddressSend) GetAddr() string{
	return a.ip
}

func (a AddressRec) GetAddr() string{
	return a.ip
}

func (a AddressSend) GetPort() int{
	return a.port
}

func (a AddressRec) GetPort() int{
	return a.port
}

//create sockets for sender and receiver
func (addr *AddressSend) CreateSocket() error{
	var err error
	addr.cnn, err = net.Dial(addr.mode, addr.ip+":"+strconv.Itoa(addr.port))
	return err
}

func (addr *AddressRec) CreateSocket() error{
	var err error
	addr.cnnUDP, err = net.ListenPacket(addr.mode, addr.ip+":"+strconv.Itoa(addr.port))
	return err
}

func (addr AddressSend) Close() error{
	return addr.cnn.Close()
}

func (addr AddressRec) Close() error{
	addr.cnnUDP.Close()
	return addr.cnnUDP.Close()
}

func (addr AddressSend) Send(bytes []byte){
	addr.cnn.Write(bytes)
}

func (addr AddressRec) Received() []byte{
	buffer := make([]byte, 1024)
	var n int
	n, _, _ = addr.cnnUDP.ReadFrom(buffer)
	return buffer[0:n]
}

//methods implemented but not used (because senders do not receive bytes and receivers do not send bytes)
func (addr AddressSend) Received() []byte{return []byte(nil)}
func (addr AddressRec) Send(bytes []byte){}
