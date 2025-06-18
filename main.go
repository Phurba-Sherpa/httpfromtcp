package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 8, 8)
	line := ""

	for {
		count, err := file.Read(buffer)

		if err != nil {
			if line != "" {
				fmt.Printf("read: %s\n", line)
				line = ""
			}

			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}

		data := string(buffer[:count])
		parts := strings.Split(data, "\n")

		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", line, parts[i])
			line = ""
		}
		line += parts[len(parts)-1]
	}

}
