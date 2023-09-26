package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}

func reverseString(s string) string {
	// Конвертируем строку в массив рун (Unicode символов)
	runes := []rune(s)
	length := len(runes)

	// Индексы начального и конечного элементов для обмена
	start := 0
	end := length - 1

	// Переворачиваем строку, меняя местами символы
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	// Конвертируем массив рун обратно в строку
	return string(runes)
}

func main() {
	var input string
	fmt.Print("Введите строку: ")
	input = Scan()
	reversed := reverseString(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
