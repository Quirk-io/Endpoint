
package main

import ("log"
	"net"
	"fmt"
	"encoding/json"
	qpeer "github.com/Quirk-io/go-qPeer/qpeer"
)

type Endpoints struct{
	PublicEndpoint Endpoint
	PrivateEndpoint Endpoint
}

type Endpoint struct{
	Ip string
	Port string
}

type RegMsg struct{
	Msgtype string `json:"msgtype"`
	PrivateEndpoint string `json:"privendpoint"`

}

const (
	signal_ip = ""
	signal_port = ""
	AES_key = ""
)

func GetPrivateIp() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr).IP
    return string(localAddr)
}

func Reg() RegMsg{
	privendpoint := Endpoint{GetPrivateIp(), "1691"}
	jsonified_privendpoint, _ := json.Marshal(privendpoint)

	reg_msg := RegMsg{"reg", string(jsonified_privendpoint)}

	return reg_msg
}


func Udp(AES_key string) (*net.UDPConn, Endpoints) {
	signalsrv, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", signal_ip, signal_port))
	localaddr, _ := net.ResolveUDPAddr("udp", ":1691")

	var err error
	l, err := net.ListenUDP("udp", localaddr)
	if err != nil{
		log.Fatal(err)
	}
	defer l.Close()
	
	jsonified_regmsg, _ := json.Marshal(Reg())
	_, err = l.WriteToUDP([]byte(qpeer.AES_encrypt(string(jsonified_regmsg), AES_key)), signalsrv)
	if err != nil{
		log.Fatal(err)
	}

	buffer := make([]byte, 2048)
		
	n, read_err := l.Read(buffer)
	if read_err != nil {
		log.Fatal(read_err)
	}
	
	recvd := buffer[:n]

	var endpoints Endpoints
	json.Unmarshal([]byte(qpeer.AES_decrypt(string(recvd), AES_key)), &endpoints)

	return l, endpoints
}