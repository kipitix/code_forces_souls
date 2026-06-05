package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const letterABig = 'A'
const letterASmall = 'a'
const letterNewLine = '\n'
const letterDistance = letterASmall - letterABig

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadBytes(letterNewLine)
	line1 = line1[:len(line1)-1]
	charIndexLine1 := 0
	for {
		charLine2, err := reader.ReadByte()
		if err == io.EOF || charLine2 == letterNewLine {
			break
		}
		charLine1 := line1[charIndexLine1]
		charIndexLine1++

		if charLine1 >= letterASmall {
			charLine1 -= letterDistance
		}
		if charLine2 >= letterASmall {
			charLine2 -= letterDistance
		}

		if charLine1 > charLine2 {
			fmt.Println(1)
			return
		} else if charLine1 < charLine2 {
			fmt.Println(-1)
			return
		}
	}

	fmt.Println(0)
}
