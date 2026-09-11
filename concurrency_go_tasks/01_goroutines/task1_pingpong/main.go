package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	// TODO: реализовать обмен сообщениями между горутинами
	var wg sync.WaitGroup
	wg.Add(2)

	var pingCh = make(chan string, 1)
	var pongCh = make(chan string, 1)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			fmt.Fprintf(w, "%s\n", <-pingCh)
			pongCh <- "pong"
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Fprintf(w, "%s\n", <-pongCh)
			if i < 4 {
				pingCh <- "ping"
			}
		}
	}()

	pingCh <- "ping"

	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
}
