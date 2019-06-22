package def

import "fmt"
import "sync"

/*
Print1 prints integers and strings received from port 0, prefixed
with the process identifier(arg[0]) string,
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
