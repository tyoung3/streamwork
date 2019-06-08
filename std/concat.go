package std

import "sync"

// var Version string = "v0.0.1"

/* 
Concat sends all port 1 input to port 0, followed by
all port 2 input, etc. until all channel input [1:N] has
been sent to port 0.
*/
func Concat(wg *sync.WaitGroup,
	arg []string,
	cs []chan interface{}) {

	cout := cs[0]

	for j := range cs[1:] {
		func(cin chan interface{}) {
			for {
				ip, ok := <-cin
				if ok != true {
					return
				}
				cout <- ip
			}
		}(cs[j+1])
	}

	close(cout)
	wg.Done()
}
