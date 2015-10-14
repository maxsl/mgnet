package link

import (
	"net"
)

type Session struct {
	id      uint64

	conn net.Conn
}

// 创建session
func NewSession(conn net.Conn) *Session {
	session := &Session{
	conn : conn,
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