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

// Находим пары стопок с минимальной разницей и пытаемся их уровнять.
// Идём от наименьшей разницы к наибольшей и проверяем можно ли решить.
// Для каждой пары:
// Оцениваем разность.
// Если разница чётная, то можно достичь равенства переместив половину разницы на другую стопку.
// Если разница нечётна, то перемещаем целую часть половины разности и ещё одну монету с соседней стопки - там всегда должна быть хотя бы одна.
func solveCase(startCoinsCount []int) int {

	slices.Sort(startCoinsCount)

	diff := startCoinsCount[1] - startCoinsCount[0]
	if diff%2 == 0 {
		return diff / 2
	} else {
		return diff/2 + 1
	}
}
