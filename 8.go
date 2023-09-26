package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setBit(num int64, i uint, value bool) int64 {
	if value {
		// Устанавливаем i-й бит в 1
		return num | (1 << i)
	} else {
		// Устанавливаем i-й бит в 0
		return num &^ (1 << i)
	}
}

func main() {
	var num int64
	var i uint
	var newValue bool

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите число (int64): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return
	}

	fmt.Print("Введите позицию бита (uint): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	i64, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Ошибка чтения позиции бита:", err)
		return
	}
	i = uint(i64)

	fmt.Print("Введите значение бита (true/false): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	newValue, err = strconv.ParseBool(input)
	if err != nil {
		fmt.Println("Ошибка чтения значения бита:", err)
		return
	}

	result := setBit(num, i, newValue)
	fmt.Printf("Число после установки %d-го бита в %v: %d\n", i, newValue, result)
}
