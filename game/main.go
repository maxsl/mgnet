package main

import (
	"fmt"
<<<<<<< HEAD
	_ "github.com/goodkele/mgnet/library/types"
	_ "github.com/golang/protobuf/proto"
=======
	"github.com/goodkele/mgnet/library/types"
	"github.com/golang/protobuf/proto"
>>>>>>> parent of bdf65c8... protobuf
)

func main() {
	
<<<<<<< HEAD
	fmt.Println("Hello mgnet")
=======
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
>>>>>>> parent of bdf65c8... protobuf
}