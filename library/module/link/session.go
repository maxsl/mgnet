package link

import (
	"net"
	"github.com/goodkele/mgnet/library/module"
	"sync/atomic"
)

var (
	globalSessionId uint64
)

type Session struct {
	id	uint64

	conn net.Conn
}

// 创建session
func NewSession(conn net.Conn, codecType module.CodecType) *Session {
	sessionId := atomic.AddUint64(&globalSessionId, 1)
	session := &Session{
		conn : conn,
		id : sessionId,
	}
	return session
}

// 网络连接
func (this *Session) Conn() net.Conn {
	return this.conn
}

// Session Id
func (this *Session) Id() uint64 {
	return this.id
}

// 关闭Session
func (this *Session) Close() {
	this.conn.Close()
}