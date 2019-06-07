package std

//import "os"
import "fmt"
import "sync"
import "strconv"

var version string = "v0.0.0"

/* Gen1 sends 'nbr' integers, beginning with 'start', incremented
by 'inc';  argcuments 1, 2, and 3 respectively.

*/
func Gen1(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {

	defer wg.Done()
	// fmt.Println("Running", arg[0])
	c := cs[0]

	//  Get arguments from os.args
	fmt.Println(arg[0], " std.Gen1", arg[1], arg[2], arg[3])
	/* stub for now */
	nbr, _ := strconv.Atoi(arg[1])
	start, _ := strconv.Atoi(arg[2])
	inc, _ := strconv.Atoi(arg[3])

	ip := start

	for n := 0; n < nbr; n++ {
		c <- ip
		ip += inc
	}

	close(c)
}
