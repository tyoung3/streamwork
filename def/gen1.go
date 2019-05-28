package def 

import "fmt"
import "sync"
import "strconv"

var Version string ="0.0.2" 


/* Gen1 sends arg[1],  arg[2] long strings to channel cs[0] (out1) 
    Strings consist of the process name only
    i.e. G3  G3 , etc
*/
func Gen1(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var n0 int = 1
	var n int

	defer wg.Done()
	
	l := len(arg)
	if l > 1 {
		n0, _ = strconv.Atoi(arg[1])
	}

	c := cs[0]

	for n = 1; n < n0+1; n++ {
		c <- fmt.Sprintf("%s-%d", arg[0], n)
	}

	close(c)
}
