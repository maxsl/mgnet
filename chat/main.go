package main

import (
	"fmt"
)

func main() {

	var p []byte
	p = make([]byte, 1024, 1024)
	
	p[0] = 'a'
	p[1] = 'b'
	
	fmt.Println(p)
	
	p = p[0:0]

	fmt.Println("hello chat")
}