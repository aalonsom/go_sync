package counter

var value int = 0
// The channel must have a parameter type, even if it not required
var mutex = make (chan int)

func init() {
	value = 0
}

func Add() {
	mutex <- 1
}

func Get() int {
	return value;
}

func Set (i int) {
	value = i
}

func Start() {
	for{
		<- mutex
		value++
	}
}

