package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	wordCountStr := scanner.Text()
	wordCount, _ := strconv.Atoi(wordCountStr)
	for i := 0; i < wordCount; i++ {
		scanner.Scan()
		word := scanner.Text()
		fmt.Println(solution(word))
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
}

func solution(word string) string {
	if len(word) <= 10 {
		return word
	}
	return fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
}
