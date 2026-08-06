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
	caseCount, _ := strconv.Atoi(scanner.Text())
	for caseI := 0; caseI < caseCount; caseI++ {
		var startCoinsCount [3]int
		scanner.Scan()
		startCoinsCount[0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		startCoinsCount[1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		startCoinsCount[2], _ = strconv.Atoi(scanner.Text())

		fmt.Println(solveCase(startCoinsCount[:]))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func solveCase(coinsCount []int) int {

	roundsCount := 0

	for coinsCount[0] != coinsCount[1] && coinsCount[0] != coinsCount[2] && coinsCount[1] != coinsCount[2] {

		slices.Sort(coinsCount)

		diffHiMid := coinsCount[2] - coinsCount[1]
		diffLoMid := coinsCount[1] - coinsCount[0]

		optimalDiff := diffHiMid
		if optimalDiff > diffLoMid {
			optimalDiff = diffLoMid
		}

		coinsCount[2] -= optimalDiff
		coinsCount[0] += optimalDiff
		roundsCount += optimalDiff
	}

	return roundsCount
}
