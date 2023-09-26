package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем два канала для обмена данными между горутинами.
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Используем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Горутина для чтения чисел из массива и отправки их в первый канал.
	wg.Add(1)
	go func() {
		defer close(inputChannel) // Закрываем канал после завершения чтения.

		numbers := []int{1, 2, 3, 4, 5} // Пример чисел из массива.
		for _, num := range numbers {
			inputChannel <- num
		}
		wg.Done()
	}()

	// Горутина для умножения чисел на 2 и отправки результатов во второй канал.
	wg.Add(1)
	go func() {
		defer close(outputChannel) // Закрываем канал после завершения обработки.

		for x := range inputChannel {
			result := x * 2
			outputChannel <- result
		}
		wg.Done()
	}()

	// Горутина для вывода данных из второго канала в stdout.
	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range outputChannel {
			fmt.Println(result)
		}
	}()

	// Ожидаем завершения всех горутин.
	wg.Wait()
}
