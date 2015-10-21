package link

import (
	"net"
	"sync"
	"github.com/goodkele/mgnet/library/module"
)

type Server struct {
	listener  net.Listener
	codecType module.CodecType

	// About sessions
	maxSessionId int64
	sessions     map[int64]*Session
	sessionMutex sync.Mutex
}

// 创建服务器
func NewServer(listener net.Listener, codecType module.CodecType) *Server{
	server := &Server{
		listener : listener,
		codecType : codecType,
		sessions : make(map[int64]*Session),
	}
	
	return server
}

// 创建session
func (this *Server) newSession(conn net.Conn) *Session {
	session := NewSession(conn, this.codecType)
	
	return session
}

// 等待一个连接
func (this *Server) Accept() (*Session, error) {
	conn, err := this.listener.Accept()
	if err != nil {
		return nil, err
	}
	return this.newSession(conn), nil
}

// 监听端口
func (this *Server) Listener() net.Listener {
	return this.listener
}

// 停止服务
func (this *Server) Stop() {
	this.listener.Close()
	
}

// 关闭
func (this *Server) Close() {
	
}