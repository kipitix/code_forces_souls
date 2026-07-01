package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const capitalDistance = 'a' - 'A'

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstChar := true

	for {
		char, err := reader.ReadByte()
		if err == io.EOF || char == '\n' || char == '\r' {
			break
		}

		if firstChar {
			if char >= 'a' && char <= 'z' {
				char -= capitalDistance
			}
			firstChar = false
		}

		fmt.Printf("%c", char)
	}
}
