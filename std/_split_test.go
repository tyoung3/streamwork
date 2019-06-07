package std

/*
	copy IPs from input  port: cs[0]
	to all output ports: "cs[1:]"
*/

import "testing"
import "fmt"
import "sync"

func TestSplit(t *testing.T) {
	var cs []chan interface{} /* A slice of channels */
	var wg sync.WaitGroup
	const nc = 2 /* Nbr of split output channels. */

	arg := []string{"split"}

	fmt.Println(arg[0], "Test")
	for n := 0; n < nc+1; n++ {
		cs = append(cs, make(chan interface{}, 3))
	}

	c := cs[0] /* input channel */

	go func() { /* Send seven IPs to in(split) */
		fmt.Println("Sending IPs")
		for i := 1; i < 8; i++ {
			c <- i
		}
		fmt.Println("Sent seven IPs")
		close(c)
		wg.Done()
	}()

	/* Receive input from all (split)out[N]... */
	for j := 1; j < nc+1; j++ {
		go func(ci chan interface{}, n int) {
			fmt.Println("Getting channel", n)
			defer wg.Done()
			for {
				s, ok := <-ci
				if ok == true {
					fmt.Println("Out", n, s)
				} else {
					fmt.Println("End of channel", n)
					return
				}
			}
		}(cs[j], j)
	}

	wg.Add(nc - 1)
	go Split(&wg, arg, cs)
	wg.Wait()

}
