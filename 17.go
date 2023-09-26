package main

import (
	"fmt"
	"sort"
)

func main() {
	// Пример отсортированного среза данных.
	data := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Элемент, который мы ищем.
	target := 7

	// Используем функцию sort.Search для выполнения бинарного поиска.
	index := sort.Search(len(data), func(i int) bool {
		return data[i] >= target
	})

	// Проверяем, найден ли элемент.
	if index < len(data) && data[index] == target {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
