package link

import (
	"net"
	"github.com/goodkele/mgnet/library/module"
	"sync/atomic"
	"container/list"
	"sync"
)

var (
	globalSessionId uint64
)

type Session struct {
	id			uint64				// Session Id
	conn 		net.Conn			// 连接
	encoder 	module.Encoder		// 编码器
	decoder 	module.Decoder		// 解码器
	
	syncMutexSend 	sync.Mutex		// 锁，发送时
	syncMutexReceive sync.Mutex		// 锁，收到时
	
	closeFlag 	int32				// 是否关闭，0：正常，1：关闭
	
	callbacks 	*list.List			// 回调队列
}

// 创建session
func NewSession(conn net.Conn, codecType module.CodecType) *Session {

	session := &Session{
		id 		: atomic.AddUint64(&globalSessionId, 1),
		conn 	: conn,
		encoder : codecType.NewEncoder(conn),
		decoder : codecType.NewDecoder(conn),
		callbacks : list.New(),
	}
	
	return session
}

// 网络连接
func (this *Session) Conn() net.Conn {
	return this.conn
}

// Session Id
func (this *Session) Id() uint64 {
	return this.id
}

// Session closeFlag
func (this *Session) IsClose() (bool) {
	return atomic.LoadInt32(&this.closeFlag) != 0
}

// 收到消息
func (this *Session) Receive(msg interface{}) (err error) {
	this.syncMutexReceive.Lock()
	defer this.syncMutexReceive.Unlock()

	if err := this.decoder.Decode(msg); err != nil {
		return err
	}
	return 
}

// 发送消息
func (this *Session) Send(msg interface{}) (err error) {
	this.syncMutexSend.Lock()
	defer this.syncMutexSend.Unlock()
	
	if err := this.encoder.Encode(msg); err != nil {
		return err
	}
	return
}

// 关闭
func (this *Session) Close() {
	if (atomic.CompareAndSwapInt32(&this.closeFlag, 0, 1)) {
		this.ExecCloseCallback()
		this.conn.Close()
	}
}



// 回调
type closeCallback struct {
	Handler interface{}
	Func    func()
}

// 增加回调
func (this *Session) AddCloseCallback(handler interface{}, fun func()) {
	if (this.IsClose()) {
		return
	}

	this.callbacks.PushBack(&closeCallback{handler, fun})
}

// 删除回调
func (this *Session) DelCloseCallback(handler interface{}) {
	if (this.IsClose()) {
		return
	}
	for element := this.callbacks.Front(); element != nil; element = element.Next() {
		if element.Value.(closeCallback).Handler == handler {
			this.callbacks.Remove(element)
		}
	}
}

// 执行回调
func (this *Session) ExecCloseCallback() {
	for element := this.callbacks.Front(); element != nil; element = element.Next() {
		element.Value.(*closeCallback).Func()
	}
}