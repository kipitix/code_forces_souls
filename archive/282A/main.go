package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	result := ProcessOperations(scanner)
	fmt.Println(result)
}

// ProcessOperations обрабатывает операции из сканера и возвращает результат
func ProcessOperations(scanner *bufio.Scanner) int {
	scanner.Scan()
	operationsCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0
	}

	result := 0
	for i := 0; i < operationsCount; i++ {
		if !scanner.Scan() {
			break
		}
		operation := scanner.Text()
		operationParts := strings.Split(operation, "X")
		trueOperation := operationParts[0]
		if trueOperation == "" {
			trueOperation = operationParts[1]
		}
		switch trueOperation {
		case "++":
			result++
		case "--":
			result--
		default:
			panic("Invalid operation")
		}
	}
	return result
}
