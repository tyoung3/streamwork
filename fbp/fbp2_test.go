package fbp

import (
	"github.com/tyoung3/streamwork/poc"
	"sync"
	"testing"
)

func TestFBP(t *testing.T) {
	var cs []chan interface{}
	var wg1 sync.WaitGroup

	for i := 0; i < 6; i++ {
		cs = append(cs, make(chan interface{}, 1024))
	}

	Launch(&wg1, []string{"P1"}, poc.Print1, cs[3:4])
	Launch(&wg1, []string{"Q1", "7"}, poc.Gen2, cs[3:4])

	Launch(&wg1, []string{"P2"}, poc.Print1, cs[0:1])
	Launch(&wg1, []string{"Q2", "2"}, poc.Gen2, cs[1:2])
	Launch(&wg1, []string{"Q3", "3"}, poc.Gen2, cs[2:3])
	wg1.Add(1)
	poc.MergeInt(&wg1, cs[0], cs[1], cs[2])

	wg1.Wait()

}

/*
	/* NETWORK:
			Partition 1:
	     (Q1 Gen1)out1 -> in1(A Print1)

	     	Partition 2:
		 (Q2 Gen1)out1 -> in1(B Print1)
	     (Q3 Gen1)out1 -> in1(B Print1)

*/
