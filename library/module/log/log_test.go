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
	
}