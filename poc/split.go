package poc

import "sync"

// var version string = "v0.0.1"

/*
Split copies all input from port 0, to all active ports 1,..N
*/
func Split(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var wg2 sync.WaitGroup
	var c2s []chan interface{}

	defer wg.Done()

	for range cs[1:] {
		c2s = append(c2s, make(chan interface{}))
	}

	for j := range cs[1:] { /* Start a routine for each output */
		wg2.Add(1)
		go func(c chan interface{}, c2 chan interface{}) {
			defer wg2.Done()
			for {
				ip, ok := <-c2
				if ok != true {
					close(c)
					return
				}
				c <- ip
			}
		}(cs[j+1], c2s[j])
	}

	wg2.Add(1)
	go func() {
		for {
			ip, ok := <-cs[0]
			if ok != true {
				for j := range c2s {
					close(c2s[j])
				}
				wg2.Done()
				return
			}
			for j := range c2s {
				c2s[j] <- ip
			}
		}
	}()

	wg2.Wait()
}
