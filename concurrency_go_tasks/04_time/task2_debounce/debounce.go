package debounce

import "time"

// Debounce принимает значения и отдаёт только последнее после паузы d.
func Debounce(d time.Duration, in <-chan int) <-chan int {
	// TODO: реализовать дебаунс значений из канала

	out := make(chan int)
	a := time.NewTimer(d)

	lastVal := 0
	isOpen := false
	go func() {
		defer a.Stop()
		for {
			select {
			case val, ok := <-in:
				if !ok {
					if isOpen {
						<-a.C
						out <- lastVal
					}
					close(out)
					return
				}
				a.Reset(d)
				lastVal = val
				isOpen = true
			case <-a.C:
				if isOpen {
					out <- lastVal
				}
				isOpen = false
			}
		}
	}()
	
	return out
}
