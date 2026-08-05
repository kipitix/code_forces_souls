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

	fmt.Println(n)

	// Coins values sorted in decreased sequence
	coins := []int{}
	sum := 0

	for scanner.Scan() {
		coin, _ := strconv.Atoi(scanner.Text())
		sum += coin

		//
		insertPosition := len(coins)
		for i, v := range coins {
			if v < coin {
				insertPosition = i
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
