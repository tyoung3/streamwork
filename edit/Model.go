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
	var version string = "v0.0.0"

	defer wg.Done()
	cfg := pkgConfig("/home/tyoung3/.sw/edit.toml")
	seqno, _ := cfg.IntOr("edit.seqno", 1)
	bs, _ := cfg.IntOr("edit.buffersize", 1)
	fmt.Println(
		"Running", arg[0], version, "bs =", bs)

	if seqno != 1234 {
		fmt.Println(
			"Seqno not = 1234.  Missing config file?")
	}

	cs[1] <- 1

	j := 0
	for j >= 0 {
		_ = <-cs[j]
		j--
	}
}
