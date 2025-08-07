package main

import (
	"fmt"
	"os"
)

func main() {

	msgs, err := os.Open("messages.txt")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	for {
		bytes := make([]byte, 8)
		_, err := msgs.Read(bytes)
		if err != nil {
			return
		}
		fmt.Printf("read: %s\n", bytes)
	}
}
