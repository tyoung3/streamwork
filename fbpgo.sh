#!/bin/bash

# FBPGO.sh
pgm=fbpgo
version="0.0.1"  ;# v0.0.0

# NOTE: Go ignores files and directories beginning with '_', as in _OLD/

Die() {
		echo "$0/Die $*"
		exit 1
}

Verbose() {
	if [ ! -z $verbose ]; then
	 	echo "${pgm}.sh: $*"
	fi 	
}

NotesFBP() {
	cat << EOF
* * * * * * * * * * * 
#TODO

##Upload to github/tyoung3

	* Find good name
		* fgbase, flogo, flow, flowcontrol,goflow, flow2go, appgo (taken)
		 by trustmaster, etc.
		 flo, flow,  dataflow, cogoflow 
		 graph, chart, flowgraph
		 lego, plumbing, snap together, tab/slot, link, plug/socket 
		 goplug, gograph 
		 GOLINK, glink, gflow, fountain , fount, 
		 port, goport, node, flowport, PORTAGO
		 river, delta, oxflow,  
		 pipework, plumbery, pipeline, tubing, canals, 
		 STREAMWORKS, bitstream, bitflow, data current 
		 GoG(Go Graph)  [noun: Haste; ardent desire to go.]
		 net, network,  nodelink 
		 
	* Fix copyright
	* Fix README.md 
	* git push (add to fbpgo.sh)

##Trials
	. Try nesting fbp calls in a component. Callit Nest
	. Try method calls.
	
##Roadmap 
	###Collate component
	
	###Frontend
	. .net conversion to running program.  
		. .net > (AddSemiColons) > (expand variables) > (tokeninize) >
		  (interpret) > (buildGo) > (Launch) 
		  
##Guidelines		  
		    	
##Signals
	. Trap signals
	
##Copyright 

##Publish 

##Run Go doc

##Changes

#$pgm
What follows are some thoughts on designing a comprehensive FBP-based Go language program, $pgm.  

Such a system makes possible the direct invocation of a program from its editable flow graph or from the flow graph's stored network text file. 

A comprehensive system would allow a wide community of developers to  share code effectively.  

##Goroutines 
Goroutines are lightweight(green) threads communicating with each other within a single address space.  Go may schedule goroutines onto multiple operating system threads and therefore onto multiple CPU cores. 
 
Go provides a number of facilities to help solve and to avoid race and deadlock problems.  

##Channels
Go design is very FBP-like.  Goroutines typically communicate over
Go channels, which are ring buffers of specified capacity(default 1).  
Channels and goroutines are built into the Golang syntax(unlike thread libraries in other languages) clarifying and simplifying their operation.
Channels can help to avoid extensive locking of critical code sections. 

##Packaging
The basic Go unit is a package, which is imported into a main Go program or another Go package.  The single main Go program defines a Go command, an executable program.

Go 'modules' are collections of related packages.  Modules will default
with the advent of Go-2.0.  Go modules are a recent development. 
Go 'plugins' allow for third parties to create interfacing code. 
The Go ecosystem avoids complex make files by the way it arranges dependent 
code in sub-directories.  Package libraries are compiled into the ../pkg directory tree. 
$pgm uses modules to allow for third parties to develop $pgm code 
independently. 

##Performance
Much has been done to make Go programs fast.  
A benchmark go program, available at ./github.com/tyoung3/flowfib, can run millions of goroutines communicating over millions of channels in a few seconds.

##Implementation 
From a graphable text file, FOO.net (which may include other files
and reference subnets defined elsewhere), generate and run a Go
'main' program, FOO.go.  FOO.go will import and invoke standard $pgm go components, passing a structure(interface), providing the component's process name, interface(s), and string arguments.  

###Components
Each component is required to be reentrant.  The same component may be invoked multiple times in the network definition, but each time with a different name. 
Components should be pre-compiled or compiled on the fly and cached.

###Ports
FOO.go will connect component ports as specified in FOO.net.   FOO.go will intercept OS signals, SIGUSR1, SIGUSR2, etc. and act accordingly.  

###Nodes
FOO.go will also connect with external nodes and processes as specified(or not) in FOO.go.

#Done
	Replace Launch4
	
	
EOF
	
}

