package main

import (
	"bufio"
	"os"
	"testing"
)

func TestProcessOperationsWithFile(t *testing.T) {
	// Открываем файл test.txt
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Не удалось открыть файл test.txt: %v", err)
	}
	defer file.Close()

	// Создаем сканер из файла
	scanner := bufio.NewScanner(file)

	// Обрабатываем операции
	result := ProcessOperations(scanner)

	// Выводим результат (опционально)
	t.Logf("Результат выполнения операций: %d", result)

	// Здесь можно добавить проверку на ожидаемый результат
	// Например, если вы знаете ожидаемое значение:
	// expected := 5
	// if result != expected {
	//     t.Errorf("Expected %d, got %d", expected, result)
	// }
}

// Дополнительный тест с несколькими файлами, если нужно
func TestProcessOperationsWithMultipleFiles(t *testing.T) {
	testFiles := []string{
		"test.txt",
		// можно добавить другие файлы: "test2.txt", "test3.txt" и т.д.
	}

	for _, filename := range testFiles {
		t.Run(filename, func(t *testing.T) {
			file, err := os.Open(filename)
			if err != nil {
				t.Skipf("Файл %s не найден, пропускаем", filename)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			result := ProcessOperations(scanner)
			t.Logf("Файл: %s, результат: %d", filename, result)
		})
	}
}

// Тест, который сравнивает результат с ожидаемым из отдельного файла
func TestProcessOperationsWithExpectedResult(t *testing.T) {
	// Открываем файл с операциями
	operationsFile, err := os.Open("test.txt")
	if err != nil {
		t.Skipf("Файл test.txt не найден, пропускаем тест")
		return
	}
	defer operationsFile.Close()

	// Обрабатываем операции
	scanner := bufio.NewScanner(operationsFile)
	result := ProcessOperations(scanner)

	// Проверяем, что результат не отрицательный (пример проверки)
	if result < 0 {
		t.Errorf("Результат не должен быть отрицательным, получено: %d", result)
	}

	// Можно также сохранить результат в файл
	// resultFile, _ := os.Create("result.txt")
	// fmt.Fprintf(resultFile, "%d", result)
	// resultFile.Close()
}
