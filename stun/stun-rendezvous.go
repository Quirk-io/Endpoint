
package main

import ("log"
	"net"
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
		
		n, public_addr, read_err := srv.ReadFromUDP(buffer)
		if read_err != nil {
			log.Fatal(read_err)
		}
		
		kenc_regmsg := buffer[:n]
		regmsg := Dkenc_Regmsg(AES_key, string(kenc_regmsg))

		private_endpoint := ImportPrivateEndpoint(regmsg.PrivateEndpoint)
		public_endpoint := Endpoint{public_addr.IP.String(), public_addr.Port}

		endpoints := Endpoints{public_endpoint, private_endpoint}
		
  	}

}

