#!/bin/bash

# FBPGO.sh
pgm=fbpgo
version="0.0.4"  ; 

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

##Tutorials 
	
###Quick Start

	* Download and install Go
		* Check GOPATH and GOROOT	
	* Create a directory in $GOPATH/mod.  Ex . mkdir $GOPATH/mod/foo
	* Change directory to $GOPATH/mod/foo
	* Create main.go with package foo and import "github/tyoung3/streamwork"
	* Run 'go mod init'
	* run 'go run'

### Ignore this 
	 
	* git clone https://github.com/tyoung3/streamwork 
	* Run test 'go test ./..
	* To update: 'git pull origin master' 
	
	
###Build a component 

####Component Skeleton Generation	
	* Run script with import path and component name(initial lower case letter). Ex. '... std foo' creates a new component, foo.go, and test, foo_test.go in .../std/   
	* Modify foo_test.go (test driven development) 
	* cd .../std;  run 'go test foo*'  
	* Modify foo.go 
	* repeat until working	
		
####Guidelines

	* defer wg.Done() at the beginning 
	* Test w/-race before committing
	* Avoid global variables
	* Include process name in messages (standard error code @TODO)
			  
##Trials

	* Try nesting fbp calls within a component.  
	* Try method calls.
	
##Roadmap 
	
###Frontend

	* Develop as a separate module
	
	* networkDefinition conversion to running program.  
		* networkDefinition > (AddSemiColons) > (expand variables)
		 > (tokeninize) > (interpret) > (buildGo) > (Launch) 
	
		  
###Backend 
	Backend prototype is working. 
 		  
		    	
##Signals
	. Trap signals
	
##Copyright 

##Git 
	

##Publish 
	
	* go mod tidy 					   [Clean up  go.sum]
	* go list [-u] -m all				. See all dependencies
	* go get -u=patch					. Update to latest patch version(s)


		Check branch!!
	* git status  [must be master branch]
	* git pull origin master
	* git add --all 
	* git commit -m "Good commit message"
	* git tag -a v1.2.3 -m "Version_Notice_wo_spaces"
	* git push origin v0.0.1

##Changes


#DONE

	* Run Go doc
	* Replace Launch4
	* Collate component
	
##GitHub Flow

###New branch 

	* git pull origin master 
	* git checkout -b "Usefull Branch Name"
	* git push origin "Usefull Branch Name"
		   https://github.com/tyoung3/streamwork/pull/new/Fix_Comp_names
	* Make changes
	* git push --set-upstream origin Fix_Comp_names
	* git branch -a   [Check branch]
	
	
	* git commit -m "Good commit message"
	* git add --all
	* git status
	* git push origin Fix_Comp_names
	
	
https://githubflow.github.io/	
	
	* Anything in the master branch is deployable
	* To work on something new, create a descriptively named branch 
	  off of master (ie: new-oauth2-scopes)
	* Commit to that branch locally and regularly push your work 
	  to the same named branch on the server
	* When you need feedback or help, or you think the branch is 
	  ready for merging, open a pull request
	* After someone else has reviewed and signed off on the feature, 
	  you can merge it into master
	* Once it is merged and pushed to 'master', you can and 
	  should deploy immediately
	
	
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
EOF
	
}

Notesx() { 
	cat <<- EOF
		Workflow: https://gist.github.com/blackfalcon/8428401
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

src=~/streamwork

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
	

SeeDocs() {

	URLS="
	https://gist.github.com/blackfalcon/8428401/
	https://appliedgo.net/flow/ 
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
	pkg=$1; shift; 
	name=$1
	inp=$2
	outp=$3
	shift 3
	[ -z $pkg ] && Die  Usage: $0 gs PKG NAME ... 
	[ -z $name ] && Die Usage: $0 gs PKG NAME ...
	[ -d $src/$pkg ] || Die $src/$pkg is missing.
	src2=$src/$pkg
	pushd $src2 || Die Cannot cd  $src2
	uname=name
	[ -f ${name}.go ]  || cat << EOF > ${name}.go 
	 	package $pkg
	 	
		`IdSkel` 
		 
		import "fmt"
		import "sync"
		
		var Version string="v$version"
		
	    func $uname(wg *sync.WaitGroup, arg []string, cs []chan interface{} ){
	    	
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
		package $pkg

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
	 for  {
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
	go $uname(&wg, arg,  cs)
	wg.Wait()	

}
	
EOFY
}


CheckOut() {
	go get github.com/tyoung3/dummy_repository 
}

case $1 in	
	d)  shift; SeeDocs --new-window  $*;;
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
	 	gd			. Run and view 'godocs'
	 	gs PKG	COMPNAME [INPORTS [OUTPORTS]] . Generate skeleton components
					COMPNAME must begin with an upper case letter
					PKG directory mus exist
	 	n			. Browse $pgm notes
	 	r			. Run fbpgo.go
	 	t			. Test fbp packages 
	 	v			. Display version
	 	x			. Edit $0
	 	--help			. Display this usage
		
	 	Example: $0 x
EOFX
	;;
esac	

