package poc

/*  	Network
		(G2 Gen1 "2")out -> in1(M Join) -> in(Sink Print1)
        (G3 Gen1 "3")out -> in2(M)
        (G5 Gen1 "5")out -> in3(M)
*/

import "testing"
import "fmt"
import "sync"
import "github.com/tyoung3/streamwork/fbp"

func TestJoin(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		cs = append(cs, make(chan interface{}))
	}

	fmt.Println("Testing Join-0.0.3")
	wg.Add(1)
	go Join(&wg, []string{"Join"}, cs[0:4])
	fbp.Launch(&wg, []string{"G2", "2"}, Gen2, cs[1:2])
	fbp.Launch(&wg, []string{"G3", "3"}, Gen2, cs[2:3])
	fbp.Launch(&wg, []string{"G5", "5"}, Gen2, cs[3:4])
	fbp.Launch(&wg, []string{"Sink"}, Print1, cs[0:1])

	fmt.Println("TestJoin waiting")
	wg.Wait()
	fmt.Println("TestJoin ended")
}
