package edit

/*  Generated StreamWork component, View,
    by sw.sh on Wed Jun 26 19:56:37 EDT 2019
	1  input  ports
	1 output ports
*/

import "testing"
import "fmt"
import "sync"

func TestSkel_View(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup
	var domore bool = true

	arg := []string{"TestSkel_View"}

	fmt.Println("\u001b[32m",arg[0],"\u001b[0m")
	for i := 0; i < 1+1; i++ {
		cs = append(cs, make(chan interface{},1))
	}

	wg.Add(2)
	go func() {
	
	  i := 3
	  
	  for domore   {
		select {
		case cs[0] <- i+100: 
			i--
			if i<1 {
				fmt.Println("\u001b[32mclosing\u001b[0m")
				close (cs[0])
			}	 
		case ip, ok := <-cs[1]:
			if ok != true {
				domore = false
			} else {
				fmt.Println("\u001b[32mchan:", 1, 
					"IP:", ip,"\u001b[0m")
			}
		}
	  }	
		
/*
		i := 0
		for i >= 0 {
			cs[i] <- i
			close(cs[i])
			i--
		}
		
		j := 1
		for j >= 1 {
			ip, ok := <-cs[j]
			if ok != true {
				break
			}
			fmt.Println("chan:", j, "IP:", ip)
			j--
		}
*/
		fmt.Println("TestSkel_View Ended")
		wg.Done()
		return
	}()

	go View(&wg, arg, cs)
	wg.Wait()

}
