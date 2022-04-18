package upnp

import (
	upnp "github.com/jcuga/go-upnp"
	stun "github.com/quirkio/Endpoint/stun"
	"log"
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