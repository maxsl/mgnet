package rpc

import (
	"github.com/goodkele/mgnet/library/module/proto"
	"github.com/goodkele/mgnet/library/module/link"
	"errors"
	"fmt"
)

type RpcService struct {
	// RPC函数参数： rpcIndex, refer, session, msg, gmt
	funcs map[uint32] func(uint32, uint32, *link.Session, proto.Message, uint32) (proto.Message, uint32, error)
}

// 新建RPC函数列表
func New() *RpcService {
	return &RpcService{
		funcs	:	make(map[uint32] func(uint32, uint32, *link.Session, proto.Message, uint32) (proto.Message, uint32, error)),
	}
}

// 注册RPC
func (this *RpcService) Register(rcpIndex uint32, funcHandle func(uint32, uint32, *link.Session, proto.Message, uint32) (proto.Message, uint32, error)) {
	this.funcs[rcpIndex] = funcHandle
}

// 移除RPC
func (this *RpcService) Remove(rcpIndex uint32) {
	delete(this.funcs, rcpIndex)
}

// 是否存在
func (this *RpcService) IsExists(rcpIndex uint32) bool {
	_, ok := this.funcs[rcpIndex]
	return ok
}

// 执行RPC函数
func (this *RpcService) Exec(rpcIndex uint32, refer uint32, session *link.Session, msg proto.Message, gmt uint32) (proto.Message, uint32, error) {
	funcHandle, ok := this.funcs[rpcIndex]
	if ok == false {
		return nil, 0, errors.New(fmt.Sprintf(ERROR_RPC_NOT_EXISTS,  rpcIndex))
	}
	res, code, err := funcHandle(rpcIndex, refer, session, msg, gmt)
	return res, code, err
}