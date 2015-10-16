package main

import (
	"fmt"
	"github.com/goodkele/mgnet/library/types"
	"github.com/golang/protobuf/proto"
)

func main() {
	
//	var pro = proto.Marshal
	

	req := &types.SearchRequest{}
	
	//req2 := &types.SearchRequest{}
	
	req.Query = "asdfasdf"
	
	pb := proto.MarshalTextString(req)

	//pb,_ := proto.MarshalMessageSetJSON(req)
	// pb,_ := proto.Marshal(req)
	
	fmt.Println(pb)

	//proto.Unmarshal(pb, req2)
	//proto.unm(pb, req2)
	
	//fmt.Println(req2)
	
	
	fmt.Println(req)
}