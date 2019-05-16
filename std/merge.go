package  std

import "sync"


func Merge(wg1 *sync.WaitGroup, arg []string, cs []chan interface{})  {
	var wg sync.WaitGroup

	defer wg1.Done()
	wg.Add(len(cs) - 1)

	for _, c := range cs[1:] {
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				cs[0] <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(cs[0])
	}()

	return    
}
