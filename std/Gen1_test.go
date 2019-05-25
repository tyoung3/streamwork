package std

/*
	Gem1 Number Start Increment 
	Emit Number integer IPs bginning with Start, incementing IPs by Increment.
	E.x Gen1 "5" "2" "3" should emit 2 5 8 11 14
*/
	
import "testing"
import "fmt"
import "sync"

func TestGen1(t *testing.T) {
	var cs  []chan interface{}
	var wg  sync.WaitGroup
	
	arg := []string{"ABC","4","2","5"}
	
	cs = append(cs,make(chan interface{}))
	c  := cs[0]
	
	go func() { 
	  defer wg.Done()
	  for  {
		s, ok := <-c
		if ok == true {
			fmt.Println("tGen1", s)
		} else {
			fmt.Println("TestGen1 Ended")
			return
		}
	   }	
	}() 
	
	wg.Add(2)
	go Gen1(&wg, arg,  cs)
	wg.Wait()	

}
	
