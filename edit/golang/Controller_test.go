package edit

/*  Generated StreamWork component, Controller,
    by sw.sh on Wed Jun 26 20:00:23 EDT 2019
	2  input  ports
	2 output ports
*/

import "testing"
import "fmt"
import "sync"

func TestSkel_Controller(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	arg := []string{"TestSkel_Controller"}

	fmt.Println(arg[0])
	for i := 0; i < 2+2; i++ {
		cs = append(cs, make(chan interface{}))
	}

	wg.Add(2)
	go func() {

		j := 3
		for j >= 2 {
			ip, ok := <-cs[j]
			if ok != true {
				break
			}
			fmt.Println("chan:", j, "IP:", ip)
			j--
		}

		i := 1
		for i >= 0 {
			cs[i] <- i
			close(cs[i])
			i--
		}
		fmt.Println("TestSkel_Controller Ended")
		wg.Done()
		return
	}()

	go Controller(&wg, arg, cs)
	wg.Wait()

}
