
package main

import ("log"
	"net"
	"fmt"
	"encoding/json"
	qpeer "github.com/Quirk-io/go-qPeer/qpeer"
)

func Udp_Rendezvous(AES_key string) (*net.UDPConn, Endpoints){
	addr, _ := net.ResolveUDPAddr("udp", ":1691")

	var err error
	srv, err := net.ListenUDP("udp", addr)
	if err != nil{
		log.Fatal(err)
	}
	defer srv.Close()

	for {
    	buffer := make([]byte, 2048)
		
		n, read_err := conn.Read(buffer)
		if read_err != nil {
			log.Fatal(read_err)
		}
		
		regmsg := buffer[:n]
		
  	}
  }

}

