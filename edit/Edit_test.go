package edit

/*  Generated StreamWork component, Edit,
    by sw.sh on Wed Jun 26 19:54:41 EDT 2019
	1  input  ports
	1 output ports
*/

import "testing"
import "fmt"
import "sync"

func TestSkel_Edit(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	arg := []string{"TestSkel_Edit"}

	fmt.Println(arg[0])
	for i := 0; i < 1+1; i++ {
		cs = append(cs, make(chan interface{}))
	}

	wg.Add(2)
	go func() {

		j := 1
		for j >= 1 {
			ip, ok := <-cs[j]
			if ok != true {
				break
			}
			fmt.Println("chan:", j, "IP:", ip)
			j--
		}

		i := 0
		for i >= 0 {
			cs[i] <- i
			close(cs[i])
			i--
		}
		fmt.Println("TestSkel_Edit Ended")
		wg.Done()
		return
	}()

	go Edit(&wg, arg, cs)
	wg.Wait()

}
