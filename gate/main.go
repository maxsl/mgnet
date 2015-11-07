package main

import (
//	"net"
//	"time"
	"runtime"
	
//	"github.com/goodkele/mgnet/gate/module/api"
	"github.com/goodkele/mgnet/gate/module/epool"
//	"github.com/goodkele/mgnet/library/module/constant"
	"github.com/goodkele/mgnet/library/module/mglog"
	//"github.com/goodkele/mgnet/library/module/link"
//	"github.com/goodkele/mgnet/library/module/protocol"
//	"github.com/goodkele/mgnet/library/module/types"
)

func init() {
	// 初始化日志
	mglog.InitLog("./", "gate", 0, &mglog.SWITCHER_DAY)
}



func main() {
	
	runtime.SetCPUProfileRate(4)
	
	var ptId		uint32
	var serverId	uint32
	
	address := ":10011"
	
	

	server := epool.NewServer(address, ptId, serverId)
	
	
	server.Serve()
	
	//server.Stop()
	
	server.SyncGroupStop.Wait()
	
	
	
	//api.Rpc.Exec(1, 1, &link.Session{}, nil, 10)


	
	
	//time.Sleep(5 * time.Second)

	
}