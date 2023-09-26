package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Создаем sync.Map для конкурентной записи
	data := new(sync.Map)
	// Количество горутин для конкурентной записи
	numWorkers := 1000
	// Функция для конкурентной записи данных
	worker := func(key string, value int) {
		defer wg.Done()
		data.Store(key, value)
	}
	// Запускаем горутины для записи данных
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		key := fmt.Sprintf("Key%d", i)
		value := i * 10
		go worker(key, value)
	}
	// Ожидаем завершения всех горутин
	wg.Wait()
	// Выводим результаты
	data.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}
