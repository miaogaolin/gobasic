package testatomic

import "sync"

func AddNumber(num int) int {
	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			i += num
			wg.Done()
		}(i)
	}
	wg.Wait()
	return num
}
