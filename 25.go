package main

import (
	"fmt"
	"time"
)

// Sleep - собственная функция sleep, принимающая количество секунд в качестве параметра
func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")
	Sleep(3) // Вызываем нашу функцию sleep на 3 секунды
	fmt.Println("Прошло 3 секунды")
}
