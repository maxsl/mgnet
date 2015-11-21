package epool

import (
	"time"
	"runtime"
	"sync"
	"github.com/goodkele/mgnet/library/module/mglog"
	"github.com/goodkele/mgnet/library/module/constant"
	"github.com/goodkele/mgnet/library/module/link"
	"github.com/goodkele/mgnet/library/module/types"
	"github.com/goodkele/mgnet/library/module/protocol"
)

//var (
//	GameServer *Server
//)

type Connect struct {
	connId	uint32
	address string
	status	int8
	conn 	*link.Session
	
}

func NewConnect(connId	uint32, address string) {
	
}

// 链接到其他服务
func (this *Connect) connect() {
	
}

// 收到
func (this *Connect) receive() {
	
}

// 发送
func (this *Connect) send() {
	
}



type Server struct {
	address 	string
	ptId		uint32
	serverId	uint32
	sessions 	map[string]*link.Session
	conns		map[uint32]*link.Session
	
	receiveChan	chan *types.Routing		// 收到消息
	sendChan	chan *types.Routing		// 发送消息
	updateTick	*time.Ticker
	SyncGroupStop 	sync.WaitGroup		// 等待锁，服务器关闭时等待所有session关闭

	now			time.Time
}

// 创建服务
func NewServer(address string, ptId uint32, serverId uint32) (*Server) {
	
	return &Server {
		address 	: 	address,
		ptId 		:	ptId,
		serverId	:	serverId,
		sessions 	: 	make(map[string]*link.Session),
		conns		: 	make(map[uint32]*link.Session),
		
		receiveChan	:	make(chan *types.Routing, 1024),
		sendChan	:	make(chan *types.Routing, 1024),
		updateTick	:	time.NewTicker(10 * time.Millisecond),	// 1秒更新100次
		
		now			:	time.Now(),
	}
}

// 接受请求
func (this *Server) Receive(session *link.Session) {
	routing := &types.Routing{}
	
	for {
		err := session.Receive(routing)
		if err != nil {
			mglog.Error(constant.ERROR_RECEIVE, "gate", err)
			// fixme 出现执行错误，应该把玩家踢下线
			session.Close()
			break
		}
		
		this.receiveChan <- routing
	}
}

// 收到消息
func (this *Server) RoutingReceive(routing *types.Routing) {
	this.sendChan <- routing
}

// 发送消息
func (this *Server) RoutingSend(routing *types.Routing) {
	
	
	
	this.conns[constant.SERVER_GAME].Send(routing)
}

// 更新函数
func (this *Server) Update(now time.Time) {
	this.now = now

}

// 路由
func (this *Server) Routing() {
	for {
		select {
			case receive := <- this.receiveChan :
				this.RoutingReceive(receive)
			case send := <- this.sendChan :
				this.RoutingSend(send)
			case now := <- this.updateTick.C :
				this.Update(now)
		}
	}
}

// 开始提供服务
func (this *Server) Serve() {
	serve, err := link.Serve("tcp", this.address, &protocol.CodecType{})
	if err != nil {
		mglog.Error(constant.ERROR_SERVE, "gate", err)
		return
	}
	mglog.Info("Gate: Start server")
	this.SyncGroupStop.Add(1)
	
	go func() {
		mglog.Info("Gate: Start gate waiting accept")
		
		for {
			session, err := serve.Accept()
			mglog.Debug("Gate: accept, SessionId : %s", session.Id())
			if err != nil {
				mglog.Error(constant.ERROR_ACCEPT, "gate", err)
			}

			go this.Receive(session)
		}
	}()
	
	runtime.Gosched()
	
	go this.Routing()
}

// 关闭服务器 
func (this *Server) Stop() {
	this.SyncGroupStop.Done();
	
}

// 链接到其他服务器
func (this *Server) Connect(connType uint32, address string) {
	for {
		session, err := link.Connect("tcp", address, &protocol.CodecType{})
		if err != nil {
			mglog.Error(constant.ERROR_CONNECT, "gate", connType, err)
			time.Sleep(3 * time.Second)
			continue
		}
	
		mglog.Info("Gate: Connect %d, address:%s", connType, address)
		
		this.conns[connType] = session
	
		msg := &types.Routing{1,2,3,"sessId1", []byte{}, 4, 5, "error6" }
	
		mglog.Debug("%v", msg)
	
		this.conns[connType].Send(msg)
	
		go this.Receive(session)
		break
	}
}

