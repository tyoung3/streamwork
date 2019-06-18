package poc

import "sync"
import "reflect"

func merge(cs []chan interface{}) {
	var state = 0
	
	ip0 := *new(interface{})
	ip1 := *new(interface{})
	
	for {
		switch state {
		case 0, 2, 8: /* get 0 and add 1 to state. If EOF, add 4 to state */
			ip0 = get0(&state, cs[2])
		case 1, 4: /* get 1 and add 2 to state. if EOF state+=8 */
			ip1 = get1(&state, cs[1])
		case 3:
			if ip0 == ip1 {
				cs[0] <- ip0  
				state--
			} else {
				val0 := reflect.ValueOf(ip0).Int()
				val1 := reflect.ValueOf(ip1).Int()
				if val0 > val1 {
					cs[0] <- ip1
					state = 1
				} else {
					cs[0] <- ip0
					state = 2
				}
			}
		case 6:  
			cs[0] <- ip1
			state = 4
		case 9: 
			cs[0] <- ip0
			state = 8
		case 12:
			close(cs[0])
			return

		}
	}
		
}

/* 
Merge compares input from channels 1 and 2 sending
   the lowest data to channel 0. 
*/
func Merge(wg *sync.WaitGroup, 
			 arg []string, 
			 cs []chan interface{}) {
			 
	defer wg.Done()
	
	merge(cs)	
	
}
