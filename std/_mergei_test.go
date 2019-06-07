package std

// Fix -race bug

import "testing"

// import "fbp/internal/merge"
import "fmt"
import "sync"

func TestMergeI(t *testing.T) {
	fmt.Println("TestMerge Start ")
	var wg sync.WaitGroup

	cs := *new([]chan interface{})

	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	cs = append(cs, make(chan interface{}))
	c1 := cs[0]
	c2 := cs[1]
	c3 := cs[2]
	c4 := cs[3]

	wg.Add(1)
	go func() {
		defer wg.Done()
		c1 <- "A:IP1"
		c3 <- "C:IP1"
		c1 <- "A:IP2"
		c4 <- "D:IP1"
		c2 <- "B:IP1"
		c2 <- "B:IP2"
		fmt.Println("End of input")
		close(c1)
		close(c2)
		close(c3)
		close(c4)
	}()

	fmt.Println("TestMerge")

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			i, ok := <-cs[0]
			if ok != true {
				return
			}
			fmt.Println(i)
		}
	}()

	wg.Add(1)
	Merge(&wg, cs[0], cs[1], cs[2], cs[3], cs[4])

	wg.Wait()
	fmt.Println("TestMerge Ended")
}
