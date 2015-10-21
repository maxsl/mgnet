package protocol

import (
	"io"
	"github.com/goodkele/mgnet/library/module/proto"
	"github.com/goodkele/mgnet/library/module"
)

// 创建解析器
type CodecType struct {
	
}

func (this *CodecType) NewEncoder(w io.Writer) module.Encoder {
	return &Encode{w}
}

func (this *CodecType) NewDecoder(r io.Reader) module.Decoder {
	return &Decode{r, make([]byte, 1024, 1024)}
}

// 序列化
type Encode struct {
	write io.Writer
}

// 写入
func (this *Encode) Encode(msg interface{}) error {
	
	var err error
	var message []byte
	if buf, ok := msg.(proto.Message); ok == true {		
		if message, err = proto.Marshal(buf); err == nil {
			_, err = this.write.Write(message)
		}
	}
	return err
	
}

// 反序列化
type Decode struct {
	read io.Reader
	p []byte
}

// 读取
func (this *Decode) Decode(msg interface{}) error {

	var err error
	if buf, ok := msg.(proto.Message); ok == true {	
		_, err = this.read.Read(this.p[0:])
		if err == nil {
			err = proto.Unmarshal(this.p, buf)
		}
	}
	return err
	
}

