package std

import "sync"

// var Version string = "v0.0.1"

func Concat(wg *sync.WaitGroup, 
			arg []string, 
			cs []chan interface{}) {

	cout := cs[0]
	
		for j := range(cs[1:]) {
			func(cin chan interface{}) {
			   for {
				 ip,ok := <- cin
				 if ok != true {
				 	 return
				 }
				 cout <- ip
			   }	 
			}(cs[j+1])	
		}	

	close(cout)
	wg.Done()
}
