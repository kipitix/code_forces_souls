package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	caseCount, _ := strconv.Atoi(scanner.Text())

	for range caseCount {

		scanner.Scan()
		numCount, _ := strconv.Atoi(scanner.Text())

		arrayB := make([]int, numCount)
		arrayBSum := 0

		for numI := 0; numI < numCount; numI++ {

			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			arrayB[numI] = num
			arrayBSum += num
		}

		result := solve(arrayB, arrayBSum)
		if result == nil {
			fmt.Println(-1)
		} else {
			for i := 0; i < len(result)-1; i++ {
				fmt.Printf("%d ", result[i])
			}
			fmt.Println(result[len(result)-1])
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

}

func solve(arrayB []int, arrayBSum int) []int {
	n := len(arrayB)

	if arrayBSum < 1 {
		return nil
	}

	sortedVals := append([]int(nil), arrayB...)
	slices.Sort(sortedVals)

	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}

	result := make([]int, n)
	sum := 0
	for k := 0; k < n; k++ {
		threshold := -sum
		idx := sort.Search(n, func(i int) bool { return sortedVals[i] > threshold })
		idx = find(idx)

		sum += sortedVals[idx]
		result[k] = sum
		parent[idx] = idx + 1
	}

	return result
}
