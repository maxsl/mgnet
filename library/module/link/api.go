package link

import (
	"net"
	"time"
	"github.com/goodkele/mgnet/library/module"
)

// 创建服务
func Serve(network, address string, codecType module.CodecType) (*Server, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return NewServer(listener, codecType), nil
}

// 创建连接
func Connect(network, address string, codecType module.CodecType) (*Session, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewSession(conn, codecType), nil
}

// 创建带超时连接
func ConnectTimeout(network, address string, timeout time.Duration, codecType module.CodecType) (*Session, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return NewSession(conn, codecType), nil
}