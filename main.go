package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 8, 8)

	for {
		count, err := file.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}

		data := string(buffer[:count])
		fmt.Printf("read: %s\n", data)

	}

}
