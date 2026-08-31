package main

const (
	maxBits = 31
)

func main() {

}

func solve(a []int, b []int, size int) bool {
	// Вычисляем массив C, на основании которого будем оценивать достижимость решения
	c := make([]int, size)
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	// Создаём 2 базиса
	// Straight Basis Operator
	sbo := NewXorBasisOperator(maxBits, size+1)
	// Reverse Basis Operator
	rbo := NewXorBasisOperator(maxBits, size+1)
	// Заполняем SBO всеми элементами из A, исключая последний
	for i := 0; i < size-1; i++ {
		sbo.Add(a[i])
	}
	// Проходим в обратном порядке смотря на массив C, используя для оценки массив A
	// и объединённые базисы
	for i := size - 1; i >= 0; i-- {
		// Проверяем значение
		mergedBasis := sbo.Basis().Merge(rbo.Basis())
		if !mergedBasis.Check(c[i]) {
			return false
		}
		// Корректируем базисы
		sbo.Rewind()
		rbo.Add(a[i])
	}
	return true
}

// XOR Базис - хранит базис и позволяет 2 операции с ним
type XorBasis []int

// Проверить возможность составления значения на основании базиса
func (xb XorBasis) Check(value int) bool {
	if value == 0 {
		// Проверка на то, что базис НЕ пустой
		// return !xb.IsZero()
		return true
	}
	for index := len(xb) - 1; index >= 0; index-- {
		if (value>>index)&1 == 1 {
			if xb[index] == 0 {
				return false
			}
			value ^= xb[index]
		}
	}
	return value == 0
}

// Проверка, что базис пустой
func (xb XorBasis) IsZero() bool {
	for _, v := range xb {
		if v != 0 {
			return false
		}
	}
	return true
}

// Объединить 2 базиса
func (xb XorBasis) Merge(other XorBasis) XorBasis {
	// Merged Basis Operator
	mbo := NewXorBasisOperator(len(xb), len(xb)*2+1)
	for _, v := range xb {
		mbo.Add(v)
	}
	for _, v := range other {
		mbo.Add(v)
	}
	return mbo.Basis()
}

// Оператор XOR базиса
// Формирует базис на основании данных подаваемых в него
// Умеет делать шаг назад за счёт хранения истории базиса
type XorBasisOperator struct {
	maxBits           int
	basisStory        []XorBasis
	currentBasisIndex int
}

// Создание нового оператора базиса
func NewXorBasisOperator(maxBits int, storyCapacity int) *XorBasisOperator {
	xb := &XorBasisOperator{
		maxBits:           maxBits,
		basisStory:        make([]XorBasis, storyCapacity),
		currentBasisIndex: 0,
	}
	for i := 0; i < storyCapacity; i++ {
		xb.basisStory[i] = make([]int, maxBits)
	}
	return xb
}

// Текущий базис
func (xbo *XorBasisOperator) Basis() XorBasis {
	return xbo.basisStory[xbo.currentBasisIndex]
}

// Добавить новое значение базиса с сохранением истории
func (xbo *XorBasisOperator) Add(value int) {
	copy(xbo.basisStory[xbo.currentBasisIndex+1], xbo.basisStory[xbo.currentBasisIndex])
	xbo.currentBasisIndex++
	for i := xbo.maxBits - 1; i >= 0; i-- {
		if (value>>i)&1 == 1 {
			if xbo.basisStory[xbo.currentBasisIndex][i] == 0 {
				xbo.basisStory[xbo.currentBasisIndex][i] = value
				return
			}
			value ^= xbo.basisStory[xbo.currentBasisIndex][i]
		}
	}
}

// Отмотать историю на один шаг назад
func (xbo *XorBasisOperator) Rewind() {
	xbo.currentBasisIndex--
}
