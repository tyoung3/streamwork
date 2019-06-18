package poc

import (
	"fmt"
	"github.com/tyoung3/streamwork/fbp"
	"sync"
	"testing"
)

func TestMerge(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	fmt.Println("TestMerge:")
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))

	fbp.Launch(&wg, []string{"PrintM"},  Print1, cs[0:1])

	wg.Add(1)
	go func() {
		cs[1] <- 3
		cs[1] <- 4
		cs[1] <- 4
		cs[1] <- 6
		cs[1] <- 9
		cs[1] <- 12
		close(cs[1])
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cs[2] <- 2
		cs[2] <- 4
		cs[2] <- 6
		cs[2] <- 6
		cs[2] <- 8
		cs[2] <- 10
		cs[2] <- 12
		cs[2] <- 19
		cs[2] <- 21
		close(cs[2])
		wg.Done()
	}()

	go fbp.Launch(&wg, []string{"Merge" }, Merge, cs[0:3])
	wg.Wait()
	fmt.Println("TestMerge ended");

}
