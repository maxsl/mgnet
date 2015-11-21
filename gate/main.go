package main

import (
	"runtime"
	"github.com/goodkele/mgnet/gate/module/epool"
	"github.com/goodkele/mgnet/library/module/mglog"
	"github.com/goodkele/mgnet/library/module/constant"
)

func init() {
	// 初始化日志
	mglog.InitLog("./log", "gate", 0, &mglog.SWITCHER_DAY)
}

func main() {
	runtime.SetCPUProfileRate(4)


	var ptId		uint32
	var serverId	uint32
	
	address := ":10011"
	server := epool.NewServer(address, ptId, serverId)
	server.Serve()
	// 链接Game游戏服务器
	server.Connect(constant.SERVER_GAME, ":10012")
	
	//server.Stop()
	
	server.SyncGroupStop.Wait()
	
	//api.Rpc.Exec(1, 1, &link.Session{}, nil, 10)
	//time.Sleep(5 * time.Second)
}