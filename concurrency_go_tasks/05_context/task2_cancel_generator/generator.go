package generator

import "context"

// Generate возвращает канал, из которого можно читать возрастающие числа,
// начиная с нуля. Генерация прекращается при отмене ctx.
func Generate(ctx context.Context) <-chan int {
	// TODO: реализовать генератор чисел с учётом отмены

	numChan := make(chan int)
	num := 0

	if ctx.Err() != nil {
		close(numChan)
		return numChan
	}
	
	go func() {
		defer close(numChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			select {
			case <-ctx.Done():
				return
			case numChan <- num:
				num++
			}
		}
	}()
	
	return numChan
}
