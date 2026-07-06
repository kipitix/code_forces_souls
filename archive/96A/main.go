package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	oneCount := 0
	zeroCount := 0
	for {
		char, err := reader.ReadByte()
		if err == io.EOF || char == '\n' || char == '\r' {
			break
		}
		if char == '1' {
			oneCount++
			zeroCount = 0
		} else if char == '0' {
			zeroCount++
			oneCount = 0
		}
		if oneCount >= 7 || zeroCount >= 7 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
