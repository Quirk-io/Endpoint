
package stun

import (
	"net"
	"log"
	"encoding/json"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
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
	Msgtype string 
	PrivateEndpoint string
}

func AES_encrypt(msg string, key string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	enc_msg := gcm.Seal(nonce, nonce, []byte(msg), nil)
	return string(enc_msg)
}

func AES_decrypt(enc_msg string, key string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal(err)
	}

	nonceSize := gcm.NonceSize()
	if len(enc_msg) < nonceSize {
		log.Fatal(err)
	}

	nonce, enc_msg := enc_msg[:nonceSize], enc_msg[nonceSize:]
	
	msg, err := gcm.Open(nil, []byte(nonce), []byte(enc_msg), nil)
    if err != nil {
        log.Fatal(err)
    }

	return string(msg)
}

func GetPrivateIp() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr).IP
    return localAddr.String()
}

func Reg() RegMsg{ //Register msg
	privendpoint := Endpoint{GetPrivateIp(), "1691"}
	jsonified_privendpoint, _ := json.Marshal(privendpoint)

	regmsg := RegMsg{"reg", string(jsonified_privendpoint)}

	return regmsg
}

func ImportPrivateEndpoint(jsonified_privendpoint string) Endpoint{ //From json to Endpoint
	var endpoint Endpoint
	json.Unmarshal([]byte(jsonified_privendpoint), &endpoint)

	return endpoint
}

func Kenc_Regmsg(AES_key string) string{ //Encrypting Regmsg with AES_key
	jsonified_regmsg, _ := json.Marshal(Reg())
	kenc_regmsg := base64.StdEncoding.EncodeToString([]byte(AES_encrypt(string(jsonified_regmsg), AES_key)))

	return kenc_regmsg
}

func Dkenc_Regmsg(AES_key string, kenc_regmsg string) RegMsg{
	b64dec_regmsg, _ := base64.StdEncoding.DecodeString(kenc_regmsg)
	jsonified_regmsg := AES_decrypt(string(b64dec_regmsg), AES_key)

	var regmsg RegMsg
	json.Unmarshal([]byte(jsonified_regmsg), &regmsg)

	return regmsg
}

func Kenc_Endpoints(AES_key string, endpoints Endpoints) string{
	jsonified_endpoints, _ := json.Marshal(endpoints)
	kenc_endpoints := base64.StdEncoding.EncodeToString([]byte(AES_encrypt(string(jsonified_endpoints), AES_key)))

	return kenc_endpoints
}

func Dkenc_Endpoints(AES_key string, kenc_endpoints string) Endpoints{
	b64dec_endpoints, _ := base64.StdEncoding.DecodeString(kenc_endpoints)
	jsonified_endpoints := AES_decrypt(string(b64dec_endpoints), AES_key)

	var endpoints Endpoints
	json.Unmarshal([]byte(jsonified_endpoints), &endpoints)

	return endpoints
}
