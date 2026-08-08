package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	const maxCapacity = 300 * 1024 // 300 КБ
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	caseCount, _ := strconv.Atoi(scanner.Text())

	for _ = range caseCount {
		scanner.Scan()
		// charCount, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		originalString := scanner.Text()
		compactedLength, isCase1, isCase2 := compactedStringLength(originalString)
		if isCase2 {
			compactedLength -= 2
		} else if isCase1 {
			compactedLength--
		}
		fmt.Println(compactedLength)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// Вычисляется длина сжатой строки.
// На выходе:
// 1. Длина.
// 2. Признак того, что была ситуация, что была последовательность из одного символа
// 3. Признак того, что была ситуация, что шла последовательность символов,
// потом 1 символ и потом снова последовательность из таких же символов.
func compactedStringLength(str string) (int, bool, bool) {

	result := 0
	prevChar := '\x00'
	isCase1 := false
	isCase2 := false
	sequenceLen := 0
	prevSeqChar := '\x00'
	prevPrevSeqChar := '\x00'

	for _, char := range str {
		if prevChar != char {

			result++

			prevPrevSeqChar = prevSeqChar
			prevSeqChar = prevChar
			if char == prevPrevSeqChar && sequenceLen == 1 {
				isCase2 = true
			}

			if sequenceLen == 1 && prevPrevSeqChar != '\x00' {
				isCase1 = true
			}
			sequenceLen = 0
		}

		sequenceLen++
		prevChar = char
	}

	return result, isCase1, isCase2
}
