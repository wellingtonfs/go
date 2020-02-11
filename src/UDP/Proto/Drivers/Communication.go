package communication

import "net"
import "strconv"
import "errors"
import "io"

const(
	ipDefault = "127.0.0.1"
	portDefault = 5000
	modeDefault = "udp"
)

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
	cnnTCP net.Listener
	cnnClient net.Conn
}

type Socket interface{
	SetAddr(newAddr string)
	SetPort(newPort int)
	GetAddr() string
	GetPort() int
	ChangeMode(newMode string) error
	CreateSocket() error
	AcceptConnection() error
	Send(bytes []byte)
	Received() ([]byte, error)
	Close()
}

//function init
func CommandSender() Socket{
	var conn Socket
	conn = &AddressSend{ipDefault, portDefault, modeDefault, nil}
	return conn
}

func CommandReceiver() Socket{
	var conn Socket
	conn = &AddressRec{ipDefault, portDefault, modeDefault, nil, nil, nil}
	return conn
}

//getters and setters for address
func (a *AddressSend) SetAddr(newAddr string){
	a.ip = newAddr
}

func (a *AddressRec) SetAddr(newAddr string){
	a.ip = newAddr
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
	if addr.ip == "localhost"{
		addr.ip = "127.0.0.1"	
	}
	addr.cnn, err = net.Dial(addr.mode, addr.ip+":"+strconv.Itoa(addr.port))
	return err
}

func (addr *AddressRec) CreateSocket() error{
	var err error
	if addr.ip == "localhost"{
		addr.ip = "127.0.0.1"	
	}
	if addr.mode == "udp"{
		addr.cnnUDP, err = net.ListenPacket(addr.mode, addr.ip+":"+strconv.Itoa(addr.port))
	}else{
		addr.cnnTCP, err = net.Listen(addr.mode, addr.ip+":"+strconv.Itoa(addr.port))
	}
	return err
}

func (addr *AddressSend) Close(){
	if addr.cnn != nil{
		addr.cnn.Close()
		addr.cnn = nil
	}
}

func (addr *AddressRec) Close(){
	if addr.cnnUDP != nil{
		addr.cnnUDP.Close()
	}
	if addr.cnnTCP != nil{
		addr.cnnTCP.Close()
	}
	if addr.cnnClient != nil{
		addr.cnnClient.Close()
	}
	addr.cnnUDP = nil
	addr.cnnTCP = nil
	addr.cnnClient = nil
}

func (addr *AddressSend) ChangeMode(newMode string) error{
	if addr.cnn == nil{
		if newMode == "udp" || newMode == "tcp"{
			addr.mode = newMode
		}
		return nil
	}else{
		return errors.New("[Error] Socket still open!")
	}
}

func (addr *AddressRec) ChangeMode(newMode string) error{
	if addr.cnnUDP == nil && addr.cnnTCP == nil{
		if newMode == "udp" || newMode == "tcp"{
			addr.mode = newMode
		}
		return nil
	}else{
		return errors.New("[Error] Some Socket still open!")
	}
}

func (addr AddressSend) Send(bytes []byte){
	addr.cnn.Write(bytes)
}

func (addr AddressRec) Received() ([]byte, error){
	buffer := make([]byte, 1024)
	var n int
	var err error
	if addr.mode == "udp"{
		n, _, err = addr.cnnUDP.ReadFrom(buffer)
		return buffer[0:n], err

	}else if addr.cnnClient != nil{
		n, err = addr.cnnClient.Read(buffer)

		if err == io.EOF{
			addr.cnnClient = nil
			return nil, errors.New("Cliente encerrou a conexao!")
		}

	}
	return buffer[0:n], err
}

//funções usadas somente no protocolo tcp

//Metodo de aceitar novas conexões
func (addr *AddressRec) AcceptConnection() error{
	if addr.mode == "udp"{return errors.New("Socket UDP nao tem o metodo AcceptConnection!")}

	if addr.cnnClient != nil{
		addr.cnnClient.Close()
	}

	var err error
	addr.cnnClient, err =  addr.cnnTCP.Accept()
	return err
}

//metodo para enviar dados do recepitor tcp
func (addr AddressRec) Send(bytes []byte){
	if addr.mode == "tcp" && addr.cnnClient != nil{
		addr.cnnClient.Write(bytes)
	}
}

//metodo para receber dados no remetente
func (addr AddressSend) Received() ([]byte, error){
	if addr.mode == "tcp" && addr.cnn != nil{
		buffer := make([]byte, 1024)
		n, err := addr.cnn.Read(buffer)

		return buffer[0:n], err
	}
	return nil, errors.New("[Erro] Metodo funciona somente no protocolo TCP e com uma conexao ativa!")
}

//methods implemented but not used (because senders do not receive bytes and receivers do not send bytes)
func (addr AddressSend) AcceptConnection() error{return nil}