Notesx() { 
	cat <<- EOF
		Code generation from BNF: https://github.com/goccmack/gocc
		Using gocc: https://medium.freecodecamp.org/write-a-compiler-in-go-quick-guide-30d2f33ac6e0
	Tutorial: https://tour.golang.org/welcome/1
	Specification: https://golang.org/ref/spec#Notation
		http://www.cs.sfu.ca/CourseCentral/383/tjd/	syntaxAndEBNF.html#defining-language-syntax-with-extended-backus-naur-form

FLOWGO:  
	Expand DVD  (A go.comp1)out->in(B go.comp2)out=>in(S Stdout); 
		This will generate two system processes: 1)a GO process running
		two goroutines, and 2)S, which sends B output to the console. 
	Generate a FBP in GO from a go.net definition.  
		. Use GOIMPORTS to generate inport lines.
		. Generate a Go channel for each flow
		. Connect components using goroutines
		. SELECT to schedule goroutines 
		. ? provide for connections to external processes.
EOF

}

src=~/go/mod/fbp

IdentSrc() {
	cat << EOF
/* $src2
 *		 	 DO NOT EDIT!!!
 *		generated by ${pgm}.sh-v$version  on `date`
**/			 	 
EOF
	
}
	
MakeDir2() {
	dir2=$1
	ddate=`date`
	Verbose Making $dir2
	[ -d $dir2 ] 		\
	|| ( mkdir -p $dir2		\
	&& cat << EOF > $dir2/_README.md
	$dir2	generated by ${pgm}.sh-v$version on $ddate	 
EOF
	)
}

MakeDir() {
	for dir in $*; do 
		MakeDir2 $dir
	done
}
	
Genfbpgotest() { 
	gsrc=$src/${pgm}_test.go	
	cat <<- EOF > $gsrc
	package main

		`IdentSrc`

	import (
		"testing"
		"fmt"
		"time"
		"fbp/strings/comp1"
		"fbp/comp2"
	)

	func TestFBP(t *testing.T) {
			var cs []chan interface{}
			
			fmt.Println("Welcome to test $pgm!")
			
		
			d1 := make(chan struct{})
			d2 := make(chan struct{})	
			cs = append(cs,make(chan interface{}))
			cs = append(cs,make(chan interface{}))
			go comp1.Launch(d1, _, "A1",cs[0:1])
			go comp2.Launch(d2,"C2",cs[1:2])
		    <- d1
		    <- d2  		
			fmt.Println("The time is", time.Now())
	}
EOF
	go fmt $gsrc > /dev/null
}

