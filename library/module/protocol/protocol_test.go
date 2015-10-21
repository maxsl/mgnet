package protocol

import (
	"fmt"
	"testing"
	"github.com/goodkele/mgnet/library/module/types"
	"os"
)


func Test_Encode(t *testing.T) {
	
	fileName := "../../../log/wulei.txt"

	fd, _ := os.OpenFile(fileName, os.O_CREATE, 0755)

	codecType := &CodecType{}
	
	encode := codecType.NewEncoder(fd)
	
	seReq := &types.SearchRequest{"name1", 110, 20}
	
	encode.Encode(seReq)
	
	fmt.Println(seReq)

	fd.Sync()
	fd.Close()
}

func Test_Decode(t *testing.T) {
	
	fileName := "../../../log/wulei.txt"

	fd, _ := os.OpenFile(fileName, os.O_CREATE, 0755)

	codecType := &CodecType{}

	seReq := &types.SearchRequest{}
	
	decode := codecType.NewDecoder(fd)
	
	decode.Decode(seReq)
	
	fmt.Println(seReq)
	
	fd.Sync()
	fd.Close()
}