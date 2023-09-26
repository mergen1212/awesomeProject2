package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}

func reverseWordsInString(inputString string) string {
	words := strings.Fields(inputString) // Разбиваем строку на слова
	// Переворачиваем каждое слово
	for i, _ := range words {
		words[i] = words[len(words)-i-1]
	}
	reversedString := strings.Join(words, " ")

	return reversedString
}

func main() {
	var input string
	input = Scan()
	result := reverseWordsInString(input)
	fmt.Println(result)
}
