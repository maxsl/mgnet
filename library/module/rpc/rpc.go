package rpc

import (
	"github.com/goodkele/mgnet/library/module/proto"
	"github.com/goodkele/mgnet/library/module/link"
	"errors"
	"fmt"
)

type RpcService struct {
	// RPC函数参数： rpcIndex, refer, session, msg, gmt
	funcs map[int32] func(int32, int32, *link.Session, proto.Message, int32) (proto.Message, int32, error)
}

func New() *RpcService {
	return &RpcService{
		funcs	:	make(map[int32] func(int32, int32, *link.Session, proto.Message, int32) (proto.Message, int32, error)),
	}
}

// 注册RPC
func (this *RpcService) Register(rcpIndex int, funcHandle func(int32, int32, *link.Session, proto.Message, int32) (proto.Message, int32, error)) {
	this.funcs[rcpIndex] = funcHandle
}

// 移除RPC
func (this *RpcService) Remove(rcpIndex int) {
	delete(this.funcs, rcpIndex)
}

// 是否存在
func (this *RpcService) IsExists(rcpIndex int) bool {
	_, ok := this.funcs[rcpIndex]
	return ok
}

// 执行RPC函数
func (this *RpcService) Exec(rpcIndex int32, refer int32, session *link.Session, msg proto.Message, gmt int32) (proto.Message, int32, error) {
	funcHandle, ok := this.funcs[rcpIndex]
	if ok == false {
		return nil, errors.New(fmt.Sprintf(ERROR_RPC_NOT_EXISTS,  rcpIndex))
	}
	res, code, err := funcHandle(rpcIndex, refer, session, msg, gmt)
	return res, code, err
}