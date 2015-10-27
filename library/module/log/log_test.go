package log

import (
	"testing"
	"time"
	"fmt"
)

func Test_All(t *testing.T) {
	
	
	fmt.Println(time.Now())
	
	dur,_ := time.ParseDuration("2s")
	
	time.Sleep(dur)
	
	fmt.Println(time.Now())

	fmt.Println(time.Now().Zone())


	switcher := &SwitcherHour{}
	
	du := switcher.NextDuration()
	
	fmt.Println(du)
	fmt.Println(switcher.Filename("hoho"))

	// dd,_ := time.ParseDuration("1s")

	tick := time.NewTicker(1 * time.Second)
	
	time.Sleep( 5 * time.Second)
	
	for t := range tick.C {
		fmt.Println(t)
	}
	
	tick.Stop()
	
//	c := time.Tick(1 * time.Second)
//	for now := range c {
//	    fmt.Printf("%v %s\n", now)
//	}
}