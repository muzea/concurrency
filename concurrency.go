package concurrency

func dispatch(concurrencyLimit, total int, loop chan int, execEnd, done chan bool) {
	var concurrency = 0
	var runIndex = 0
	for total > runIndex && concurrencyLimit > concurrency {
		loop <- runIndex
		runIndex++
		concurrency++
	}
	for {
		<-execEnd
		if total > runIndex {
			loop <- runIndex
			runIndex++
		} else {
			concurrency--
			if concurrency == 0 {
				done <- true
				return
			}
		}
	}
}

func Run(callback func(int), concurrencyLimit, total int) {
	loop := make(chan int, concurrencyLimit)
	execEnd := make(chan bool, concurrencyLimit)
	done := make(chan bool, 1)
	var exec = func(index int) {
		callback(index)
		execEnd <- true
	}
	go dispatch(concurrencyLimit, total, loop, execEnd, done)
	for {
		select {
		case index := <-loop:
			{
				go exec(index)
				break
			}
		case <-done:
			{
				return
			}
		}
	}
}
