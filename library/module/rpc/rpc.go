package rpc

import (
	"github.com/goodkele/mgnet/library/module/proto"
	"errors"
	"fmt"
)

type RpcService struct {
	
	funcs map[int] func(proto.Message) (proto.Message, error)
}

func New() *RpcService {
	return &RpcService{
		funcs	:	make(map[int] func(proto.Message) (proto.Message, error)),
	}
}

// 注册RPC
func (this *RpcService) Register(rcpIndex int, funcHandle func(proto.Message) (proto.Message, error)) {
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
func (this *RpcService) Exec(rcpIndex int, req proto.Message) (proto.Message, error) {
	funcHandle, ok := this.funcs[rcpIndex]
	if ok == false {
		return nil, errors.New(fmt.Sprintf(ERROR_RPC_NOT_EXISTS,  rcpIndex))
	}
	res, err := funcHandle(req)
	return res, err
}