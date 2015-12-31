package link

import (
	"io"
	"testing"
	"sync"
	"github.com/goodkele/mgnet/library/module/types"
	"github.com/goodkele/mgnet/library/module/protocol"
	"runtime"
	//"net"
	//"time"
	//"fmt"
)

func Test_server(t *testing.T) {
	address := "192.168.1.21:10011"
	server, err := Serve("tcp", address, &protocol.CodecType{})
	if err != nil {
		t.Errorf("%v", err)
	}
	
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

	runtime.Gosched()
	
	var syncGroupWait sync.WaitGroup

	go func () {
		syncGroupWait.Add(1)
		defer syncGroupWait.Done()
		
		sess, err := Connect("tcp", "192.168.1.21:10011", &protocol.CodecType{})
		if err != nil {
			t.Errorf("%v", err)
		}
		
		t.Logf("%s", sess.Id())
		
		sendReq := &types.SearchRequest{"name1", 1, 22}
		
		receiveReq := &types.SearchRequest{}

		sendErr := sess.Send(sendReq)
		if sendErr != nil {
			t.Errorf("Send : %v", sendErr)
		}
		
		receiveErr := sess.Receive(receiveReq)
		
		if receiveErr != nil {
			t.Errorf("%v", receiveErr)
		}
		
		t.Logf("Receive :　%v", receiveReq)

		sess.Close()
	}()
	
	syncGroupWait.Wait()

	isStop := server.Stop()
	if isStop != true {
		t.Errorf("Server : Can't stop")
	}
}