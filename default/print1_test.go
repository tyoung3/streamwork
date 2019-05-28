package strings

/* /home/tyoung3/go/mod/fbp/comp2/comp2_test.go
 *		 	 DO NOT EDIT!!!
 *		generated by fbpgo.sh-v0.0.1  on Wed May  8 15:47:11 EDT 2019
**/

import "testing"
import "fmt"
import "sync"

func TestPrint1(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup
	arg := []string{"TestPrint1", "xx"}

	fmt.Println("TestPrint1")
	cs = append(cs, make(chan interface{}))
	c := cs[0]

	go func() {
		c <- 7
		c <- 11
		c <- "OK?"
		close(c)
	}()

	wg.Add(1)
	go Print1(&wg, arg, cs)
	fmt.Println("TestPrint1 wait c ")
	wg.Wait()
	fmt.Println("TestPrint1 done.")
}