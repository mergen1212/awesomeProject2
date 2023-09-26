package main

// пример 1

//package main
//
//import (
//	"context"
//	"fmt"
//	"sync"
//)
//func worker(stopCh chan bool, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		select {
//		case <-stopCh:
//			fmt.Println("Горутина завершена.")
//			return
//		default:
//			// Выполняем работу горутины
//		}
//	}
//}
//
//func main() {
//	stopCh := make(chan bool)
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go worker(stopCh, &wg)
//	// Ожидаем нажатия клавиши, затем останавливаем горутину
//	fmt.Println("Нажмите Enter для остановки горутины...")
//	fmt.Scanln()
//	close(stopCh)
//	wg.Wait()
//	fmt.Println("завершена горутина worker")
//}

//пример 2

//package main
//
//import (
//	"context"
//	"fmt"
//	"sync"
//)
//
//func worker(ctx context.Context, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("Горутина завершена.")
//			return
//		default:
//			// Выполняем работу горутины
//		}
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//	go worker(ctx, &wg)
//
//	// Ожидаем нажатия клавиши, затем останавливаем горутину
//	fmt.Println("Нажмите Enter для остановки горутины...")
//	fmt.Scanln()
//	cancel() // Отправляем сигнал для остановки горутины
//	wg.Wait()
//	fmt.Println("Программа завершена.")
//}

//пример 3

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var mu sync.Mutex
//var stopFlag bool
//
//func worker(wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		mu.Lock()
//		if stopFlag {
//			mu.Unlock()
//			fmt.Println("Горутина завершена.")
//			return
//		}
//		mu.Unlock()
//
//		// Выполняем работу горутины
//
//		time.Sleep(time.Second)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//	go worker(&wg)
//
//	// Ожидаем нажатия клавиши, затем останавливаем горутину
//	fmt.Println("Нажмите Enter для остановки горутины...")
//	fmt.Scanln()
//
//	mu.Lock()
//	stopFlag = true
//	mu.Unlock()
//
//	wg.Wait()
//	fmt.Println("Программа завершена.")
//}
