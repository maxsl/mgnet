package gate

import (
	"github.com/goodkele/mgnet/library/module/proto"
	"github.com/goodkele/mgnet/library/module/link"
	"fmt"
)


func API_Echo(rpcIndex uint32, refer uint32, session *link.Session, msg proto.Message, gmt uint32) (proto.Message, uint32, error) {
	fmt.Println(rpcIndex)
	fmt.Println(refer)
	fmt.Println(session)
	fmt.Println(msg)
	fmt.Println(gmt)
	
	return nil, 0, nil
}


func API_Receive(rpcIndex uint32, refer uint32, session *link.Session, msg proto.Message, gmt uint32) (proto.Message, uint32, error) {
	
	fmt.Println(rpcIndex)
	fmt.Println(refer)
	fmt.Println(session)
	fmt.Println(msg)
	fmt.Println(gmt)
	
	return nil, 0, nil
}
