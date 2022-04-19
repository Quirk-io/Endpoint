package upnp

import (
	upnp "github.com/jcuga/go-upnp"
	stun "github.com/quirkio/Endpoint/stun"
	"log"
	"strconv"
)

func GetIp() string{
	//setup router
	u, err := upnp.Discover()
	if err != nil{
		log.Fatal(err)
	}

	//get Ip
	ip, err := u.ExternalIP()
	if err != nil{
		log.Fatal(err)
	}

	return ip
}

func OpenPort(port string) stun.Endpoint{
	proto := "TCP"
	ip := GetIp()
	port_int, _ := strconv.Atoi(port)

	u, uErr := upnp.Discover()
	if uErr != nil{
		log.Fatal(uErr)
	}

	fErr := u.Forward(uint16(port_int), "Forwarding req by Endpoint", proto) 
	if fErr != nil{
		log.Fatal(fErr)
	}

	endpoint := stun.Endpoint{ip, port}

	return endpoint
}

func ClosePort(port string){
	
}