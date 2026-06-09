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
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		sequence := []int{}
		for j := 0; j < n; j++ {
			scanner.Scan()
			a, _ := strconv.Atoi(scanner.Text())
			inserted := false
			for i, ai := range sequence {
				if a > ai {
					sequence = append(sequence, 0)
					copy(sequence[i+1:], sequence[i:])
					sequence[i] = a
					// sequence = append(sequence[:i], append([]int{a}, sequence[i:]...)...)
					inserted = true
					break
				}
			}
			if !inserted {
				sequence = append(sequence, a)
			}
		}

		if len(sequence) < 2 {
			fmt.Println(-1)
		} else if len(sequence) == 2 {
			fmt.Println(sequence[0], sequence[1])
		} else if checkSequence(sequence) {
			fmt.Println(sequence[0], sequence[1])
		} else {
			fmt.Println(-1)
		}
	}
}

func checkSequence(sequence []int) bool {
	for i := 0; i < len(sequence)-2; i++ {
		if sequence[i]%sequence[i+1] != sequence[i+2] {
			return false
		}
	}
	return true
}
