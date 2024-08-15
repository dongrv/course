package train

import "sync"

func RunGoroutine() {
	var wg sync.WaitGroup
	var i *struct{ Id int }
	println(i.Id)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println("goroutine id is ", i)
		}(i)
	}
	wg.Wait()
}
