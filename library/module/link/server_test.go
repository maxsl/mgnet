package link

import (
	"fmt"
	"io"
	"testing"
	"net"
	"github.com/goodkele/mgnet/library/module/protocol"
)

func Test_server(t *testing.T) {
	
	address := ":10011"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	server := NewServer(listener, &protocol.CodecType{})

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

	
	for i:=0; i<10; i++ {
		
	}
	
	
	server.Stop()
	

}