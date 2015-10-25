package main

import (
	"fmt"
	"sync"
//	"flag"
//	"sync/atomic"
//	"github.com/goodkele/mgnet/tool/module/a"
//	"github.com/goodkele/mgnet/tool/module/b"
)


func main() {
	
	
//	a.Say()
//	b.Say()

	fmt.Println("asdf")
	
	var syncMutex sync.Mutex
	
	syncMutex.Lock()
	
	syncMutex.Unlock()
	
	
	
	/*
	var gopherType string
	const (
        defaultGopher = "pocket"
        usage         = "the variety of gopher"
    )
	
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	
	
	flag.Parse()
	
	
	str := flag.Args()
	
	fmt.Println(str)
	
	fmt.Println(gopherType)
	
	fmt.Println("hello")
	
	var uu uint64
	atomic.AddUint64(&uu, 4)
	
	fmt.Println(uu)
	*/
	
	


	
	
	


	
}

