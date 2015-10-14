package link

import (
	"fmt"
	"io"
	"testing"
	"net"
)

func Test_server(t *testing.T) {
	
	address := ":10011"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	server := NewServer(listener)

	go func() {
		for {
			session,err := server.Accept()
			if err != nil {
				fmt.Println(err)
				break
			}
			
			go io.Copy(session.Conn(), session.Conn())
		}
	}()

	
	server.Stop()
	

}