package main

import (
	"runtime"
	"github.com/goodkele/mgnet/game/module/epool"
	"github.com/goodkele/mgnet/library/module/mglog"
)

func init() {
	// 初始化日志
	mglog.InitLog("./", "game", 0, &mglog.SWITCHER_DAY)
}

func main() {
	runtime.SetCPUProfileRate(4)
	
	var ptId		uint32
	var serverId	uint32
	
	address := ":10012"
	server := epool.NewServer(address, ptId, serverId)
	
	server.Serve()
	//server.Stop()
	server.SyncGroupStop.Wait()
	//api.Rpc.Exec(1, 1, &link.Session{}, nil, 10)
	//time.Sleep(5 * time.Second)
}