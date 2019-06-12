package poc

/*  	Network
		(G2 Gen1 "2")out -> in1(C concat) -> in(Sink Print1)
        (G3 Gen1 "3")out -> in2(C)
        (G5 Gen1 "5")out -> in3(C
*/

import "testing"
import "fmt"
import "github.com/tyoung3/streamwork"
import "github.com/tyoung3/streamwork/strings"
import "sync"

func TestConcat(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		cs = append(cs, make(chan interface{}, 1024))
	}

	fmt.Println("Testing Concat")
	//wg.Add(1)
	go fbp.Launch(&wg, []string{"Cat"}, Concat, cs[0:4])
	fbp.Launch(&wg, []string{"G2", "2"}, strings.Gen1, cs[1:2])
	fbp.Launch(&wg, []string{"G3", "3"}, strings.Gen1, cs[2:3])
	fbp.Launch(&wg, []string{"G5", "5"}, strings.Gen1, cs[3:4])
	fbp.Launch(&wg, []string{"Sink"}, strings.Print1, cs[0:1])

	wg.Wait()

}
