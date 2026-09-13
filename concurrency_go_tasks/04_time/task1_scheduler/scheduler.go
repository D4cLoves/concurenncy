package scheduler

import (
	"sync"
	"time"
)

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	// TODO: периодический вызов функции с возможностью остановки
	var ticker = time.NewTicker(d)
	var stopCh = make(chan struct{})
	var once sync.Once

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-stopCh:
				return
			}
		}
	}()

	stop = func() {
		once.Do(func() {
			close(stopCh)
		})
	}

	return stop
}
