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
	testCount, _ := strconv.Atoi(scanner.Text())

	for range testCount {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		a := make([]int, n)
		for i := range a {
			scanner.Scan()
			a[i], _ = strconv.Atoi(scanner.Text())
		}

		b := make([]int, n)
		for i := range b {
			scanner.Scan()
			b[i], _ = strconv.Atoi(scanner.Text())
		}

		if solve(a, b) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// solve определяет, можно ли преобразовать массив a в массив b.
//
// Ключевое наблюдение (из разбора задачи): любая последовательность операций
// приводит к массиву, который является перестановкой либо исходного массива a,
// либо массива, полученного из a ровно одной операцией.
//
// Одна операция с выбранным элементом a_i меняет общий XOR массива на a_i
// (так как n чётно, то n-1 нечётно): S' = S XOR a_i. Значит необходимый
// элемент равен x = S_a XOR S_b.
func solve(a []int, b []int) bool {
	sa := 0
	for _, v := range a {
		sa ^= v
	}
	sb := 0
	for _, v := range b {
		sb ^= v
	}

	sortedA := slices.Clone(a)
	slices.Sort(sortedA)
	sortedB := slices.Clone(b)
	slices.Sort(sortedB)

	// Случай 0 операций: b является перестановкой a.
	if slices.Equal(sortedA, sortedB) {
		return true
	}

	// Иначе b обязано быть перестановкой результата ровно одной операции.
	x := sa ^ sb

	idx := -1
	for i, v := range a {
		if v == x {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false
	}

	// Применяем операцию к индексу idx: XOR элемента x во все остальные.
	c := slices.Clone(a)
	for i := range c {
		if i != idx {
			c[i] ^= x
		}
	}
	slices.Sort(c)

	return slices.Equal(c, sortedB)
}
