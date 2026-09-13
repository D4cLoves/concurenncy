package pool

import "sync"

// RunPool обрабатывает задачи параллельно в заданном количестве воркеров
// и возвращает сумму результатов.
func RunPool(jobs []int, workers int) int {
	// TODO: реализовать пул воркеров и сбор результатов


	if len(jobs) == 0 {
		return 0
	}
	if workers <= 0 {
		workers = 1
	}

	var workCh = make(chan int, len(jobs))
	var resultCh = make(chan int, len(jobs))
	
	for _, job := range jobs {
		workCh <- job
	}
	close(workCh)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range workCh {
				resultCh <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	
	sum := 0
	for val := range resultCh {
		sum += val
	}
	
	return sum
}
