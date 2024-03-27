// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала
// и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// воркеры для считывания из канала
func worker(job <-chan int, id int) {
	for {
		fmt.Println("worker", id, "done job", <-job)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	var N int

	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	job := make(chan int, 100)
	for i := 0; i < N; i++ {
		go worker(job, i)
	}

	// graceful shutdown
	go func() {
		<-sigCh
		fmt.Printf("\nSIGTERM captured, %d goroutines killed\n", N)
		os.Exit(0)
	}()

	for {
		job <- rand.Int()
	}

}
