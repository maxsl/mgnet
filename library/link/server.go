package link

import (
//	"fmt"
	"net"
	"io"
	"sync"
)

type CodecType interface {
	EncodeType
	DecodeType
}

type EncodeType interface {
	NewEncoder(w io.Writer) Encoder
}

type DecodeType interface {
	NewDecoder(r io.Reader) Decoder
}

type Encoder interface {
	Encode(msg interface{}) error
}

type Decoder interface {
	Decode(msg interface{}) error
}


type Server struct {
	listener  net.Listener
	codecType CodecType

	// About sessions
	maxSessionId uint64
	sessions     map[uint64]*Session
	sessionMutex sync.Mutex
}

// 创建服务器
func NewServer(listener net.Listener) *Server{
//	, codecType CodecType
	server := &Server{
		listener : listener,
//		codecType : codecType,
		sessions : make(map[uint64]*Session),
	}
	
	return server
}

// 创建session
func (this *Server) newSession(conn net.Conn) *Session {
	session := NewSession(conn)
	
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