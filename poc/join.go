package poc

import "sync"

/*
Join sends all input from channels [1:n] to channel [0], port 0.
*/
func Join(wg1 *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var wg sync.WaitGroup

	defer wg1.Done()
	nc := len(cs)
	wg.Add(nc - 1)

	for _, c := range cs[1:] {
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				cs[0] <- v
			}
		}(c)
	}

	wg.Wait()
	close(cs[0])

}
