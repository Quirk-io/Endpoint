
package main

import ("log"
	"net"
	"fmt"
)

func Udp(AES_key, signal_ip, signal_port string) (*net.UDPConn, Endpoints) {
	signalsrv, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", signal_ip, signal_port))
	localaddr, _ := net.ResolveUDPAddr("udp", ":1691")

	var err error
	l, err := net.ListenUDP("udp", localaddr)
	if err != nil{
		log.Fatal(err)
	}
	
	_, err = l.WriteToUDP([]byte(Kenc_Regmsg(AES_key)), signalsrv)
	if err != nil{
		log.Fatal(err)
	}

	buffer := make([]byte, 2048)
		
	n, read_err := l.Read(buffer)
	if read_err != nil {
		log.Fatal(read_err)
	}
	
	recvd := buffer[:n]

	endpoints := Dkenc_Endpoints(AES_key, string(recvd))

	return l, endpoints
}