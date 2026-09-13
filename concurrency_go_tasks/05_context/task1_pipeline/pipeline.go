package pipelinectx

import "context"

// Run строит конвейер из двух стадий: удвоение и суммирование.
// Конвейер должен останавливаться, если ctx отменён.
// Возвращает итоговую сумму и ошибку контекста при отмене.
func Run(ctx context.Context, nums []int) (int, error) {
	// TODO: реализовать конвейер с остановкой по ctx

	if err := ctx.Err(); err != nil {
    	return 0, err
	}
	
	var numChan = make(chan int)
	sum := 0

	go func() {
		defer close(numChan)
		for _, num := range nums {
			select {
			case <-ctx.Done():
				return
			case numChan <- num * 2:
			}
		}
	}()

	for num := range numChan {
		sum += num
	}

	if err := ctx.Err(); err != nil {
		return 0, err
	}
	
	return sum, nil
}
