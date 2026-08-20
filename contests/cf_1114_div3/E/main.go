package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

		// Исходный перемешанный массив разностей
		arrayB := make([]int, numCount)
		// Сумма всех элементов исходного массива - равна значению последнего элемента
		arrayBSum := 0
		// arrayBMin := math.MaxInt
		// arrayBMax := math.MinInt

		for numI := 0; numI < numCount; numI++ {

			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			arrayB[numI] = num
			arrayBSum += num
			// if num > arrayBMax {
			// 	arrayBMax = num
			// }
			// if num < arrayBMin {
			// 	arrayBMin = num
			// }
		}

		// fmt.Println(arrayBSum, arrayBMin, arrayBMax, arrayBMax+arrayBMin)

		// diff = Ai - Ai - 1
		// Ai - 1 = Ai - diff

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
	// fmt.Println("---")
	// fmt.Println(arrayB)

	if arrayBSum < 1 {
		return nil
	}

	result := make([]int, len(arrayB))
	slices.Sort(arrayB)

	firstElement := math.MaxInt
	firstElementIndexInArrayB := -1
	for index, value := range arrayB {
		if value < firstElement && value > 0 {
			firstElement = value
			firstElementIndexInArrayB = index
		}
		if value == 1 {
			break
		}
	}

	arrayB = append(arrayB[:firstElementIndexInArrayB], arrayB[firstElementIndexInArrayB+1:]...)

	// fmt.Println(arrayB)
	headIndex := 0
	tailIndex := len(arrayB) - 1

	result[0] = firstElement
	result[len(result)-1] = arrayBSum
	for resultIndex := len(result) - 2; resultIndex > 0; resultIndex-- {
		prevValue := result[resultIndex+1]
		variant1 := prevValue - arrayB[headIndex]
		variant2 := prevValue - arrayB[tailIndex]
		if variant1 < 1 && variant2 < 1 {
			return nil
		}
		if variant1 < 1 {
			result[resultIndex] = variant2
			tailIndex--
			continue
		}
		if variant2 < 1 {
			result[resultIndex] = variant1
			headIndex++
			continue
		}
		if variant1 < variant2 {
			result[resultIndex] = variant1
			headIndex++
			continue
		} else {
			result[resultIndex] = variant2
			tailIndex--
			continue
		}
	}

	return result
}
