package main

import (
	"fmt"
	"strings"
)

// Функция, которая проверяет, что все символы в строке уникальные и регистронезависимые
func areAllCharactersUnique(input string) bool {
	// Приводим строку к нижнему регистру
	input = strings.ToLower(input)

	// Создаем карту для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Проходим по строке
	for _, char := range input {
		// Если символ уже есть в карте, строка не уникальна
		if charMap[char] {
			return false
		}
		// Добавляем символ в карту
		charMap[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	inputString := "Aa"
	result := areAllCharactersUnique(inputString)
	fmt.Println(result) // Выводит false, так как символ "l" повторяется
}
