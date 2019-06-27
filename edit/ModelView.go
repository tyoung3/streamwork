/* package edit is generated by StreamWork.*/
package edit

/*  Generated StreamWork component, ModelView,
    by sw.sh on Wed Jun 26 20:00:48 EDT 2019
	2  input  ports
	2 output ports
*/

import (
	"fmt"
	"sync"
)

func ModelView(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var version string = "v0.0.0"

	defer wg.Done()
	cfg := pkgConfig("/home/tyoung3/.sw/foo2.toml")
	seqno, _ := cfg.IntOr("edit.seqno", 1)
	bs, _ := cfg.IntOr("edit.buffersize", 1)
	fmt.Println(
		"Running", arg[0], version, "bs =", bs)

	i := 3

	for i >= 2 {
		cs[i] <- i
		close(cs[i])
		i--
	}

	j := 1
	for j >= 0 {
		_ = <-cs[j]
		j--
	}
}
