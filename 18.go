package main

import (
	"fmt"
	"sync"
)

// Структура Counter для счетчика.
type Counter struct {
	mu    sync.Mutex // Мьютекс для синхронизации доступа к счетчику
	count int        // Значение счетчика
}

// Метод для инкрементирования счетчика на 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	// Создаем экземпляр счетчика.

	counter := &Counter{}

	// Создаем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Запускаем несколько горутин, которые инкрементируют счетчик.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин.
	wg.Wait()

	// Выводим итоговое значение счетчика.
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.count)

}
