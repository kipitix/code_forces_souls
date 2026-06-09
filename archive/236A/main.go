package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	differentLetters := map[rune]bool{}

	for {
		char, _, err := reader.ReadRune()
		if err != nil || char == '\n' || char == '\r' {
			break
		}
		differentLetters[char] = true
	}

	if len(differentLetters)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
