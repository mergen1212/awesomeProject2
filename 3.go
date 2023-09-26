package main

import (
	"fmt"
	"sync"
)

func calculateSquareSum(wg *sync.WaitGroup, numbers []int, resultChannel chan int) {
	defer wg.Done()
	squareSum := 0
	for _, number := range numbers {
		squareSum += number * number
	}
	resultChannel <- squareSum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	resultChannel := make(chan int)
	var wg sync.WaitGroup

	// Запускаем несколько горутин для вычисления суммы квадратов чисел
	numWorkers := 4 // Количество горутин
	chunkSize := len(numbers) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == numWorkers-1 {
			end = len(numbers)
		}

		wg.Add(1)
		go calculateSquareSum(&wg, numbers[start:end], resultChannel)
	}

	// Горутина для ожидания завершения всех вычислений
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Собираем результаты из всех горутин и суммируем их
	totalSum := 0
	for squareSum := range resultChannel {
		totalSum += squareSum
	}

	fmt.Printf("Сумма квадратов: %d\n", totalSum)
}
