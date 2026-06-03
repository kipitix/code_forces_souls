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
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	totalTasksAmount := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		lineSum := 0
		for _, xStr := range strings.Split(line, "") {
			if xStr == "1" {
				lineSum++
			}
		}
		if lineSum >= 2 {
			totalTasksAmount++
		}
	}

	fmt.Println(totalTasksAmount)

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
}
