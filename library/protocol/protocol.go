package protocol

import (
	"io"
	"github.com/golang/protobuf/proto"
)

// 创建解析器
type CodecType struct {
	
}

func (this *CodecType) NewEncoder(w io.Writer) Encoder {
	return &Encoder{w}
}

func (this *CodecType) NewDecoder(r io.Reader) Decoder {
	return &Decoder{r, make([]byte, 0, 1024)}
}



// 序列化
type Encoder struct {
	write io.Writer
}

// 写入
func (this *Encoder) Encode(msg interface{}) error {
	
	if buf, ok := msg.(proto.Message); ok == true {		
		if message, err := proto.Marshal(buf); err == nil {
			// n, err := 
			write.Write(message)
		}
	}
	
}

// 反序列化
type Decoder struct {
	read io.Reader
	p []byte
}

// 读取
func (this *Decode) Decode(msg interface{}) error {
	
	if buf, ok := msg.(proto.Message); ok == true {	
		n, err := read.Read(p[0:])
		err = proto.Unmarshal(p, buf)
	}
	
}

