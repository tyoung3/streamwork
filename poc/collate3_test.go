package poc

import (
	"fmt"
	"github.com/tyoung3/streamwork/fbp"
	"sync"
	"testing"
)

func TestCollate3(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	fmt.Println("TestCollate3:")
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))

	fbp.Launch(&wg, []string{"PrintM"},  Print1, cs[2:3])

	wg.Add(1)
	go func() {
		cs[1] <- 3
		cs[1] <- 4
		cs[1] <- 4
		cs[1] <- 6
		cs[1] <- 9
		cs[1] <- 12
		close(cs[1])
		fmt.Println("End1")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cs[0] <- 2
		cs[0] <- 4
		cs[0] <- 6
		cs[0] <- 6
		cs[0] <- 8
		cs[0] <- 10
		cs[0] <- 12
		cs[0] <- 19
		cs[0] <- 21
		close(cs[0])
		fmt.Println("End0")
		wg.Done()
	}()

	//wg.Add(1)
	go fbp.Launch(&wg, []string{"Collate","--merge"}, Collate, cs[0:3])
	wg.Wait()
	fmt.Println("TestCollate3 ended");

}
