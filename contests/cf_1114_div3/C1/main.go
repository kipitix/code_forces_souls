package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	scannerBufferSize = 1024 * 1024
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scannerBuffer := make([]byte, scannerBufferSize)
	scanner.Buffer(scannerBuffer, scannerBufferSize)

	scanner.Scan()
	caseCount, _ := strconv.Atoi(scanner.Text())

	for range caseCount {
		scanner.Scan()
		linesLen, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		lineA := scanner.Bytes()
		scanner.Scan()
		lineB := scanner.Bytes()

		if solve(lineA, lineB, linesLen) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// По задаче - можно заменять подстроки
// 100 -> 001
// 001 -> 100
// 110 -> 011
// 011 -> 110
// Можно представить это по-другому:
// по сути мы меняем два символа, разделённые третьим.
// Причём, неважно какой символ по середине (4 комбинации делают возможным и 1 и 0).
// Я буду идти от начала строки к концу и проверять каждый символ.
// Если символы равны, то идём дальше, если нет, то пытаемся его поменять с ближайшим справа.
// Важно какую позицию бита мы сейчас рассматриваем (чётная или нечётная).
// Если текущий символ 0 и позиция чётная, то ищем ближайшую справа 1 тоже на чётной позиции.
// Если текущий символ 1 и позиция чётная, то ищем ближайший 0 тоже на чётной позиции.
// Если текущий символ 0 и позиция НЕ чётная, то ищем ближайшую справа 1 тоже на НЕ чётной позиции.
// Если текущий символ 1 и позиция НЕ чётная, то ищем ближайший 0 тоже на НЕ чётной позиции.
// В данном варианте условия нет необходимости считать количество шагов,
// поэтому просто меняем символы на текущей позиции с ближайшим подходящим.
// Одновременно актуализируем индексы ближайших 0 и 1 для чётной и нечётной позиции.

// Рассмотрим как обновлять индексы, как буд-то нет разделения на чёт/нечет,
// тогда это будет непрерывная последовательность бит.
// Пример:
// State 0:
// point | _ _ _ _ _ _ _
// index 0 2 4 6 8 A C D
// lineA 1 1 0 0 1 1 0 0
// lineB 0 0 1 1 0 0 1 1
// Ближайшие индексы: 0: 4, 1: 2.
// Меняем местами ячейки 0 и 4.
// State 1:
// point _ | _ _ _ _ _ _
// index 0 2 4 6 8 A C D
// lineA 0 1 1 0 1 1 0 0
// lineB 0 0 1 1 0 0 1 1
// Ближайшие индексы: 0: 6, 1: 4.
// Меняем местами ячейки 2 и 6.
// State 2:
// point _ _ | _ _ _ _ _
// index 0 2 4 6 8 A C D
// lineA 0 0 1 1 1 1 0 0
// lineB 0 0 1 1 0 0 1 1

// Обновлять индексы лучше по необходимости, для оптимизации.
// Буду делать это если понятно, что нудно делать обмен.
// Также буду хранить последнее полученное значение ближайшего индекса.

// Попробую свести задачу к двум проходам по чётному и нечётному индексу.
// В таком случае сложность не увеличится.
// Будет проход по всему массиву также 1 раз.

func solve(lineA, lineB []byte, linesLen int) bool {

	for evenOdd := 0; evenOdd < 2; evenOdd++ {

		nearestIndex := [2]int{-1, -1}

		for watchIndex := evenOdd; watchIndex < linesLen; watchIndex += 2 {
			// Если равны - идём к следующему биту
			if lineA[watchIndex] == lineB[watchIndex] {
				continue
			}

			// Нужно менять

			// Преобразуем символ в значение 0 / 1
			var value, targetValue byte
			value = lineA[watchIndex] - '0'
			targetValue = 0
			if value == 0 {
				targetValue = 1
			}

			// Обновляем ближайшее значение для value
			// Если value == 0, то ищем 1 и наоборот
			// Начинаем с последнего значения или с рассматриваемого индекса + 2 (смотря что больше)
			found := false
			for searchNextIndex := max(nearestIndex[targetValue], watchIndex+2); searchNextIndex < linesLen; searchNextIndex += 2 {
				if lineA[searchNextIndex]-'0' == targetValue {
					nearestIndex[targetValue] = searchNextIndex
					found = true
					break
				}
			}
			if !found {
				nearestIndex[targetValue] = linesLen
			}

			// Меняем, если возможно
			if nearestIndex[targetValue] < linesLen {
				lineA[watchIndex], lineA[nearestIndex[targetValue]] = lineA[nearestIndex[targetValue]], lineA[watchIndex]
				continue
			}

			return false
		}
	}

	return true
}
