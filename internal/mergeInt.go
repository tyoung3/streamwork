package internal

// github/tyoung3/streamwork/internal

import "sync"
import "fmt"

/* Merge sends all input from channels cs[1:] to channel cs[0]
 */
func MergeInt(wg1 *sync.WaitGroup, cs ...chan interface{}) {
	var wg sync.WaitGroup

	// fmt.Println("Merge stgit checkout -b art")
	defer wg1.Done()
	wg.Add(len(cs) - 1)

	for _, c := range cs[1:] {
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				fmt.Println("Merge:", v)
				cs[0] <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		// fmt.Println("Merge done")
		close(cs[0])
	}()

	// return cs[0]
}
