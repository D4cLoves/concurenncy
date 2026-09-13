package pipeline

// Run строит конвейер из трёх стадий: квадрат, умножение на 2 и суммирование.
func Run(nums []int) int {
	// TODO: реализовать конвейер обработки чисел
	var ch = make(chan int, len(nums))
	var ch2 = make(chan int, len(nums))

	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num * num
		}
	}()

	go func() {
		defer close(ch2)
		for num := range ch {
			ch2 <- num * 2
		}
	}()

	var sum int
		for num := range ch2 {
			sum += num
		}

	return sum
}
