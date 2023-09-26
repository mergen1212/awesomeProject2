package main

import (
	"fmt"
	"math"
)

func main() {
	// Последовательность температурных колебаний.
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения группированных данных.
	temperatureGroups := make(map[int][]float64)

	// Определяем шаг для группировки.
	step := 10

	// Группируем температурные значения.
	for _, temp := range temperatures {
		groupKey := int(math.Floor(temp/float64(step))) * step
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temp)
	}

	// Выводим результаты.
	for key, values := range temperatureGroups {
		fmt.Printf("%d: %v, ", key, values)
	}
	fmt.Printf("etc.")
}
