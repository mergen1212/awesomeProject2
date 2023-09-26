package main

import (
	"fmt"
	"sort"
)

// Создаем тип для сортировки.
type IntSlice []int

// Реализуем методы интерфейса sort.Interface.
func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	// Пример неотсортированного массива.
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	// Преобразуем массив в IntSlice.
	intSlice := IntSlice(arr)

	// Сортируем массив с использованием быстрой сортировки.
	sort.Sort(intSlice)

	// Выводим отсортированный массив.
	fmt.Println("Отсортированный массив:", intSlice)
}
