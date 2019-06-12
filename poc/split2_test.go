package poc

/* NETWORK:
   (Q Gen1)out1 -> in1(S split)out1 -> in1(A1 Print1)
   (S)out2                           -> in1(A2 Print1)
   (S)out3					       -> in1(A3 Print1)
*/

import (
	"github.com/tyoung3/streamwork"
	"github.com/tyoung3/streamwork/strings"
	"sync"
	"testing"
)

func TestSplit2(t *testing.T) {
	var cs []chan interface{}
	var wg1 sync.WaitGroup

	for i := 0; i < 4; i++ {
		cs = append(cs, make(chan interface{}, 1024))
	}

	// wg1.Add(1)
	go fbp.Launch(&wg1, []string{"S"}, Split, cs[0:4])
	fbp.Launch(&wg1, []string{"Q", "3"}, strings.Gen1, cs[0:1])
	fbp.Launch(&wg1, []string{"A2"}, strings.Print1, cs[1:2])
	fbp.Launch(&wg1, []string{"A1"}, strings.Print1, cs[2:3])
	fbp.Launch(&wg1, []string{"A3"}, strings.Print1, cs[3:4])

	wg1.Wait()

}
