Copyright (C) 2019 Thomas W. Young, fbp@twyoung.com 

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file or its derivitaves except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

STREAMWORK
==========

Name
----

         StreamWork - a flow based program(FBP) Go language  'backend' framework.
         

Description
-----------

StreamWork is a proof-of-concept for a Go language flow-based system which will be able to read and execute a network definition consisting of Go processes and links passing information packets(IPs) between them.  The IPs are designed as nil(empty) interfaces to be filled by the source components.    

The backend provided here has a minimal set of components, but enough to 
create, test, and benchmark working programs.  

The frontend(to be developed) will process the network definition.

Everything here(including the project name) is subject to change, until after the frontend has been developed.   There is no guarantee of backward compatibility.  

Contributors are encouraged.  Comments are particularly welcome.   
Please do not submit code before contacting the project by e-mailing streamwork@twyoung.com.     


Components
----------

Components are designed to be reentrant with a common interface.  The same 
component may be invoked any number of times, but  with a different, unique, process name each time.   Component memory will be reused.  Go creates a small stack for each process.  

The current list of components include 

	1. Comp1(Generates N IPs), 
	2. Comp2(sends input to the terminal), 
	3. Split(sends input to N output channels),  
	4. Concat(sends N inputs in order to its output channel), and  			
	5. Collate(compares IPs from channels 0 and 1,  
	  sends matches to channels 2 and 3 respectively, 
	  and mismatches to channels 4 and 5 respectively. 

Each process runs as a separate goroutine, which may possibly invoke more go routines.  
 
Why another FBP Golang framework?
---------------------------------

Streamwork will interpret and run a network definition text file.   No other 
project seems to do this.  Streamwork also defines a standard component
interface(currently subject to change), such that each process knows its name. 
      
Author
------

    Tom Young, streamwork@twyoung.com
