# go_sync

This set of simple exercises are intended to show how to use the functionalities
in the go programming for developing concurrent programs. These examples 
illustrate the goroutines, the race condition problem and how to fix using
to alternatives: locks and channels. 

The included examples are:
- "testCount": this happens when a set of goroutines accesses simultaneously 
by shared memory. This example is shown with a single counter accedes from
a number goroutines for adding it. It can be shown how when a number of 
goroutines access in concurrency, the expected counter value should be 10000000, 
although it is not. The result is random. For executing this example, use the 
comments described in the file.

- "testCountChannel": It is included a solution for the previous counter based
on GO channels. It can be checked that the result is correct, although the
execution time is greater than the non synchronized counter. 
The safe counter is in "counterSync"


- "testCountMutex": It is included a solution for the previous counter based
on GO locks. It can be checked that the result is correct, although the
execution time is greater than the non synchronized counter. This alternative
is faster than the channel one. The safe counter is in "counterSync".
