package main

import (
	"os"
	"fmt"
	"flag"
//	"sync/atomic"
//	"crypto/sha1"
//	"sync"
)

func main() {
	
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
	*/

	
	
	var flagSet1Name string
	var flagSet2Name string

	var flagSet1 = flag.NewFlagSet("name1", flag.ExitOnError)
	var flagSet2 = flag.NewFlagSet("name2", flag.ExitOnError)
	
	flagSet1.StringVar(&flagSet1Name, "name", "defaultName1", "pleat enter name 1")
	flagSet2.StringVar(&flagSet2Name, "name", "defaultName2", "pleat enter name 2")
	
	

	if os.Args[0] == "name1" {
//		fmt.Println("main name1")
		flagSet1.Parse(os.Args[1:])
	} else {
//		fmt.Println("main name2")
		flagSet2.Parse(os.Args[1:])
	}

//	flagSet1.Parse(os.Args[1:])
	
	fmt.Println(os.Args[0:]);
	
	fmt.Println(flagSet1Name)
	fmt.Println(flagSet2Name)
	
}

