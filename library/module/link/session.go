package link

import (
	"net"
	"github.com/goodkele/mgnet/library/module"
)

type Session struct {
	id	int64

	conn net.Conn

}

// 创建session
func NewSession(conn net.Conn, codecType module.CodecType) *Session {
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
func (this *Session) Id() int64 {
	return this.id
}

// 关闭Session
func (this *Session) Close() {
	this.conn.Close()
}