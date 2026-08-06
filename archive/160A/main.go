package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	coins := make([]int, n)
	allSum := 0

	for i := range n {
		scanner.Scan()
		coin, _ := strconv.Atoi(scanner.Text())

		allSum += coin

		coins[i] = coin
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.Sort(coins)

	targetSum := allSum/2 + 1
	minCoinsCount := 0
	mySum := 0

	for i := n - 1; i >= 0 && mySum < targetSum; i-- {
		mySum += coins[i]
		minCoinsCount++
	}

	fmt.Println(minCoinsCount)
}
