package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result := solution(line)
		fmt.Println(result)
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
}

func solution(input string) string {
	w, _ := strconv.Atoi(input)
	// Число должно быть четным и больше 2
	// Больше 2 потому что 2 разделится на 2 нечетные части 1 и 1
	if w%2 == 0 && w > 2 {
		return "YES"
	}
	return "NO"
}
