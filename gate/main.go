package main

import (
//	"net"
	"time"
	"runtime"
	"github.com/goodkele/mgnet/library/module/constant"
	"github.com/goodkele/mgnet/library/module/mglog"
	"github.com/goodkele/mgnet/library/module/link"
	"github.com/goodkele/mgnet/library/module/protocol"
	"github.com/goodkele/mgnet/library/module/types"
)

func init() {
	// 初始化日志
	mglog.InitLog("./", "gate", 0, &mglog.SWITCHER_DAY)
}


func receive(session *link.Session) {
	
	routing := &types.Routing{}
	
	for {
		
		err := session.Receive(routing)
		if err != nil {
			mglog.Error(constant.ERROR_GATE_RECEIVE, err)
		}
		
		//routing.
		
	}
}

func main() {

	address := ":10011"

	serve, err := link.Serve("tcp", address, &protocol.CodecType{})
	if err != nil {
		mglog.Error(constant.ERROR_GATE_SERVE, err)
		return
	}
	mglog.Info("Gate: Start server")

	go func() {
		mglog.Info("Gate: Start gate waiting accept")
		
		for {
			session, err := serve.Accept()
			mglog.Debug("Gate: accept, SessionId : %d", session.Id())
			if err != nil {
				mglog.Error(constant.ERROR_GATE_ACCEPT, err)
			}

			go receive(session)
		}
	}()

	
	runtime.Gosched()
	
	

	mglog.Debug("%v", serve)
	
	mglog.Info("assssdf")
	
	
	time.Sleep(5 * time.Second)

	
	
}