package epool

import (
	"time"
	"github.com/goodkele/mgnet/library/module/mglog"
	"github.com/goodkele/mgnet/library/module/link"
)

type Server struct {
	
	address 	string
	
	ptId		uint32
	serverId	uint32
	
	
	
	//updateChan	chan 
	
	sessions map[string]*link.Session
	
	updateTick	time.Ticker
	
}

// 创建服务
func NewServer(address string, ptId, uint32, serverId uint32) (*Server) {
	return &Server {
		address : address,
		
		ptId 	:	ptId,
		serverId:	serverId,
		
		sessions : make(map[string]*link.Session),
		
		serverTime	:	time.Now()
		
		
	}
}

// 接受请求
func (this *Server) Receive(session *link.Session) {

	routing := &types.Routing{}
	
	for {
		
		err := session.Receive(routing)
		if err != nil {
			mglog.Error(constant.ERROR_GATE_RECEIVE, err)
			// fixme 出现执行错误，应该把玩家踢下线
			continue
		}
		
		//routing.
		
	}
	
}

// 更新函数
func (this *Server) Update() {
	
}



// 开始提供服务
func (this *Server) Serve() {
	
	serve, err := link.Serve("tcp", this.address, &protocol.CodecType{})
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

			go this.Receive(session)
		}
	}()
	
	runtime.Gosched()
	
}

