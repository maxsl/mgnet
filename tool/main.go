package main

import (
	"fmt"
	"flag"
)

func main() {
	
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
}