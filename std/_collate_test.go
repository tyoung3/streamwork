package std

import "testing"
import "fmt"
import "sync"

func TestSkel_collate(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	arg := []string{"collate", "dummy arg[1]", "dummy arg[2]"}

	fmt.Println(arg[0])
	cs = append(cs, make(chan interface{}))
	c := cs[0]

	go func() {
		for 1 == 1 {
			s, ok := <-c
			if ok == true {
				fmt.Println("Out1", s)
			} else {
				fmt.Println("Test_collate Ended")
				wg.Done()
				return
			}
		}
	}()

	wg.Add(2)
	go Launch(&wg, arg, cs)
	wg.Wait()

}
