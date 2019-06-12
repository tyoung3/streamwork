package poc

import (
	"fmt"
	"github.com/tyoung3/streamwork"
	"github.com/tyoung3/streamwork/strings"
	"sync"
	"testing"
)

func TestCollate(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	fmt.Println("TestCollate2:")
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))

	fbp.Launch(&wg, []string{"Match0"}, strings.Print1, cs[2:3])
	fbp.Launch(&wg, []string{"Match1"}, strings.Print1, cs[3:4])
	fbp.Launch(&wg, []string{"Miss0"}, strings.Print1, cs[4:5])
	fbp.Launch(&wg, []string{"Miss1"}, strings.Print1, cs[5:6])

	wg.Add(1)
	go func() {
		cs[1] <- 3
		cs[1] <- 6
		cs[1] <- 9
		cs[1] <- 12
		cs[1] <- 12
		cs[1] <- 15
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
		cs[0] <- 14
		close(cs[0])
		fmt.Println("End0")
		wg.Done()
	}()

	//wg.Add(1)
	go fbp.Launch(&wg, []string{"Collate"}, Collate, cs[0:6])
	wg.Wait()

}
