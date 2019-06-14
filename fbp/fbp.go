/*Package fbp contains code to launch Streamwork generated Go programs
sw creates a Go program from a Streamwork network definition.  The 
generated Go  program uses package fbp to launch the components specified 
in the  network definition.  
*/
package fbp

/*
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
*/

/*
TODO(ty): Document  code

*/

import "sync"

/*Launch starts a standard fbp process as a goroutine.

Launch requires four arguments:

1) wg, a pointer to a wait group.
   Launch increments the WaitGroup for every process
   launched.The process
   must import "sync" and issue wg.Done()  immediately
   prior to returning. This is easily and reliably
   accomplished by issuing 'defer wg.Done()' at the
   start of the process.

2)arg, a slice of argument strings.

    THe first string is
	required to be the unique process name.  Developers should
	strive to standardize argument meanings within their
	own packages.

3)f, the component's function signature.

	Launch passes on wg, arg, and cs to the process.
	Packages using Launch will need
	to import the function's package.  (Ex. to launch split, import
	"fbp/std" and specify  fbp.Launch(&wg,arg, std.Split,cs).)
	A bash script, 'fbpgo.sh -gs', is available to generate
	component skeleton code.

4)cs. a slice of channel interfaces.

    cs defines a slice of input channels followed by
	output channels, corresponding to in1, in2,...out1, out2 ...
	or 1,2,3,... ports in a network definition.

	Components are required to close all their output channels,
	and none of their input channels.

	Interfaces allow processes to exchange type-verified data
	over the channel.  Data may be of any valid Go type,
	so long as the receiving channel accepts it.
	Receiving processes may use select type
	instructions to handle multiple interface types, if desirable.

	It is possible to bypass Launch, by incrementing wg  and
	invoking the component directly.  This can be useful for
	independent testing, but is not supported.

*/
func Launch(
	wg *sync.WaitGroup,
	arg []string,
	f func(wg *sync.WaitGroup,
		arg []string,
		cs []chan interface{}),
	cs []chan interface{}) {
	/*if len(arg) < 1 {
		os.exit(2)
	} */
	wg.Add(1)
	go f(wg, arg, cs)
	/* DO NOT put code here.  It will most likely not run.*/
}
