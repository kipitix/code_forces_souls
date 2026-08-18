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

	for range caseCount {
		scanner.Scan()
		numCount, _ := strconv.Atoi(scanner.Text())
		// Мапка с данными для дальнейшего восстановления исходного массива
		// Ключ - значение из исходного массива - значение - слайс индексов, содержащих такое значение
		valuesIndexes := make(map[int][]int)
		for numIndex := range numCount {
			scanner.Scan()
			number, _ := strconv.Atoi(scanner.Text())
			indexes, _ := valuesIndexes[number]
			if indexes == nil {
				valuesIndexes[number] = []int{numIndex}
			} else {
				valuesIndexes[number] = append(indexes, numIndex)
			}
		}
		// Решаем
		result := solve(valuesIndexes, numCount)
		// Выводим
		if result == nil {
			fmt.Println(-1)
		} else {
			iLimit := numCount - 1
			for i, value := range result {
				if i != iLimit {
					fmt.Printf("%d ", value)
				} else {
					fmt.Println(value)
				}
			}
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func solve(valuesIndexes map[int][]int, numCount int) []int {
	// Подготовка результата
	result := make([]int, numCount)
	// Получаем из все уникальные значения исходного массива теней
	// Это все ключи входной мапки
	values := []int{}
	for value := range valuesIndexes {
		values = append(values, value)
	}
	// Сортируем по возрастанию
	slices.Sort(values)
	// Обоити все
	// Исключение - если всё заполнено одним значением (это должен быть 0)
	if len(values) < 2 {
		// это должен быть 0 - иначе решения нет
		if values[0] != 0 {
			return nil
		}
		// Генерируем ответ - все 1
		result[0] = 1
		// Exponentially duplicate the data block
		for j := 1; j < len(result); j *= 2 {
			copy(result[j:], result[:j])
		}
		return result
	}
	// Типовой случай
	lastValue := -1
	for valueIndex := 1; valueIndex < len(values); valueIndex++ {
		shadow := values[valueIndex]
		prevShadow := values[valueIndex-1]
		if (shadow-prevShadow)%len(valuesIndexes[prevShadow]) != 0 {
			return nil
		}
		targetValue := (shadow - prevShadow) / len(valuesIndexes[prevShadow])
		if targetValue < lastValue {
			return nil
		}
		lastValue = targetValue
		for _, resultIndex := range valuesIndexes[prevShadow] {
			result[resultIndex] = targetValue
		}
	}
	// Заполняем последнее значение
	lastValue++
	for _, resultIndex := range valuesIndexes[values[len(values)-1]] {
		result[resultIndex] = lastValue
	}

	return result
}
