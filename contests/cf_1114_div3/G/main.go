package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// Бинарная min-куча на срезе int без boxing.
type IntHeap []int

func (h *IntHeap) Push(x int) {
	*h = append(*h, x)
	i := len(*h) - 1
	arr := *h
	for i > 0 {
		parent := (i - 1) / 2
		if arr[parent] <= arr[i] {
			break
		}
		arr[parent], arr[i] = arr[i], arr[parent]
		i = parent
	}
}

func (h *IntHeap) Pop() int {
	arr := *h
	top := arr[0]
	n := len(arr)
	arr[0] = arr[n-1]
	arr = arr[:n-1]
	*h = arr
	i := 0
	for {
		left := 2*i + 1
		if left >= len(arr) {
			break
		}
		smallest := left
		right := left + 1
		if right < len(arr) && arr[right] < arr[left] {
			smallest = right
		}
		if arr[i] <= arr[smallest] {
			break
		}
		arr[i], arr[smallest] = arr[smallest], arr[i]
		i = smallest
	}
	return top
}

// Слияние двух куч по принципу "маленькое в большое".
func mergeHeaps(dst, src *IntHeap) *IntHeap {
	if dst == nil {
		return src
	}
	if src == nil {
		return dst
	}
	if len(*dst) < len(*src) {
		dst, src = src, dst
	}
	for _, v := range *src {
		dst.Push(v)
	}
	return dst
}

func solve(n int, a []int, parent []int) []int {
	// parent[i] - 0-индексированный родитель узла i (для i>=1), parent[0] не используется.
	heaps := make([]*IntHeap, n)
	spares := make([]int, 0, n)

	for i := n - 1; i >= 1; i-- {
		finalize(i, a, heaps, &spares)
		p := parent[i]
		heaps[p] = mergeHeaps(heaps[p], heaps[i])
		heaps[i] = nil
	}
	finalize(0, a, heaps, &spares)

	root := heaps[0]
	leafCount := len(*root)

	sum := 0
	for _, v := range *root {
		sum += v
	}

	sort.Sort(sort.Reverse(sort.IntSlice(spares)))

	result := make([]int, n)
	for k := 1; k < leafCount; k++ {
		result[k-1] = -1
	}

	acc := sum
	result[leafCount-1] = sum
	for k := leafCount + 1; k <= n; k++ {
		acc += spares[k-leafCount-1]
		result[k-1] = acc
	}

	return result
}

// Обрабатывает узел i: если это лист (heaps[i] == nil), создаёт кучу из одного
// элемента. Иначе извлекает минимум из объединённой кучи потомков, сравнивает
// со значением узла и одно из двух значений откладывает как "запасное".
func finalize(i int, a []int, heaps []*IntHeap, spares *[]int) {
	if heaps[i] == nil {
		h := IntHeap{a[i]}
		heaps[i] = &h
		return
	}
	h := heaps[i]
	x := h.Pop()
	if a[i] > x {
		h.Push(a[i])
		*spares = append(*spares, x)
	} else {
		h.Push(x)
		*spares = append(*spares, a[i])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt := func() int {
		n := 0
		c, _ := reader.ReadByte()
		for c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			c, _ = reader.ReadByte()
		}
		neg := false
		if c == '-' {
			neg = true
			c, _ = reader.ReadByte()
		}
		for c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
			c, _ = reader.ReadByte()
		}
		if neg {
			n = -n
		}
		return n
	}

	t := readInt()
	for ; t > 0; t-- {
		n := readInt()
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = readInt()
		}
		parent := make([]int, n)
		for i := 1; i < n; i++ {
			parent[i] = readInt() - 1
		}

		result := solve(n, a, parent)

		buf := make([]byte, 0, n*7)
		for i, v := range result {
			if i > 0 {
				buf = append(buf, ' ')
			}
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		buf = append(buf, '\n')
		writer.Write(buf)
	}
}
