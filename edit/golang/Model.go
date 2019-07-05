package edit

/*  Generated StreamWork component, Model,
    by sw.sh on Wed Jun 26 19:55:26 EDT 2019
	1  input  ports
	1 output ports
*/

import (
	"fmt"
	"sync"
)

/* Model keeps a current record of the network definitiion 
*/
func Model(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {

	defer wg.Done()
	cfg := pkgConfig("/home/tyoung3/.sw/edit.toml")
	bs, _ := cfg.IntOr("edit.buffersize", 1)
	fmt.Println(
		"Running", arg[0], version, "bs =", bs)

	cs[1] <- 1

	j := 0
	for j >= 0 {
			ip, ok := <-cs[j]
			if ok != true {
				break
			}
			fmt.Println(arg[0], "chan:", j, "IP:", ip)
		j--
	}
}
