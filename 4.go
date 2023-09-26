package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChannel <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				return // Канал закрыт, завершаем работу
			}
			fmt.Printf("Worker %d: Произвольные данные: %d\n", id, data)
		case <-ctx.Done():
			return // Завершаем работу по сигналу от контекста
		}
	}
}

func main() {
	numWorkersPtr := flag.Int("workers", 5, "Количество воркеров")
	flag.Parse()

	if *numWorkersPtr <= 0 {
		fmt.Println("Количество воркеров должно быть больше нуля.")
		os.Exit(1)
	}

	dataChannel := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем воркеров
	for i := 1; i <= *numWorkersPtr; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, ctx)
	}

	// Подписываемся на сигналы Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Главный поток постоянно отправляет данные в канал
	for i := 1; ; i++ {
		select {
		case <-sigCh:
			// Получен сигнал Ctrl+C, завершаем работу
			fmt.Printf("Получен сигнал Ctrl+C")
			cancel()
			close(dataChannel)
			wg.Wait()
			return
		default:
			dataChannel <- i
		}
	}
}
