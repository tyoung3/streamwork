package poc

/* /home/tyoung3/go/mod/fbp/comp1/comp1.go
 *		 	 DO NOT EDIT!!!
 *		generated by fbpgo.sh-v0.0.1  on Wed May  8 15:47:10 EDT 2019
**/

import "fmt"
import "sync"
import "strconv"

// var Version string = "0.0.1"

/*
Gen2 sends arg[1],  arg[2] long strings to channel cs[0] (out1)
    Strings consist of the process name
   and seq. no; i.e. G3-1, G3-2, etc.

   BUG cannot set length, yet.
*/
func Gen2(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var n0 = 7
	var n int
	//var nb int = 80

	defer wg.Done()

	//pname := arg[0]
	l := len(arg)
	if l > 1 {
		n0, _ = strconv.Atoi(arg[1])
		/* if l > 2 {
		    _,_ := strconv.Atoi(arg[2])
		}
		*/
	}

	c := cs[0]

	for n = 1; n < n0+1; n++ {
		c <- fmt.Sprintf("%s-%d", arg[0], n)
	}

	close(c)
}
