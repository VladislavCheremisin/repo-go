package main

import (
	"fmt"
	"github.com/pkg/profile"
	"runtime"
)

func sayHi(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sayBay(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	go sayHi("Hi people")
	go sayBay("Bay man")
	say("Wow")
}
