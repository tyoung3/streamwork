Copyright (C) 2019 Thomas W. Young, streamwork@twyoung.com 

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file or its derivitaves except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[![Go Report Card](https://goreportcard.com/badge/tyoung3/streamwork)](https://goreportcard.com/report/tyoung3/streamwork)
[![GoDoc](https://godoc.org/github.com/tyoung3/streamwork?status.svg)](https://godoc.org/github.com/tyoung3/streamwork)

STREAMWORK-0.3.0 
================

Name
----

         StreamWork - a flow based, Go language framework.
         

Description
-----------

StreamWork is a proof-of-concept for a Go language 
flow-based system which reads and executes a network 
definition consisting of goroutines and channel connections
between them.  
The channels pass information packets(IPs) designed as 
nil(empty) interfaces.     

The network definition is a simple, graphable way of 
clearly describing process connections. Sw creates the code
that makes and assigns the channels, and lauches the goroutines.   

The backend provided here has a minimal set of components, but enough to 
create, test, and benchmark working programs.  

The frontend, sw, processes the network definition and 
generates the Go code.  Sw can also produce an .svg or 
.jpeg image of the network definition.

Versions will be backward compatible within the same major version. 
Ex. your code depending on v0.0.1 will still work on v0.8.7, 
but may fail on v1.0.0.  
Streamwork is developed using the Go module facility, 
so that dependencies are clear.   

Contributors are encouraged.  Comments are particularly welcome.   
Please do not submit code before contacting the project by e-mailing streamwork@twyoung.com.     


Components
----------

Components are designed to be reentrant with a common interface. 
The same component may be invoked any number of times, but 
with a different,
unique, process name each time.   Component memory will be reused. 
Go creates a small stack for each process.  

An incomplete list of components includes:

	1. Gen1(Generates N integer IPs), 
	2. Gen2(Generates N string IPs),
	3. Print1(sends input to the terminal), 
	4. Split(sends input to N output channels),  
	5. Concat(sends N inputs in order to its 
	   output channel)		  			
	6. Collate(compares IPs from channels 0 and 1,  
	   sends matches to channels 2 and 3 respectively, 
	   and mismatches to channels 4 and 5 respectively. 
	   Note that channels 2 and 3 may actually and usefully be the same channel.
	   Other combinations would not be so useful.
	    
Each process runs as a separate goroutine, which may well invoke more goroutines.  
In such cases, the component is responsible for waiting until all its goroutines are
finished, before returning.
 
QuickStart
----------

	* Download and install Go [v1.11.0 or better is best. 
	  1.9.7 and up may also work. 
	  Previous versions will need to run under $GOPATH/src 
	  without module/version support].   	
	* Create a new directory, $GOPATH/mod/foo.  
	* Change directory to $GOPATH/mod/foo.  
	* Create $GOPATH/mod/foo/main.go:
```	
package  main

import "fmt"
import "sync"
import "github.com/tyoung3/streamwork/fbp"
import "github.com/tyoung3/streamwork/def"

func main() {
        var cs []chan interface{}
        var wg sync.WaitGroup
        
        fmt.Println("foo Start")
        cs = append(cs, make(chan interface{}))

        fbp.Launch(&wg, []string{"TestPrint1"},    def.Print1, cs)
        fbp.Launch(&wg, []string{"TestGen1","11"}, def.Gen1, cs)
        wg.Wait()
        fmt.Println("foo end")
}

```
	* Run 'go mod init foo/foo'
	  This should create .../mod/foo/go.mod 

	* Run 'go run main.go'
		Go should find and install 
	"github.com/tyoung3/streamwork/strings latest"; 
	 creating .../mod/foo/go.sum, 
     and compile and run the program, Foo. 

    * Run 'go build' to create the executable, foo (or foo.exe). 
    
Why another FBP Go framework?
---------------------------------

	* Streamwork is able to interpret and run a program
	  generated from a network definition text file. 
	  No other project seems to do this, so far. 
	  
	* Streamwork ensures clear component interfaces.  
	
	* Streamwork components are contained in separately compiled 
	  members of separately tested packages, rather than being 
	  mashed together in a single program.

	* Streamwork defines a standard component 
	  interface, such that each
	  process knows its name; very useful for 
	  reporting errors, etc.. 
	  
	* Streamwork uses precompiled code(unless the source
	  is changed or recently imported).  
	  Compilation, when it is required, is 
	  transparent unless errors occur.
	  
      
Author
------

    Tom Young, streamwork@twyoung.com
