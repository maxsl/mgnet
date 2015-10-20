package protocol

import (
	"fmt"
	"testing"
	//"github.com/goodkele/mgnet/library/proto"
	"github.com/goodkele/mgnet/library/types"
	"os"
)

func Test_Encode(t *testing.T) {
	
	fileName := "../../log/wulei.txt"

//os.O_WRONLY|
	fd, _ := os.OpenFile(fileName, os.O_CREATE, 0755)


	codecType := &CodecType{}
	
//	encode := codecType.NewEncoder(fd)
	
	seReq := &types.SearchRequest{}
	
	//encode.Encode(seReq)
	
	decode := codecType.NewDecoder(fd)
	
	
	
	er := decode.Decode(seReq)
	
	fmt.Println(er)
	
	
	fmt.Println("asdf")
	fmt.Println(seReq)
	
	
	
	
	
	
	
	

	
	//fd.WriteString("package orm\n\n")

	



	fd.Sync()
	fd.Close()

	fmt.Println(t)
}