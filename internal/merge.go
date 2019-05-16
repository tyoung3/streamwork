package internal

import "sync"
//import "fmt"


func Merge(wg1 *sync.WaitGroup,cs ...chan interface{})   {
	var wg sync.WaitGroup

	//fmt.Println("Merge/Fix race bug")
	wg.Add(len(cs)-1)

	for _, c := range cs[1:] {
		go func(c <- chan interface{}) {
			defer wg.Done()
			for v := range c {
				cs[0] <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(cs[0])
        wg1.Done()
	}()

	// return cs[0]
}
