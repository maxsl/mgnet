package link

import (
	"net"
	"sync"
	"sync/atomic"
	"github.com/goodkele/mgnet/library/module"
)

type Server struct {
	listener  		net.Listener		// 监听
	codecType 		module.CodecType	// 创建器

	stopFlag 		int32				// 服务器是否暂停
	syncGroupStop 	sync.WaitGroup		// 等待锁，服务器关闭时等待所有session关闭

	sessions     	map[uint64]*Session	// Sessions
	syncMutexSession sync.Mutex			// 锁，创建与销毁session时
	
	online			uint64				// 当前在线人数
}

// 创建服务器
func NewServer(listener net.Listener, codecType module.CodecType) *Server {
	server := &Server{
		listener 	: listener,
		codecType 	: codecType,
		sessions 	: make(map[uint64]*Session),
	}
	return server
}

// 等待一个连接
func (this *Server) Accept() (*Session, error) {
	conn, err := this.listener.Accept()
	if err != nil {
		return nil, err
	}
	return this.NewSession(conn), nil
}

// 监听端口
func (this *Server) Listener() net.Listener {
	return this.listener
}

// 关闭
func (this *Server) Stop() bool {
	if atomic.CompareAndSwapInt32(&this.stopFlag, 0, 1) {		
		this.listener.Close()
		
		for _, session := range this.sessions {
			session.Close()
		}
		
		this.syncGroupStop.Wait()
		return true
	}
	return false
}

// 创建session
func (this *Server) NewSession(conn net.Conn) (*Session) {
	this.syncMutexSession.Lock()
	defer this.syncMutexSession.Unlock()
	
	session := NewSession(conn, this.codecType)
	this.sessions[session.Id()] = session
	// 增加回调
	session.AddCloseCallback(this, func(){
		this.DelSession(session)
	})
	
	this.syncGroupStop.Add(1)
	return session
}

// 删除session
func (this *Server) DelSession(session *Session) {
	this.syncMutexSession.Lock()
	defer this.syncMutexSession.Unlock()
	
	delete(this.sessions, session.Id())
	this.syncGroupStop.Done()
}


