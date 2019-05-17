package strings


import "testing"
import "fmt"
import "sync"

func TestComp1(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	arg := []string{"TestComp1", "7"}

	fmt.Println(arg[0])
	cs = append(cs, make(chan interface{}))
	c := cs[0]

	go func() {
		for  {
			s, ok := <-c
			if ok == true {
				fmt.Println("TestCompcomp1.go1", s)
			} else {
				fmt.Println("TestComp1 Ended")
				return
			}
		}
	}()

	wg.Add(1)
	go Comp1(&wg, arg, cs)
	wg.Wait()

}
