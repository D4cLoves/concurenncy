package limiter

import (
	"sync"
	"time"
)

// Limiter ограничивает количество событий до 5 в секунду.
type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
	once   sync.Once
}

// NewLimiter создаёт новый лимитер с ёмкостью 5 токенов.
func NewLimiter() *Limiter {
	// TODO: инициализировать канал токенов и запуск пополнения
	tokens := make(chan struct{}, 5)
	stop := make(chan struct{})

	for i := 0; i < 5; i++ {
		tokens <- struct{}{}
	}
	
	ticker := time.NewTicker(200 * time.Millisecond)
	
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
					case tokens <- struct{}{}:
					default:
				}
				
			case <-stop:
				return
				
			}
		}
	}()
	return &Limiter{tokens: tokens, stop: stop}
}

// Allow возвращает true, если событие разрешено в текущий момент.
func (l *Limiter) Allow() bool {
	// TODO: реализовать получение токена из канала
	select {
	case <-l.tokens:
		return true
		
	default:
		return false
	}
}

// Stop останавливает лимитер.
func (l *Limiter) Stop() {
	// TODO: остановить пополнение токенов
	l.once.Do(func() {
		close(l.stop)
	})
}