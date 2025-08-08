package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	// listen on port 42069
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	for {
		// waiting to accept connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		fmt.Println("connection accepted")
		msgs := getLinesChannel(conn)
		for msg := range msgs {
			fmt.Printf("%s\n", msg)
		}
		fmt.Println("connection closed")

	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	var curr_line string
	msg_ch := make(chan string)

	go func() {
		defer f.Close()
		for {
			// reading 8 bytes at a time
			bytes := make([]byte, 8)
			n, err := f.Read(bytes)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if len(curr_line) > 0 {
						msg_ch <- curr_line
					}
					close(msg_ch)
					break
				}
				fmt.Print(err.Error())
				close(msg_ch)
				break
			}
			//split on newlines
			parts := strings.Split(string(bytes[:n]), "\n")

			for i := 0; i < len(parts)-1; i++ {
				curr_line = curr_line + parts[i]
				msg_ch <- curr_line
				curr_line = ""
			}
			curr_line = curr_line + parts[len(parts)-1]
		}
	}()
	return msg_ch
}
