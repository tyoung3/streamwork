package unit

	/* NETWORK:
	     (Q comp1)out1 -> in1(S split)out1 -> in1(A1 comp2) 
	     (S)out2                           -> in1(A2 comp2) 
	     (S)out3					       -> in1(A3 comp2)
	*/
	
import (
	"testing"
	"github.com/tyoung3/streamwork"
	"github.com/tyoung3/streamwork/strings"
	"github.com/tyoung3/streamwork/std"
	"sync"
)

func TestSplit(t *testing.T) {
	var cs []chan interface{} 
	var wg1 sync.WaitGroup
	
	for i:=0; i<4; i++ {
		cs = append(cs,make(chan interface{},1024))
	}	
      
	  fbp.Launch(&wg1, []string{"S"},      std.Split,  cs[0:4])
	  fbp.Launch(&wg1,[]string{"Q","3"},   strings.Comp1, cs[0:1]) 
	  fbp.Launch(&wg1, []string{"A2"}, 	   strings.Comp2, cs[1:2])
      fbp.Launch(&wg1,[]string{"A1"}, 	   strings.Comp2, cs[2:3])
      fbp.Launch(&wg1,[]string{"A3"}, 	   strings.Comp2, cs[3:4])
      
	  wg1.Wait()
 
}

