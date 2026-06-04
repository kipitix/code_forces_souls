package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func nextRoundPlayersCount(string) {

}
