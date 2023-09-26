package main

import "fmt"

func intersection(set1, set2 []int) []int {
	// Создаем карты для хранения элементов множеств.
	set1Map := make(map[int]bool)
	set2Map := make(map[int]bool)

	// Заполняем карты элементами из множеств.
	for _, item := range set1 {
		set1Map[item] = true
	}

	for _, item := range set2 {
		set2Map[item] = true
	}

	// Создаем результат пересечения.
	var result []int

	// Проверяем каждый элемент первого множества во втором множестве.
	// Если элемент существует в обоих множествах, добавляем его в результат.
	for item := range set1Map {
		if set2Map[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	// Примеры неупорядоченных множеств.
	set1 := []int{14, 8, 100, 76, 47, 39}
	set2 := []int{34, 5, 100, 1, 6, 39}

	// Выполняем операцию пересечения множеств.
	intersected := intersection(set1, set2)

	// Выводим результат.
	fmt.Println("Результат пересечения:", intersected)
}
