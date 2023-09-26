package main

import (
	"fmt"
	"sync"
)

// Функция calculateSquare принимает число, вычисляет его квадрат
// и отправляет результат в канал resultChannel. Она также
// уменьшает счетчик ожидания в sync.WaitGroup после завершения.
func calculateSquare(wg *sync.WaitGroup, number int, resultChannel chan int) {
	defer wg.Done() // Уменьшаем счетчик ожидания при выходе из функции
	square := number * number
	resultChannel <- square // Отправляем результат в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	resultChannel := make(chan int, len(numbers)) // Создаем буферизованный канал для результатов
	var wg sync.WaitGroup                         // Создаем WaitGroup для ожидания завершения горутин

	// Запускаем горутины для вычисления квадратов чисел
	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик ожидания
		go calculateSquare(&wg, number, resultChannel)
	}

	// Горутина для ожидания завершения всех горутин и закрытия канала
	go func() {
		wg.Wait()            // Ожидаем завершения всех горутин
		close(resultChannel) // Закрываем канал после завершения всех горутин
	}()

	// Читаем результаты из канала и выводим их на стандартный вывод
	for square := range resultChannel {
		fmt.Println(square)
	}
}