Gen() {
	cd   $src   
	[ -f ${pgm}.sh ] || Die  
	[ -f go.mod ] || go mod init fbp 
	
	src2=$src/fbp_test.go
	cat << EOF >  $src2
		package fbp
		
		`IdentSrc`
		
		import "testing"
		import "sync" 
		
	
		func dummyComponent2(wg *sync.WaitGroup,
							 arg []string, 
							 cs []chan interface{}) {
			defer wg.Done()	
			close(cs[0]);			 
		}
		
		func TestLaunch(t *testing.T) {
			var cs2 []chan interface{}
			var wg sync.WaitGroup
			
			arg := []string{"x2"}
			// done2 := make(chan struct{})
			cs2 = append(cs2,make(chan interface{}))
    	    Launch4(&wg, arg, dummyComponent2, cs2) 
			wg.Wait()			    
    	    
    	}	
EOF
	go fmt $src2 > /dev/null
	
	src2=$src/fbp.go
	cat << EOF > $src2
	 	package fbp
	 	
		`IdentSrc`
		
		import "fmt"
		import "sync"

	 	func Launch4(
	 			wg *sync.WaitGroup, 
	 			arg []string,
	 			f func( wg *sync.WaitGroup,
	 					arg []string, 
	 					cs []chan interface{}),
	 			cs []chan interface{} ) {
	 	  			
	 		fmt.Println("fbp: Launching",arg[0])
	 		wg.Add(1)
	 	  	go f(wg, arg, cs);			
		}		
		
	 	func Launch3(
	 			arg []string,
	 			f func(arg []string, cs []chan interface{}),
	 			cs []chan interface{} ) {
	 	  		
	 	  	done3 := make(chan struct{},0)	
	 		fmt.Println("fbp: Launching",arg[0])
	 	  	f(arg, cs);			
			close(done3)
		}		
		
	 	func Launch2(
	 			done2 chan struct{},
	 			arg []string,
	 			f func(arg []string, cs []chan interface{}),
	 			cs []chan interface{} ) {
	 	  					
			defer close(done2)
	 		fmt.Println("fbp: Launching",arg[0])
	 	  	f(arg, cs);	
		}		
	 	 	
EOF
	go fmt $src2 > /dev/null
	
	MakeDir $src/internal/merge
	src2=$src/internal/merge/merge_test.go
	cat  << EOF > $src2
package merge

`IdentSrc`

import "testing"
import "fmt"

func TestMerge(t *testing.T) {
	var cs []chan interface{}
	 
	cs = append(cs,make(chan interface{}))
	cs = append(cs,make(chan interface{}))
	cs = append(cs,make(chan interface{}))
	cs = append(cs,make(chan interface{}))
	cs = append(cs,make(chan interface{}))
	c1 := cs[0]
	c2 := cs[1]
	c3 := cs[2]
	c4 := cs[3]
	
	fmt.Println("TestMerge Start ")
		go func () {
			c1 <- "A:IP1"
			c3 <- "C:IP1"
			c1 <- "A:IP2"
			c4 <- "D:IP1"
			c2 <- "B:IP1"
			c2 <- "B:IP2"
			close(c1)
			close(c2)
			close(c3)
			close(c4)
		}()
			
	Launch( cs[0], cs[1], cs[2], cs[3], cs[4] )
	fmt.Println("TestMerge")
	
	go func() {
		for   {
			i, ok := <- cs[0]
			if ok != true {
				return
			}
			fmt.Println(i)
		}	
	}()

	fmt.Println("TestMerge Ended")
}
EOF
	go fmt $src2 > /dev/null
	
	src2=$src/internal/merge/merge.go
	cat  << EOF > $src2
package merge

import "sync"

`IdentSrc`

func Launch( cs ...chan interface{} ) chan interface{} {
	var wg sync.WaitGroup
	
	wg.Add(len(cs)-1)
	
	for _, c := range cs[1:] {
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				cs[0] <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(cs[0])
	}()
	
	return cs[0]
}	
EOF
	go fmt $src2 > /dev/null
	
		
	MakeDir $src/strings
	src2=$src/strings/comp1_test.go
	cat << EOF > $src2
		package strings

`IdentSrc`

import "testing"
import "fmt"
import "sync"

func TestComp1(t *testing.T) {
	var cs  []chan interface{}
	var wg  sync.WaitGroup
	
	arg := []string{"TestComp1","7"}
	
	fmt.Println(arg[0])
	cs = append(cs,make(chan interface{}))
	c  := cs[0]
	
	go func() {
	 for 1 == 1 {
		s, ok := <-c
		if ok == true {
			fmt.Println("TestCompcomp1.go1", s)
		} else {
			fmt.Println("TestComp1 Ended")
			return
		}
	  }	
	}() 
	
	wg.Add(1)
	go Comp1(&wg, arg,  cs)
	wg.Wait()		
	// <- done

}
EOF
	go fmt $src2 > /dev/null
	
	src2=$src/strings/comp1.go
	cat << EOF > $src2
	 	package strings
	 	
		`IdentSrc`
		 
		import "fmt"
		import "sync"
		import "strconv"
		
	    func Comp1(wg *sync.WaitGroup, arg []string, cs []chan interface{} ){
	    	var n0 int=7
	    	var n  int 
	    	
	    	
	    	defer wg.Done()
	    	//pname := arg[0]
	    	l := len(arg)
	    	if l > 1 {
				n0, _ = strconv.Atoi(arg[1]) 
			}	
			
	    	c := cs[0]
	    	
			for n=1; n<n0+1; n++ {
				c <- fmt.Sprintf("%s-%d",arg[0],n)
			} 
			
			close(c)
	    }
	 	 	
EOF
	go fmt $src2 > /dev/null

	MakeDir $src/comp2 
	src2=$src/comp2/comp2_test.go 
	cat <<- EOF > $src2
package comp2

`IdentSrc`

import "testing"
import "fmt"
import "sync"

func TestComp2(t *testing.T) {
	var cs []chan interface{}
	var wg sync.WaitGroup
	arg :=[]string{"TestComp2", "xx"}
	
	fmt.Println("TestComp2")
	cs = append(cs,make(chan interface{}))
	c := cs[0]
	
	go func() {
		c <- 7
		c <- 11
		c <- "OK?"
		close(c)
	}()
	
	wg.Add(1)
	go Launch(&wg, arg,  cs)
	fmt.Println("TestComp2 wait c ")
	wg.Wait()
	fmt.Println("TestComp2 done.")
}
EOF
	go fmt $src2 > /dev/null
		
	src2=$src/comp2/comp2.go
	cat <<- EOF > $src2
	 	package comp2
	 	
		`IdentSrc`
		
	 	import "fmt"
	 	import "sync"
	 	
	    func Launch(wg *sync.WaitGroup, 
	    			arg []string, 
	    			cs []chan interface{}) {
	    	defer wg.Done()		
	    	c:=cs[0]
	    	s:=arg[0]
			for 1==1 {
	    		ip, ok := <- c
	    		if ok != true {
	    			return
	    		}
	    		switch  ip.(type) {
	    			case string:
	    				fmt.Println(s,ip)
	    			case int:
	    				fmt.Println(s,"Int:",ip)
	    			default:
	    				fmt.Println(s,"type mismatch")
	    		}			
	    	}	
		}
	 	 	
EOF
	go fmt $src2 > /dev/null
	go build
	# Genfbpgotest	
} 
SeeDocs() {

	URLS="https://appliedgo.net/flow/ 
	https://golang.org/
	https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
	http://reo.project.cwi.nl/v2/projects/
	http://scipipe.org/
	https://www.commonwl.org/
	"
	
	$BROWSER $* $URLS
	go doc cmd/cgo |less
	go help c |less 
	go --help |less
}

