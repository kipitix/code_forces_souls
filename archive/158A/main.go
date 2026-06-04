package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(nextRoundPlayersCount(os.Stdin))
}

func nextRoundPlayersCount(reader io.Reader) int {
	result := 0

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	lastScore := 0
	for i := 0; i < k; i++ {
		if scanner.Scan() {
			if x, _ := strconv.Atoi(scanner.Text()); x > 0 {
				result++
				lastScore = x
			}
		}
	}

	if lastScore == 0 {
		return result
	}

	for i := k; i < n; i++ {
		if scanner.Scan() {
			if x, _ := strconv.Atoi(scanner.Text()); x == lastScore {
				result++
			} else {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
