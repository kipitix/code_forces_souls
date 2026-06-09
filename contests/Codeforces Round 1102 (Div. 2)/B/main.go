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
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		r := n % 12
		if r == 10 {
			if n >= 22 {
				fmt.Println(22, n-22)
			} else {
				fmt.Println(-1)
			}
		} else {
			fmt.Println(r, n-r)
		}
	}
}
