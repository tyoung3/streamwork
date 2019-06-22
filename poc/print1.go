package poc

import "fmt"
import "sync"

/*
Print1 displays integer and string data from port 0 on the
console, prefixed by the process name(arg[0].
*/
func Print1(wg *sync.WaitGroup,
	arg []string,
	cs []chan interface{}) {

	defer wg.Done()
	c := cs[0]
	s := arg[0]

	for {
		ip, ok := <-c
		if ok != true {
			return
		}
		switch ip.(type) {
		case string:
			fmt.Println(s, ip)
		case int:
			fmt.Println(s, "Int:", ip)
		default:
			fmt.Println(s, "type mismatch")
		}
	}
}
