package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
            {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(str)
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)

}

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}
