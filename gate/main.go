package main

import (
	
	"net"
	"time"
	"github.com/goodkele/mgnet/library/module/constant"
	"github.com/goodkele/mgnet/library/module/mglog"
	"github.com/goodkele/mgnet/library/module/link"
	"github.com/goodkele/mgnet/library/module/protocol"
)

func init() {
	// 初始化日志
	mglog.InitLog("./", "gate", 0, &mglog.SWITCHER_DAY)
}

func main() {

	address := ":10011"
	listener, err := net.Listen("tcp", address)
	if err != nil {		
		mglog.Error(constant.ERROR_GATE_LISTENER, err)
		return
	}

	

//	func Serve(network, address string, codecType module.CodecType) (*Server, error) {
	link.Serve("TCP", address, )

	mglog.Debug("%v", listener)
	
	mglog.Info("assssdf")
	
	
	time.Sleep(5 * time.Second)

	
	
}