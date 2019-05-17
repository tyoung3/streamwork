package  std

import "sync"
import "fmt"

func Merge(wg1 *sync.WaitGroup, arg []string, cs []chan interface{})  {
	var wg sync.WaitGroup

	fmt.Println(arg[0],"std.Merge: starting ")
	defer wg1.Done()
	nc := len(cs) 
	wg.Add(nc-1)

	for _, c := range cs[1:] {
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				fmt.Println(arg[0],"std.Merge: ", v, nc)
				cs[0] <- v
			}
			fmt.Println(arg[0],"std.Merge: exit ")
		}(c)
	}

	wg.Wait()
	fmt.Println(arg[0],"std.Merge: ending")
	close(cs[0]) 
 
}
