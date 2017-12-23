package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	var listener net.Listener
	var error error
	var conn net.Conn

	mapUsers = make(map[string]int)

	fmt.Println("Starting the server ...")

	listener, error = net.Listen("tcp", "localhost:50000")
	checkError(error)
	for {
		conn, error = listener.Accept()
		checkError(error)
		go doServerStuff1(conn)
	}
}

func doServerStuff1(conn net.Conn) {
	var buf []byte
	var err error

	for {
		buf = make([]byte, 512)
		_, err = conn.Read(buf)
		checkError(err)
		input := string(buf)
		if strings.Contains(input, "SH") {
			fmt.Println("Server shutting down now.")
			os.Exit(0)
		}
		if strings.Contains(input, "WHO") {
			displayUser()
		}

		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data :-- %v --", string(buf))

	}
}

func displayUser() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error:" + err.Error())
	}
}
