package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	halfN := (n + 1) / 2
	result := 0
	if i <= halfN {
		result = i*2 - 1
	} else {
		result = (i - halfN) * 2
	}

	fmt.Println(result)
}