HTML=fbpgo.html


IdSkel() {
	cat << EOF
	/*  Generated skeleton fbp component by fbpgo.sh on `date` 
		input  ports: "$inp"
		output ports: "$outp"
	*/ 
EOF
}

#  Generate a skeleton component
GenSkel() {
	name=$1
	inp=$2
	outp=$3
	shift 3
	[ -z $name ] && Die Usage: $0 gs NAME ...
	[ -f $src/$name ] && Die $src/$name exists 
	MakeDir $name && pushd $name || Die Cannot make directory $name
	[ -f ${name}.go ]  || cat << EOF > ${name}.go 
	 	package $name
	 	
		`IdSkel` 
		 
		import "fmt"
		import "sync"
		
		var Version string="v$version"
		
	    func Launch(wg *sync.WaitGroup, arg []string, cs []chan interface{} ){
	    	
	    	defer wg.Done()
	    	fmt.Println("Running",arg[0])
	    	c := cs[0]
	    	c <- "out1 IP1"
	    	c <- "out1 IP2"
			close(c)
	    }
	    
EOF
	go fmt ${name}.go
	
	[ -f ${name}_test.go ]  ||cat << EOFY > ${name}_test.go
		package $name

	`IdSkel`
	
import "testing"
import "fmt"
import "sync"

func TestSkel_$name(t *testing.T) {
	var cs  []chan interface{}
	var wg  sync.WaitGroup
	
	arg := []string{"$name","dummy arg[1]","dummy arg[2]"}
	
	fmt.Println(arg[0])
	cs = append(cs,make(chan interface{}))
	c  := cs[0]
	
	go func() {
	 for 1 == 1 {
		s, ok := <-c
		if ok == true {
			fmt.Println("Out1", s)
		} else {
			fmt.Println("Test_$name Ended")
			wg.Done()
			return
		}
	  }	
	}() 
	
	wg.Add(2)
	go Launch(&wg, arg,  cs)
	wg.Wait()	

}
	
EOFY
}


CheckOut() {
	go get github.com/tyoung3/dummy_repository 
}

case $1 in	
	d)  shift; SeeDocs --new-window  $*;;
	#####g)  Gen;; 
	####g2) Genfbpgotest;;
	gd)shift;   godoc -http=:6060 &
		pushd ../; $BROWSER taos:6060/streamwork
		;;
	gs|skel)shift;GenSkel $*;;
	n)shift; NotesFBP | pandoc -s --toc  -o  $HTML	\
	&& 	$BROWSER --new-window $* $HTML;;
	r)  go run internal/$pgm/*.go;;  #/* */
	t)  shift; go test $* ./... 2>&1 |less ;;
	v)	echo $pgm-v$version;;
	x)  $EDITOR ${pgm}.sh internal/${pgm}/${pgm}.go &;;
	*)  cat <<- EOFX
	
	$0-v$version USAGE:  
	 	d			. View GO docs 
	 	g			. Generate GO source code and run
	 	gd			. Run and view 'godocs'
	 	gs 		    . ? Gen fbpgo test??
	 	gs|skel	NAME [INPORTS [OUTPORTS]] . Generate skeleton components
	 	n			. Browse $pgm notes
	 	r			. Run fbpgo.go
	 	t			. Test fbp packages 
	 	v			. Display version
	 	x			. Edit $0
	 	--help			. Display this usage
		
	 	Example: $0 x
EOFX
	;;
#	c)  shift; CheckOut $*;;
#	l|list)shift; pushd ~/go/src/cmd >/dev/null; ls;;
#	r)  cmd=$2;
#		[ -d ~/go/src/cmd/$cmd ] || Die Missing ~/go/src/cmd/$cmd
#		shift 2; cd ~/go/src/cmd/$cmd; go run *.go $*;; 
#	 	c					. Check out flowfib from github
#	 	l					. List all commands in ~/go/src/cmd
#	 	r CMD [CMD options] . Run all go code in ~/go/src/cmd/CMD
esac	

