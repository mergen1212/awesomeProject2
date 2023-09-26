package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем значения a и b с использованием библиотеки math/big
	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем значения a и b (больше чем 2^20)
	a.Exp(big.NewInt(2), big.NewInt(22), nil) // a = 2^22
	b.Exp(big.NewInt(2), big.NewInt(21), nil) // b = 2^21

	// Умножение
	resultMultiply := new(big.Int)
	resultMultiply.Mul(a, b)

	// Деление
	resultDivide := new(big.Rat).SetFrac(a, b)

	// Сложение
	resultAdd := new(big.Int)
	resultAdd.Add(a, b)

	// Вычитание
	resultSubtract := new(big.Int)
	resultSubtract.Sub(a, b)

	// Вывод результатов
	fmt.Printf("Умножение: %s\n", resultMultiply.String())
	fmt.Printf("Деление: %s\n", resultDivide.FloatString(10))
	fmt.Printf("Сложение: %s\n", resultAdd.String())
	fmt.Printf("Вычитание: %s\n", resultSubtract.String())
}
