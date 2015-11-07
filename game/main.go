package main

import (
	"fmt"
	"time"
	_ "github.com/goodkele/mgnet/library/module/types"
	_ "github.com/goodkele/mgnet/library/module/proto"
)

func main() {
	
	dur,_ := time.ParseDuration("2s")
	
	tick := time.NewTicker(1 * time.Millisecond)
	
//	for now := range  tick.C {
//		fmt.Println(now)
//	}
	
	for {
		select {
			
			case now := <- tick.C :
			fmt.Println(now)
		}
	}
	
	fmt.Println("hello server")
}