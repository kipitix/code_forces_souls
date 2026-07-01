package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	oneCounter := 0
	twoCounter := 0
	threeCounter := 0

	for {
		char, err := reader.ReadByte()
		if err == io.EOF || char == '\n' || char == '\r' {
			break
		}

		switch char {
		case '1':
			oneCounter++

		case '2':
			twoCounter++

		case '3':
			threeCounter++
		}
	}

	isEmptyOutput := true

	for i := 0; i < oneCounter; i++ {
		if isEmptyOutput {
			fmt.Print("1")
			isEmptyOutput = false
		} else {
			fmt.Print("+1")
		}
	}

	for i := 0; i < twoCounter; i++ {
		if isEmptyOutput {
			fmt.Print("2")
			isEmptyOutput = false
		} else {
			fmt.Print("+2")
		}
	}

	for i := 0; i < threeCounter; i++ {
		if isEmptyOutput {
			fmt.Print("3")
			isEmptyOutput = false
		} else {
			fmt.Print("+3")
		}
	}
}
