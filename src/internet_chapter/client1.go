package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var conn net.Conn
	var err error
	var inputReader *bufio.Reader
	var input string
	var clientName string

	conn, err = net.Dial("tcp", "localhost:50000")
	checkError1(err)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, err = inputReader.ReadString('\n')
	checkError1(err)
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to shutdown server.")
		input, err = inputReader.ReadString('\n')
		checkError1(err)
		trimmedInput := strings.Trim(input, "\r\n")

		if strings.ToLower(trimmedInput) == "q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput + "\n"))
		checkError1(err)
	}

}

func checkError1(err error) {
	if err != nil {
		panic("Error:" + err.Error())
	}
}
