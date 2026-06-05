package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	row, column := recognizeCoordinatesOfOne(os.Stdin)
	moves := math.Abs(float64(row-3)) + math.Abs(float64(column-3))
	fmt.Println(moves)
}

func recognizeCoordinatesOfOne(reader io.Reader) (row int, column int) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	position := 0
	for scanner.Scan() {
		if scanner.Text() == "1" {
			break
		}
		position++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	row = position/5 + 1
	column = position%5 + 1
	return
}
