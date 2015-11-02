// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 路由类型
type Routing struct {
	Target int32  `protobuf:"varint,1,opt,name=target" json:"target,omitempty"`
	Refer  int32  `protobuf:"varint,2,opt,name=refer" json:"refer,omitempty"`
	Url    int32  `protobuf:"varint,3,opt,name=url" json:"url,omitempty"`
	SessId string `protobuf:"bytes,4,opt,name=sessId" json:"sessId,omitempty"`
	Msg    []byte `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Gmt    int32  `protobuf:"varint,6,opt,name=gmt" json:"gmt,omitempty"`
	Code   int32  `protobuf:"varint,7,opt,name=code" json:"code,omitempty"`
	Error  string `protobuf:"bytes,8,opt,name=error" json:"error,omitempty"`
}

func (m *Routing) Reset()         { *m = Routing{} }
func (m *Routing) String() string { return proto.CompactTextString(m) }
func (*Routing) ProtoMessage()    {}