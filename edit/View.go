/* package edit is generated by StreamWork.*/
package edit

/*  Generated StreamWork component, View,
    by sw.sh on Wed Jun 26 19:56:37 EDT 2019
	1  input  ports
	1 output ports
*/

import (
	"fmt"
	"sync"
)

func View(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var version string = "v0.0.0"

	defer wg.Done()
	cfg := pkgConfig("/home/tyoung3/.sw/edit.toml")
	bs, _ := cfg.IntOr("edit.buffersize", 1)
	fmt.Println(
		"Running", arg[0], version, "bs =", bs)
	cs[1] <- 1

	j := 0
	for j >= 0 {
		_ = <-cs[j]
		j--
	}
}
