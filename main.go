package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	msgs, err := os.Open("messages.txt")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer msgs.Close()

	var curr_line string

	for {
		// reading 8 bytes at a time
		bytes := make([]byte, 8)
		n, err := msgs.Read(bytes)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Print(err.Error())
			break
		}
		//split on newlines
		parts := strings.Split(string(bytes[:n]), "\n")

		for i := 0; i < len(parts)-1; i++ {
			curr_line = curr_line + parts[i]
			fmt.Printf("read: %s\n", curr_line)
			curr_line = ""
		}
		curr_line = curr_line + parts[len(parts)-1]
	}
	if len(curr_line) > 0 {
		fmt.Printf("read: %s\n", curr_line)
	}
}
