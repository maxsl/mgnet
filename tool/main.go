package main

import (
	"fmt"
	"crypto/sha1"
	
//	"sync"
//	"flag"
//	"sync/atomic"
//	"github.com/goodkele/mgnet/tool/module/a"
//	"github.com/goodkele/mgnet/tool/module/b"
)


func main() {
	
	data := []byte("This page intentionally left blank.")
	
	by := sha1.Sum(data)
	
	//s := string(by[0:])
	
	fmt.Printf("%x\n", by[0:])
	fmt.Printf("% x\n", by[0:])
	
	//fmt.Println(s)
	
	//af064923bbf2301596aac4c273ba32178ebc4a96

	
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

