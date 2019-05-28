package strings


import "testing"
import "fmt"
import "sync"

func TestGen1(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	arg := []string{"TestGen1", "7"}

	fmt.Println(arg[0])
	cs = append(cs, make(chan interface{}))
	c := cs[0]

	go func() {
		for  {
			s, ok := <-c
			if ok == true {
				fmt.Println("TestGen1.go1", s)
			} else {
				fmt.Println("TestGen1 Ended")
				return
			}
		}
	}()

	wg.Add(1)
	go Gen1(&wg, arg, cs)
	wg.Wait()

}
