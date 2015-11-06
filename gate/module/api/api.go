package api

import (
	"github.com/goodkele/mgnet/library/module/rpc"
	"github.com/goodkele/mgnet/gate/module/gate"
)

var (
	Rpc = rpc.New()
)
	

func init() {
	Rpc.Register(1, gate.API_Echo)

}


