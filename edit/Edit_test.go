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

	wg.Add(2)
	go func() {
		fmt.Println("TestSkel_Edit Ended")
		wg.Done()
		return
	}()

	go Edit(&wg, arg, cs)
	wg.Wait()

}
