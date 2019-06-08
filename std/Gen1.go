package std

import "fmt"
import "sync"
import "strconv"

var version = "v0.0.0"

/* 
Gen1 sends 'nbr' integers, beginning with 'start', incremented
by 'inc';  argcuments 1, 2, and 3 respectively over port 0.

#####Examples
		(Gen1) will send:
		1 2 3 4 5 6 7

	 	(Gen1 "-i" "9" "2" "3") will send:
	2 5 8 11 14 17 20 23 26

	    (Gen1 "-i" "3" "17") will send:
	17 18 19

		(Gen1 "-i" "4" "-2" "-1") will send:
	-2 -3 -4 -5
*/
func Gen1(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {

	defer wg.Done()
	c := cs[0]

	fmt.Println(arg[0], " std.Gen1", arg[1], arg[2], arg[3])
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
