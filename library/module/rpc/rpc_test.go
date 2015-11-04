package rpc

import (
	"testing"
	"github.com/goodkele/mgnet/library/module/link"
	"github.com/goodkele/mgnet/library/module/proto"
	"fmt"
	
)


func echo(rpcIndex uint32, refer uint32, session *link.Session, msg proto.Message, gmt uint32) (proto.Message, uint32, error) {
	fmt.Println(rpcIndex)
	fmt.Println(refer)
	fmt.Println(session)
	fmt.Println(msg)
	fmt.Println(gmt)
	
	return nil, 0, nil
}

func Test_All(t *testing.T) {
	
	var rpc *RpcService
	
	rpc = New()

	rpc.Register(1, echo)
	
	
	rpc.Exec(1, 1, &link.Session{}, nil, 10)

//	funcs map[uint32] func(uint32, uint32, *link.Session, proto.Message, uint32) (proto.Message, uint32, error)
	
}