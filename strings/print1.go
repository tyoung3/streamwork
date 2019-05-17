package strings

import "fmt"
import "sync"

func Print1(wg *sync.WaitGroup,
	arg []string,
	cs []chan interface{}) {
	
	defer wg.Done()
	c := cs[0]
	s := arg[0]
	
	for {
		ip, ok := <-c
		if ok != true {
			// fmt.Println(s,"ended.")
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
