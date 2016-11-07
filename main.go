package main

import (
	"fmt"
	"os"

	"github.com/armon/go-socks5"
)

func main() {
	var IPAddress, Port string

	if len(os.Args) == 2 {
		IPAddress = "localhost"
		Port = os.Args[1]
	} else if len(os.Args) == 3 {
		IPAddress = os.Args[1]
		Port = os.Args[2]
	} else {
		fmt.Println("syntax: ")
		fmt.Println("  <IPAddress> <port>  Binds to the specified IP and port")
		fmt.Println("  <port>              Binds to localhost and the specified port")
		return
	}

	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp", IPAddress+":"+Port); err != nil {
		panic(err)
	}
}
