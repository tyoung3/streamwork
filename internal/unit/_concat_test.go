package unit

		/*  	Network
			(G2 comp1 "2")out -> in1(C concat) -> in(Sink comp2) 
	        (G3 comp1 "3")out -> in2(C)
	        (G5 comp1 "5")out -> in3(C
	    */ 
	
import "testing"
import "fmt"
import "github.com/tyoung3/streamwork"
import "github.com/tyoung3/streamwork/strings"
import "github.com/tyoung3/streamwork/std"
import "sync"

func TestConcat(t *testing.T) {
	var cs  []chan interface{}
	var wg  sync.WaitGroup
	
	for i:=0; i<4; i++ {
		cs = append(cs,make(chan interface{},1024))
	}	
	
	fmt.Println("Testing Concat")
		//wg.Add(1)
	fbp.Launch(&wg, []string{"Cat"},     std.Concat, cs[0:4])
	fbp.Launch(&wg, []string{"G2","2"},  strings.Comp1, cs[1:2])
	fbp.Launch(&wg, []string{"G3","3"},  strings.Comp1, cs[2:3]) 
	fbp.Launch(&wg, []string{"G5","5"},  strings.Comp1, cs[3:4])  
	fbp.Launch(&wg, []string{"Sink"},    strings.Comp2, cs[0:1])
	  
	wg.Wait()	

}
	
