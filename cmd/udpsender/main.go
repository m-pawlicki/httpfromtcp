package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	u, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conn, err := net.DialUDP("udp", nil, u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	rdr := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		str, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
