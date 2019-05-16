package std

import "fmt"
import "sync"
import "reflect"

func get0(state *int, c chan interface{} )  (interface{}) {
	ip, ok := <- c
	if ok {
		 *state += 1
		 return ip 
	} else {
		*state += 4 
		return  nil
	}
} 

func get1(state *int, c chan interface{} )  (interface{}) {	
	ip, ok := <- c
	if ok {
		 *state += 2
		 return ip 
	} else {
		*state += 8
		return  nil
	}
} 

/* Collate compares IPs from two channels(0,1).  Matching IPs are 
   sent to channels 2 and 3, while mismatches from (0,1) are sent to (4,5) 
   respectively.
*/    
func Collate(wg *sync.WaitGroup, arg []string, cs []chan interface{}) {
	var state=0
	ip0 := *new(interface{})
	ip1 := *new(interface{})

	defer wg.Done()
	fmt.Println("Running", arg[0])
	 
	for {
		fmt.Println(state, ip0, ip1)
		switch  state { 
		case 0,2,8:   /* get 0 and add 1 to state. If EOF, add 4 to state */
				ip0 = get0(&state, cs[0]) 			
		case 1,4:  /* get 1 and state=3. if EOF state=8 */
				ip1 = get1(&state, cs[1]) 
		case 3:  
				 if ip0 == ip1 {
				 	cs[2] <- ip0    //BUG: Does not match all =s from 0.
				 	cs[3] <- ip1
				 	state=0
				 } else  {
				 	 val0 := reflect.ValueOf(ip0).Int()
					 val1 := reflect.ValueOf(ip1).Int()
					 if val0 > val1 {
					   cs[5]  <- ip1
				 	   state = 1 
					 } else   {
				 		cs[4] <- ip0
				 		state=2
				 }}
		case 6:  /*Output 1 to missed1.  */
				cs[5] <- ip1 
				state = 4
		case 9:  /*Output 0 to missed0. */
				cs[4] <- ip0
				state=8
		case 12: 
			close(cs[2]); close(cs[3]); close(cs[4]); close (cs[5]) 
			fmt.Println("Ended", arg[0])
			return
		
		}
	}
	
        
   /* ?? Cloase outputs */
}
