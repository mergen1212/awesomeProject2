package main

import "fmt"

func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range strings {
		set[str] = true
	}

	return set
}

func main() {
	// Исходная последовательность строк.
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество строк.
	stringSet := createSet(strings)

	// Выводим множество.
	fmt.Println("Множество строк:", stringSet)

}
