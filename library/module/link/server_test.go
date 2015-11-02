package link

import (
	"io"
	"testing"
	"net"
	"github.com/goodkele/mgnet/library/module/protocol"
	"sync"
	"github.com/goodkele/mgnet/library/module/types"
)

func Test_server(t *testing.T) {
	
	address := ":10011"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	
	server := NewServer(listener, &protocol.CodecType{})

	go func() {
		for {
			session,err := server.Accept()
			if err != nil {
				t.Errorf("%v", err)
				break
			}
			
			go io.Copy(session.Conn(), session.Conn())
		}
	}()


	
	var syncGroupWait sync.WaitGroup

	syncGroupWait.Add(1)

	go func () {
		sess, _ := Connect("tcp", ":10011", &protocol.CodecType{})
		
		sendReq := &types.SearchRequest{"name1", 1, 22}
		
		receiveReq := &types.SearchRequest{}

		
		sendErr := sess.Send(sendReq)
		if sendErr != nil {
			t.Errorf("Send : %v", sendErr)
		}
		
		sess.Send(sendReq)
		
		receiveErr := sess.Receive(receiveReq)
		
		if receiveErr != nil {
			t.Errorf("Receive : %v", receiveErr)
		}
		
		sess.Close()
		syncGroupWait.Done()
	}()

	syncGroupWait.Wait()
	
//	dur, _ := time.ParseDuration("4s")
//	time.Sleep(dur)

	isStop := server.Stop()
	if isStop != true {
		t.Errorf("Server : Can't stop")
	}
	
	

}