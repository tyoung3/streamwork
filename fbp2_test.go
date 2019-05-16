package fbp

import (
	"testing"
	"fbp/strings"
	"fbp/internal"
	"sync"
)

func TestFBP(t *testing.T) {
	var cs []chan interface{} 
	var wg1 sync.WaitGroup
	
	for i:=0; i<6; i++ {
		cs = append(cs,make(chan interface{},1024))
	}	
     
	  Launch(&wg1,[]string{"P1"}, 	   strings.Comp2, cs[3:4])
	  Launch(&wg1,[]string{"Q1","7"}, strings.Comp1, cs[3:4]) 
	
	  Launch(&wg1,[]string{"P2"}, 	   strings.Comp2, cs[0:1]) 
	  Launch(&wg1,[]string{"Q2","2"}, strings.Comp1, cs[1:2]) 
	  Launch(&wg1,[]string{"Q3","3"}, strings.Comp1, cs[2:3])  
	  		wg1.Add(1)
           internal.Merge(&wg1,cs[0], cs[1], cs[2]) 
  
	  wg1.Wait()
 
}

/* 
	/* NETWORK:
			Partition 1:
	     (Q1 comp1)out1 -> in1(A comp2) 
	       
	     	Partition 2: 
		 (Q2 comp1)out1 -> in1(B comp2)
	     (Q3 comp1)out1 -> in1(B comp2)
	  
*/
