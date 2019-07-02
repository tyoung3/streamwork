package edit

/*  Generated StreamWork component, Edit,
    by sw.sh on Wed Jun 26 19:54:41 EDT 2019
	1  input  ports
	1 output ports
*/

import "testing"
import "fmt"

func TestSkel_Edit(t *testing.T) {

	cfg := pkgConfig("/home/tyoung3/.sw/edit.toml")
	bs, _ := cfg.IntOr("edit.buffersize", 1)
	title, _ := cfg.String("edit.title")
	seqno, _ := cfg.IntOr("edit.seqno", 0)
	fmt.Println(title,
		"\u001b[34mRunning", version, "bs =", bs, "seqno =", seqno,
		"\u001b[0m")
	
	if seqno < 50 || seqno > 10000 {
		fmt.Println("\u001b[34mseqno < 50 or > 10000\u001b[0m")
	}	




}
