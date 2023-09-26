package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second) // Задержка в 1 секунду между отправками
	}
}

func main() {
	N := 10 // Время работы программы в секундах

	dataChannel := make(chan int)

	// Запускаем горутину-отправителя
	go sender(dataChannel)

	timer := time.NewTimer(time.Duration(N) * time.Second)
	defer timer.Stop()

	for {
		select {
		case data := <-dataChannel:
			fmt.Printf("Принято значение: %d\n", data)
		case <-timer.C:
			fmt.Println("Программа завершена.")
			return
		}
	}
}
