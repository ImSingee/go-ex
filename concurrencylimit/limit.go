package concurrencylimit

import "sync"

func DoWithLimit(limit int, tasks ...func()) {
	cc := make(chan struct{}, limit)
	defer close(cc)

	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	for _, task := range tasks {
		task := task
		go func() {
			cc <- struct{}{}
			defer func() {
				<-cc
				wg.Done()
			}()

			task()
		}()
	}

	wg.Wait()
}
