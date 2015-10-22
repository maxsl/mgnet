package main

import (
	"fmt"
//	"flag"
//	"sync/atomic"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}



func main() {
	
	
	var i Integer
	var ier LessAdder = &i

	i = 10
	i.Add(10)
	
	
	
	fmt.Println(ier.Less(20))
	
	
	
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

